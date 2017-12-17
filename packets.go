package main

import (
	"bytes"
	"encoding/binary"
	"io"
	"os"

	"github.com/Redundancy/go-sync/chunks"
	"github.com/Redundancy/go-sync/filechecksum"
	"github.com/Redundancy/go-sync/index"
	"github.com/fsnotify/fsnotify"
)

// Hey is first message packet through UDP multicast
type Hey struct {
	Hostname string
	IP       string
	Items    []Item
}

// ItemType is type for tree item
type ItemType int

const (
	// File for file item
	File ItemType = iota
	// Dir for directory item
	Dir
)

// Item is directory tree struct
type Item struct {
	RelPath  string
	Checksum string
	ModTime  int64
}

// Notification is packet at track every file changes
type Notification struct {
	Hostname      string
	Event         fsnotify.Op
	Type          ItemType
	RelPath       string
	OldRelPath    string
	ModTime       int64
	IP            string
	ChecksumIndex []byte
}

func encodeChecksumIndex(content io.Reader, fileSize int64, blockSize uint) ([]byte, error) {
	generator := filechecksum.NewFileChecksumGenerator(uint(blockSize))
	weakSize := generator.WeakRollingHash.Size()
	strongSize := generator.GetStrongHash().Size()
	b := bytes.NewBuffer(nil)
	b.Write(int64ToBytes(fileSize))
	b.Write(intToBytes(weakSize))
	b.Write(intToBytes(strongSize))
	_, err := generator.GenerateChecksums(content, b)
	if err != nil {
		return nil, err
	}

	return b.Bytes(), nil
}

// EncodeChecksumIndex encode ChecksumIndex of gosync
func EncodeChecksumIndex(content io.Reader, fileSize int64, blockSize uint) (io.ReadSeeker, error) {
	b, err := encodeChecksumIndex(content, fileSize, blockSize)
	if err != nil {
		return nil, err
	}
	return bytes.NewReader(b), nil
}

// DecodeChecksumIndex decode EncodeChecksumIndex of gosync
func DecodeChecksumIndex(reader io.Reader) (fileSize int64, idx *index.ChecksumIndex, lookup filechecksum.ChecksumLookup, err error) {
	fBlock := bytes.NewBuffer(nil)
	wBlock := bytes.NewBuffer(nil)
	sBlock := bytes.NewBuffer(nil)

	_, err = io.CopyN(fBlock, reader, 8)
	if err != nil {
		return
	}

	_, err = io.CopyN(wBlock, reader, 4)
	if err != nil {
		return
	}

	_, err = io.CopyN(sBlock, reader, 4)
	if err != nil {
		return
	}

	fileSize = int64(binary.LittleEndian.Uint64(fBlock.Bytes()))
	weakSize := int(binary.LittleEndian.Uint32(wBlock.Bytes()))
	strongSize := int(binary.LittleEndian.Uint32(sBlock.Bytes()))

	readChunks, err := chunks.LoadChecksumsFromReader(reader, weakSize, strongSize)

	if err != nil {
		return
	}

	idx = index.MakeChecksumIndex(readChunks)
	lookup = chunks.StrongChecksumGetter(readChunks)

	return
}

func createNotification(evt Event, hostname string) (*Notification, error) {
	n := &Notification{
		Hostname:   hostname,
		Event:      evt.Op,
		Type:       File,
		RelPath:    evt.RelPath,
		OldRelPath: evt.OldRelPath,
	}
	file, err := os.Open(evt.FullPath)
	if err != nil {
		return n, nil
	}
	stat, err := os.Stat(evt.FullPath)
	if err != nil {
		return n, nil
	}
	if stat.IsDir() {
		n.Type = Dir
	}

	n.ChecksumIndex, err = encodeChecksumIndex(file, stat.Size(), BlockSize)
	if err != nil {
		return n, nil
	}

	return n, nil
}
