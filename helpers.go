package main

import (
	"crypto/md5"
	"crypto/rand"
	"encoding/base64"
	"encoding/binary"
	"fmt"
	"io"
	"net"
	"os"
	"time"

	gosync "github.com/Redundancy/go-sync"
	"github.com/Redundancy/go-sync/blocksources"
	"github.com/Redundancy/go-sync/filechecksum"
)

func uid() string {
	buf := make([]byte, 10)

	if _, err := rand.Read(buf); err != nil {
		panic(err)
	}
	str := fmt.Sprintf("%d%x", time.Now().Unix(), buf[0:10])
	return base64.StdEncoding.EncodeToString([]byte(str))
}

func copy(src, dst string) error {
	from, err := os.OpenFile(src, os.O_RDONLY, 0)
	if err != nil {
		return err
	}
	defer from.Close()

	to, err := os.OpenFile(dst, os.O_WRONLY|os.O_CREATE, 0774)
	if err != nil {
		return err
	}
	defer to.Close()

	_, err = io.Copy(to, from)
	if err != nil {
		return err
	}
	return nil
}

func intToBytes(val int) []byte {
	bs := make([]byte, 4)
	binary.LittleEndian.PutUint32(bs, uint32(val))
	return bs
}

func int64ToBytes(val int64) []byte {
	bs := make([]byte, 8)
	binary.LittleEndian.PutUint64(bs, uint64(val))
	return bs
}

// MakeRSync rsync from remote
func makeRSync(local gosync.ReadSeekerAt, remote string, output io.Writer, fs gosync.FileSummary) *gosync.RSync {
	return &gosync.RSync{
		Input:  local,
		Output: output,
		Source: blocksources.NewHttpBlockSource(
			remote,
			1,
			blocksources.MakeFileSizedBlockResolver(
				uint64(fs.GetBlockSize()),
				fs.GetFileSize(),
			),
			&filechecksum.HashVerifier{
				Hash:                md5.New(),
				BlockSize:           fs.GetBlockSize(),
				BlockChecksumGetter: fs,
			},
		),
		Summary: fs,
		OnClose: nil,
	}
}

func createSummary(r io.Reader) (gosync.FileSummary, error) {
	blockSize := int64(BlockSize)
	fileSize, referenceFileIndex, checksumLookup, err := DecodeChecksumIndex(r)
	if err != nil {
		return nil, err
	}

	if fileSize == 0 {
		return &gosync.BasicSummary{}, nil
	}

	blockCount := fileSize / int64(blockSize)
	if fileSize%blockSize != 0 {
		blockCount++
	}

	fs := &gosync.BasicSummary{
		ChecksumIndex:  referenceFileIndex,
		ChecksumLookup: checksumLookup,
		BlockCount:     uint(blockCount),
		BlockSize:      BlockSize,
		FileSize:       fileSize,
	}

	return fs, nil
}

func Multicast(port int, b []byte) (int, error) {
	addr, err := net.ResolveUDPAddr("udp", fmt.Sprintf("%s:%d", MulticastAddr, port))
	if err != nil {
		return 0, err
	}
	c, err := net.DialUDP("udp", nil, addr)
	defer c.Close()
	n, err := c.Write(b)
	if err != nil {
		return 0, err
	}
	return n, nil
}
