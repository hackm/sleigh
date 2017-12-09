package main

import (
	"bytes"
	"encoding/binary"
	"io"

	"github.com/Redundancy/go-sync/chunks"
	"github.com/Redundancy/go-sync/filechecksum"
	"github.com/Redundancy/go-sync/index"
)

// Hey is first message packet through UDP multicast
type Hey struct {
	Hostname string `json:"hostname"`
	Tree     []Item `json:"tree"`
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
	Type ItemType `json:"type"`
	Name string   `json:"name"`
	Tree []Item   `json:"tree"`
}

// Notification is packet for notify diff
type Notification struct {
	Hostname  string   `json:"hostname"`
	Event     string   `json:"event"`
	Type      ItemType `json:"type"`
	Path      string   `json:"path"`
	Timestamp int64    `json:"timestamp"`
	Dst       string   `json:"dst"`
}

// Event for file change
type Event int

const (
	// Create file|dir
	Create Event = iota
	// Write file
	Write
	// Rename file|dir
	Rename
	// Delete file|dir
	Delete
)

// EncodeChecksumIndex encode ChecksumIndex of gosync
func EncodeChecksumIndex(content io.Reader, fileSize int64, blockSize uint) (io.ReadSeeker, error) {
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

	return bytes.NewReader(b.Bytes()), nil
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

// helpers
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
