package main

import (
	"os"
	"path"
	"path/filepath"
	"strings"
	"testing"

	"github.com/fsnotify/fsnotify"
)

func TestConvert_CreateFile(t *testing.T) {
	name := "tracker_test.go"
	wd, err := os.Getwd()
	if err != nil {
		t.Fatal("cannot get working directory")
	}
	file := path.Join(wd, name)
	tr := &Tracker{
		root: wd,
	}

	evt, err := tr.convert(fsnotify.Event{
		Op:   fsnotify.Create,
		Name: file,
	})

	if err != nil {
		t.Fatalf("cannot convert event: %v", err)
	}
	if evt.FullPath != file {
		t.Error("FullPath is invalid")
	}
	if evt.RelPath != name {
		t.Errorf("RelPath is invalid: %v", evt.RelPath)
	}
	if evt.Name != name {
		t.Errorf("Name is invalid: %v", evt.Name)
	}
	if evt.Dir != false {
		t.Error("Dir flag is invalid")
	}
	if evt.Parent != "." {
		t.Errorf("Parent is invalid: %v", evt.Parent)
	}
}

func TestConvert_CreateDir(t *testing.T) {
	name := "vendor"
	wd, err := os.Getwd()
	if err != nil {
		t.Fatal("cannot get working directory")
	}
	dir := path.Join(wd, name)
	tr := &Tracker{
		root: wd,
	}

	evt, err := tr.convert(fsnotify.Event{
		Op:   fsnotify.Create,
		Name: dir,
	})

	if err != nil {
		t.Fatalf("cannot convert event: %v", err)
	}
	if evt.FullPath != dir {
		t.Error("FullPath is invalid")
	}
	if evt.RelPath != name {
		t.Errorf("RelPath is invalid: %v", evt.RelPath)
	}
	if evt.Name != name {
		t.Errorf("Name is invalid: %v", evt.Name)
	}
	if evt.Dir != true {
		t.Error("Dir is invalid")
	}
	if evt.Parent != "." {
		t.Errorf("Parent is invalid: %v", evt.Parent)
	}
}

func TestConvert_RemoveFile(t *testing.T) {
	name := "removed.go"
	wd, err := os.Getwd()
	if err != nil {
		t.Fatal("cannot get working directory")
	}
	file := path.Join(wd, name)
	tr := &Tracker{
		root: wd,
	}

	evt, err := tr.convert(fsnotify.Event{
		Op:   fsnotify.Remove,
		Name: file,
	})

	if err != nil {
		t.Fatalf("cannot convert event: %v", err)
	}
	if evt.FullPath != file {
		t.Error("FullPath is invalid")
	}
	if evt.RelPath != name {
		t.Errorf("RelPath is invalid: %v", evt.RelPath)
	}
	if evt.Name != name {
		t.Errorf("Name is invalid: %v", evt.Name)
	}
	if evt.Parent != "." {
		t.Errorf("Parent is invalid: %v", evt.Parent)
	}
}

func TestGetItems(t *testing.T) {
	wd, err := os.Getwd()
	if err != nil {
		t.Fatal("cannot get working directory")
	}
	items, err := GetItems(wd, wd, func(evt Event) bool { return strings.Contains(evt.RelPath, "_test.go") })
	if err != nil {
		t.Fatal("cannot get items")
	}
	_, err = os.Stat(filepath.Join(wd, items[0].RelPath))
	if err != nil {
		t.Errorf("RelPath is invalid: %v", items[0].RelPath)
	}
	if items[0].Checksum == "" {
		t.Error("Checksum is not set")
	}
	if items[0].ModTime <= 0 {
		t.Errorf("ModTime is not set %v", items[0].ModTime)
	}
}
