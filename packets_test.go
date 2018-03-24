package main

import (
	"encoding/json"
	"io"
	"os"
	"reflect"
	"strings"
	"testing"

	"github.com/Redundancy/go-sync/filechecksum"
	"github.com/Redundancy/go-sync/index"
)

func TestItemSerialize(t *testing.T) {
	item := Item{}
	b, err := json.Marshal(item)
	if err != nil {
		t.Fatalf("cannot serialize 'Item': %v", err)
	}
	str := string(b)
	if strings.Contains(str, "RelPath") == false {
		t.Errorf("RelPath not found")
	}
	if strings.Contains(str, "Checksum") == false {
		t.Errorf("Checksum not found")
	}
	if strings.Contains(str, "ModTime") == false {
		t.Errorf("Modtime not found")
	}
}
func TestHeySerialize(t *testing.T) {
	hey := Hey{}
	b, err := json.Marshal(hey)
	if err != nil {
		t.Fatalf("cannot serialize 'Hey' packet: %v", err)
	}
	str := string(b)
	if strings.Contains(str, "Hostname") == false {
		t.Errorf("Hotname not found")
	}
	if strings.Contains(str, "Items") == false {
		t.Errorf("Imtes not found")
	}
}

func TestNotificationSerialize(t *testing.T) {
	n := Notification{}
	b, err := json.Marshal(n)
	if err != nil {
		t.Fatalf("cannot serialize 'Notification' packet: %v", err)
	}
	str := string(b)
	if strings.Contains(str, "Hostname") == false {
		t.Errorf("Hotname not found")
	}
	if strings.Contains(str, "IP") == false {
		t.Errorf("IP not found")
	}
	if strings.Contains(str, "Event") == false {
		t.Errorf("Event not found")
	}
	if strings.Contains(str, "Type") == false {
		t.Errorf("Type not found")
	}
	if strings.Contains(str, "Path") == false {
		t.Errorf("Path not found")
	}
	if strings.Contains(str, "ModTime") == false {
		t.Errorf("ModTime not found")
	}
}

func TestEncodeDecodeChecksumIndex(t *testing.T) {
	cnt, err := os.OpenFile("packets_test.go", os.O_RDONLY, 0)
	if err != nil {
		t.Fatal("open test content file")
	}
	info, err := cnt.Stat()
	if err != nil {
		t.Fatal("cannot get stat of test content file")
	}
	seaker, err := EncodeChecksumIndex(cnt, info.Size(), 1024)
	if err != nil {
		t.Fatalf("cannot encode: %v", err)
	}
	size, idx, checksum, err := DecodeChecksumIndex(seaker)
	if err != nil {
		t.Fatalf("cannot decode: %v", err)
	}
	if size != info.Size() {
		t.Errorf("file size difference between src and decoded")
	}
	if idx == nil {
		t.Errorf("idx is nil")
	}
	if checksum == nil {
		t.Errorf("checksum is nil")
	}
}

func Test_encodeChecksumIndex(t *testing.T) {
	type args struct {
		content   io.Reader
		fileSize  int64
		blockSize uint
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := encodeChecksumIndex(tt.args.content, tt.args.fileSize, tt.args.blockSize)
			if (err != nil) != tt.wantErr {
				t.Errorf("encodeChecksumIndex() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("encodeChecksumIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEncodeChecksumIndex(t *testing.T) {
	type args struct {
		content   io.Reader
		fileSize  int64
		blockSize uint
	}
	tests := []struct {
		name    string
		args    args
		want    io.ReadSeeker
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := EncodeChecksumIndex(tt.args.content, tt.args.fileSize, tt.args.blockSize)
			if (err != nil) != tt.wantErr {
				t.Errorf("EncodeChecksumIndex() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("EncodeChecksumIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDecodeChecksumIndex(t *testing.T) {
	type args struct {
		reader io.Reader
	}
	tests := []struct {
		name         string
		args         args
		wantFileSize int64
		wantIdx      *index.ChecksumIndex
		wantLookup   filechecksum.ChecksumLookup
		wantErr      bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotFileSize, gotIdx, gotLookup, err := DecodeChecksumIndex(tt.args.reader)
			if (err != nil) != tt.wantErr {
				t.Errorf("DecodeChecksumIndex() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotFileSize != tt.wantFileSize {
				t.Errorf("DecodeChecksumIndex() gotFileSize = %v, want %v", gotFileSize, tt.wantFileSize)
			}
			if !reflect.DeepEqual(gotIdx, tt.wantIdx) {
				t.Errorf("DecodeChecksumIndex() gotIdx = %v, want %v", gotIdx, tt.wantIdx)
			}
			if !reflect.DeepEqual(gotLookup, tt.wantLookup) {
				t.Errorf("DecodeChecksumIndex() gotLookup = %v, want %v", gotLookup, tt.wantLookup)
			}
		})
	}
}

func Test_createNotification(t *testing.T) {
	type args struct {
		evt      Event
		hostname string
	}
	tests := []struct {
		name    string
		args    args
		want    *Notification
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := createNotification(tt.args.evt, tt.args.hostname)
			if (err != nil) != tt.wantErr {
				t.Errorf("createNotification() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("createNotification() = %v, want %v", got, tt.want)
			}
		})
	}
}
