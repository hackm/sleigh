package main

import (
	"crypto/md5"
	"crypto/rand"
	"crypto/rsa"
	"crypto/tls"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/base64"
	"encoding/pem"
	"fmt"
	"io"
	"io/ioutil"
	"math/big"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/Redundancy/go-sync"
	"github.com/Redundancy/go-sync/blocksources"
	"github.com/Redundancy/go-sync/filechecksum"
	"github.com/fsnotify/fsnotify"
)

const (
	crtFile = ".sleigh/cer.pem"
	keyFile = ".sleigh/key.pem"
)

var blockSize int64 = 4 * MB

type Differ struct {
	hostname      string
	port          int
	root          string
	Notifications chan (Notification)
	Errors        chan (error)
}

func NewDiffer(hostname string, port int, root string) *Differ {
	return &Differ{
		hostname:      hostname,
		port:          port,
		root:          root,
		Notifications: make(chan (Notification)),
		Errors:        make(chan (error)),
	}
}

// Start differ
func (d *Differ) Start() error {
	geneCrt(crtFile, keyFile, d.hostname)

	s := http.NewServeMux()
	s.HandleFunc("/contents", d.createContentHandler())
	s.HandleFunc("/summaries", d.createSummaryHandler())

	err := http.ListenAndServeTLS(fmt.Sprintf(":%d", d.port), crtFile, keyFile, s)
	if err != nil {
		return err
	}

	go syncDeamon(d)

	return nil
}

// Close differ
func (d *Differ) Close() {
}

// Download file(diff sync)
func (d *Differ) Download(path, hostname string, port int) (*os.File, error) {
	contentURL := fmt.Sprintf("https://%s:%d/contents/%s", hostname, port, path)
	summaryURL := fmt.Sprintf("https://%s:%d/summaries/%s?blockSize=%%d", hostname, port, path)

	fs, err := getSummary(summaryURL, blockSize)
	if err != nil {
		return nil, err
	}
	local := filepath.Join(d.root, path)
	input, err := os.OpenFile(local, os.O_RDONLY|os.O_CREATE, 0)
	if err != nil {
		return nil, err
	}

	temp, err := ioutil.TempFile(hostname, uid())
	if err != nil {
		return nil, err
	}
	// defer os.Remove(temp.Name())

	rsync := makeRSync(input, contentURL, temp, fs)
	defer rsync.Close()

	err = rsync.Patch()

	if err != nil {
		return nil, err
	}

	// output, err := os.OpenFile(local, os.O_WRONLY|os.O_CREATE, 0)
	// if err != nil {
	// 	return err
	// }
	// err = output.Truncate(0)
	// if err != nil {
	// 	return err
	// }
	// _, err = io.Copy(output, temp)
	// if err != nil {
	// 	return err
	// }

	return temp, nil
}

func syncDeamon(d *Differ) {
	for {
		select {
		case n := <-d.Notifications:
			switch n.Event {
			case fsnotify.Create, fsnotify.Write:
				if n.Type == File {
					temp, err := d.Download(n.Path, n.Hostname, d.port)
					if err != nil {
						d.Errors <- err
						continue
					}
					for {
						output, err := os.OpenFile(filepath.Join(d.root, n.Path), os.O_WRONLY|os.O_CREATE, 0)
						if err != nil {
							d.Errors <- err
							time.Sleep(3 * time.Second)
							continue
						}
						err = output.Truncate(0)
						if err != nil {
							d.Errors <- err
							time.Sleep(3 * time.Second)
							continue
						}
						_, err = io.Copy(output, temp)
						if err != nil {
							d.Errors <- err
							time.Sleep(3 * time.Second)
							continue
						}
						break
					}
				} else {
					err := os.MkdirAll(filepath.Join(d.root, n.Path), os.ModeDir)
					if err != nil {
						d.Errors <- err
					}
				}
			case fsnotify.Rename, fsnotify.Remove:
				for {
					err := os.Remove(n.Path)
					if err != nil {
						d.Errors <- err
						time.Sleep(3 * time.Second)
						continue
					}
					break
				}
			}
		}
	}
}

func (d *Differ) createContentHandler() func(w http.ResponseWriter, req *http.Request) {
	// handler for content download
	return func(w http.ResponseWriter, req *http.Request) {
		path := filepath.Join(d.root, strings.Replace(req.URL.Path, "/contents/", "", 1))
		if _, err := os.Stat(path); err != nil {
			http.NotFound(w, req)
			return
		}
		http.ServeFile(w, req, path)
	}
}

