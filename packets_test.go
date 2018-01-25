package main

import (
	"encoding/json"
	"os"
	"strings"
	"testing"
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
