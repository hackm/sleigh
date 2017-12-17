package main

import (
	"bytes"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/fsnotify/fsnotify"

	"github.com/pkg/errors"
)

type RemoteUrlResolver func(n Notification) string
type LocalPathResolver func(relPath string) string

type Patcher struct {
	Notifications    chan Notification
	Errors           chan error
	remoteURLResolve RemoteUrlResolver
	localPathResolve LocalPathResolver
}

func NewPatcher(n chan Notification, remoteUrlResolve RemoteUrlResolver, localPathResolve LocalPathResolver) *Patcher {
	return &Patcher{
		Notifications:    n,
		Errors:           make(chan error, 1),
		remoteURLResolve: remoteUrlResolve,
		localPathResolve: localPathResolve,
	}
}
func (p Patcher) equals(n Notification) (bool, error) {
	fullpath := p.localPathResolve(n.RelPath)
	stat, err := os.Stat(fullpath)
	if err != nil {
		// file not exists
		return false, nil
	}
	file, err := os.Open(fullpath)
	if err != nil {
		return false, errors.Wrapf(err, "Cannot open %s: %s", fullpath, err)
	}

	current, err := encodeChecksumIndex(file, stat.Size(), BlockSize)

	return bytes.Equal(current, n.ChecksumIndex), nil
}

func (p Patcher) patch(n Notification) error {
	log.Println(n.RelPath)
	eq, err := p.equals(n)
	if err != nil {
		return errors.Wrapf(err, "Cannot check checsumIndex equality: %s", err)
	}

	if eq == true {
		return nil
	}

	contentURL := p.remoteURLResolve(n)
	fullpath := p.localPathResolve(n.RelPath)
	dir := filepath.Dir(fullpath)

	if _, err := os.Stat(dir); err != nil {
		if err := os.MkdirAll(dir, 0755); err != nil {
			return errors.Wrapf(err, "Cannot create dir: %s", err)
		}
	}

	input, err := os.OpenFile(fullpath, os.O_RDONLY|os.O_CREATE, 0755)
	if err != nil {
		return errors.Wrapf(err, "Cannot open %s: %s", fullpath, err)
	}
	defer input.Close()

	temp, err := ioutil.TempFile("", uid())
	if err != nil {
		return errors.Wrapf(err, "Cannot create file: %s", err)
	}
	defer temp.Close()

	summary, err := createSummary(bytes.NewReader(n.ChecksumIndex))
	if err != nil {
		return errors.Wrapf(err, "Cannot create summary: %s", err)
	}

	rsync := makeRSync(input, contentURL, temp, summary)
	defer rsync.Close()

	err = rsync.Patch()

	if err != nil {
		return errors.Wrapf(err, "Cannot gosync patch: %s", err)
	}

	err = copy(temp.Name(), fullpath)
	if err != nil {
		return errors.Wrapf(err, "Cannot replace %s -> %s: %s", temp.Name(), fullpath, err)
	}

	return nil
}

func (p Patcher) download(n Notification) error {
	fullpath := p.localPathResolve(n.RelPath)
	file, err := os.OpenFile(fullpath, os.O_WRONLY|os.O_CREATE, 0755)
	if err != nil {
		return errors.Wrapf(err, "Cannot prepare write out file: %s", err)
	}
	defer file.Close()

	contentURL := p.remoteURLResolve(n)
	res, err := http.DefaultClient.Get(contentURL)
	if err != nil || res.StatusCode != 200 {
		return errors.Wrapf(err, "Cannot get %s: %s, %s", contentURL, res.Status, err)
	}
	defer res.Body.Close()

	_, err = io.Copy(file, res.Body)
	if err != nil {
		return errors.Wrapf(err, "Cannot write out to %s: %s", fullpath, err)
	}

	return nil
}

func (p Patcher) remove(n Notification) error {
	fullpath := p.localPathResolve(n.RelPath)
	if _, err := os.Stat(fullpath); err != nil {
		return nil
	}
	if err := os.Remove(fullpath); err != nil {
		return errors.Wrapf(err, "Cannot remove %s: %s", fullpath, err)
	}
	return nil
}

func (p Patcher) rename(n Notification) error {
	fullpath := p.localPathResolve(n.RelPath)
	oldpath := p.localPathResolve(n.OldRelPath)
	if _, err := os.Stat(oldpath); err != nil {
		return nil
	}
	if err := os.Rename(oldpath, fullpath); err != nil {
		return errors.Wrapf(err, "Cannot rename %s: %s", fullpath, err)
	}
	return nil
}

func (p *Patcher) try(fn func() error) {
	go func() {
		for i := 0; i < RetryMax; i++ {
			if err := fn(); err != nil {
				p.Errors <- err
				continue
			}
			break
		}
	}()
}

func (p *Patcher) Patch(n Notification) {
	p.try(func() error {
		return p.patch(n)
	})
}

func (p *Patcher) Download(n Notification) {
	p.try(func() error {
		return p.download(n)
	})
}

func (p *Patcher) Remove(n Notification) {
	p.try(func() error {
		return p.remove(n)
	})
}
func (p *Patcher) Rename(n Notification) {
	p.try(func() error {
		return p.rename(n)
	})
}

func (p *Patcher) Start() {
	go func() {
		for {
			select {
			case n := <-p.Notifications:
				switch n.Event {
				case fsnotify.Create:
					p.Download(n)
				case fsnotify.Write:
					p.Patch(n)
				case fsnotify.Rename:
					p.Rename(n)
				case fsnotify.Remove:
					p.Remove(n)
				}

			}
		}
	}()
}