func (d *Differ) createSummaryHandler() func(w http.ResponseWriter, req *http.Request) {
	// handler for checksum index download
	return func(w http.ResponseWriter, req *http.Request) {
		var blockSize uint64 = 1024 * 1024
		blockSize, err := strconv.ParseUint(req.URL.Query().Get("blockSize"), 10, 32)

		path := filepath.Join(d.root, strings.Replace(req.URL.Path, "/summaries/", "", 1))
		file, err := os.OpenFile(path, os.O_RDONLY, 0)
		if err != nil {
			http.NotFound(w, req)
			return
		}

		defer file.Close()

		info, err := file.Stat()
		if err != nil {
			http.NotFound(w, req)
			return
		}

		b, err := EncodeChecksumIndex(file, info.Size(), uint(blockSize))
		if err != nil {
			http.NotFound(w, req)
			return
		}

		http.ServeContent(w, req, "", time.Now(), b)
	}
}

func uid() string {
	buf := make([]byte, 10)

	if _, err := rand.Read(buf); err != nil {
		panic(err)
	}
	str := fmt.Sprintf("%d%x", time.Now().Unix(), buf[0:10])
	return base64.StdEncoding.EncodeToString([]byte(str))
}

func geneCrt(crtFile, keyFile, hostname string) error {
	priv, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return err
	}
	notBefore := time.Now()
	notAfter := notBefore.Add(24 * time.Hour)
	serialNumberLimit := new(big.Int).Lsh(big.NewInt(1), 256)
	serialNumber, err := rand.Int(rand.Reader, serialNumberLimit)
	if err != nil {
		return err
	}

	template := x509.Certificate{
		SerialNumber: serialNumber,
		Subject: pkix.Name{
			Organization: []string{hostname},
		},

		NotBefore: notBefore,
		NotAfter:  notAfter,

		KeyUsage:              x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
		BasicConstraintsValid: true,
		DNSNames:              []string{"localhost", hostname},
	}

	derBytes, err := x509.CreateCertificate(rand.Reader, &template, &template, &priv.PublicKey, priv)
	if err != nil {
		return err
	}

	cert, err := os.Create(crtFile)
	if err != nil {
		return err
	}
	defer cert.Close()

	key, err := os.Create(keyFile)
	if err != nil {
		return err
	}
	defer key.Close()

	pem.Encode(cert, &pem.Block{Type: "CERTIFICATE", Bytes: derBytes})
	pem.Encode(key, &pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(priv)})

	return nil
}

// NewClient create HttpClient for LAN https. this client skip verify
func newClient() *http.Client {
	return &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}
}

// GetSummary get summary from remote
func getSummary(urlFormat string, blockSize int64) (gosync.FileSummary, error) {
	res, err := newClient().Get(fmt.Sprintf(urlFormat, blockSize))
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	fileSize, referenceFileIndex, checksumLookup, err := DecodeChecksumIndex(res.Body)
	if err != nil {
		return nil, err
	}

	blockCount := fileSize / blockSize
	if fileSize%blockSize != 0 {
		blockCount++
	}

	fs := &gosync.BasicSummary{
		ChecksumIndex:  referenceFileIndex,
		ChecksumLookup: checksumLookup,
		BlockCount:     uint(blockCount),
		BlockSize:      uint(blockSize),
		FileSize:       fileSize,
	}

	return fs, nil
}

// MakeRSync rsync from remote
func makeRSync(local gosync.ReadSeekerAt, remote string, output io.Writer, fs gosync.FileSummary) *gosync.RSync {
	return &gosync.RSync{
		Input:  local,
		Output: output,
		Source: blocksources.NewBlockSourceBase(
			&HttpRequester{
				Url:    remote,
				Client: http.DefaultClient,
			},
			blocksources.MakeFileSizedBlockResolver(
				uint64(fs.GetBlockSize()),
				fs.GetFileSize(),
			),
			&filechecksum.HashVerifier{
				Hash:                md5.New(),
				BlockSize:           fs.GetBlockSize(),
				BlockChecksumGetter: fs,
			},
			1,
			4*MB,
		),
		Summary: fs,
		OnClose: nil,
	}
}
