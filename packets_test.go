package main

import (
	"encoding/json"
	"strings"
	"testing"
)

func TestItemSerialize(t *testing.T) {
	item := Item{}
	b, err := json.Marshal(item)
	if err != nil {
		t.Errorf("cannot serialize 'Item': %v", err)
	}
	str := string(b)
	if strings.Contains(str, "type") == false {
		t.Errorf("type not found")
	}
	if strings.Contains(str, "name") == false {
		t.Errorf("name not found")
	}
	if strings.Contains(str, "checksum") == false {
		t.Errorf("checksum not found")
	}
	if strings.Contains(str, "modtime") == false {
		t.Errorf("modtime not found")
	}
}
func TestHeySerialize(t *testing.T) {
	hey := Hey{}
	b, err := json.Marshal(hey)
	if err != nil {
		t.Errorf("cannot serialize 'Hey' packet: %v", err)
	}
	str := string(b)
	if strings.Contains(str, `"hostname"`) == false {
		t.Errorf("hotname not found")
	}
	if strings.Contains(str, `"tree"`) == false {
		t.Errorf("tree not found")
	}
}

func TestNotificationSerialize(t *testing.T) {
	n := Notification{}
	b, err := json.Marshal(n)
	if err != nil {
		t.Errorf("cannot serialize 'Notification' packet: %v", err)
	}
	str := string(b)
	if strings.Contains(str, `"hostname"`) == false {
		t.Errorf("hotname not found")
	}
	if strings.Contains(str, `"dst"`) == false {
		t.Errorf("dst not found")
	}
	if strings.Contains(str, `"event"`) == false {
		t.Errorf("event not found")
	}
	if strings.Contains(str, `"type"`) == false {
		t.Errorf("type not found")
	}
	if strings.Contains(str, `"path"`) == false {
		t.Errorf("path not found")
	}
	if strings.Contains(str, `"timestamp"`) == false {
		t.Errorf("timestamp not found")
	}
}
