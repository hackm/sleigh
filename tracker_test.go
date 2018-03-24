package main

import (
	"os"
	"path"
	"path/filepath"
	"reflect"
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
	}, nil)

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
	}, nil)

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
	}, nil)

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

func TestGetDirs(t *testing.T) {
	type args struct {
		dir string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetDirs(tt.args.dir); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetDirs() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getChecksum(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getChecksum(tt.args.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("getChecksum() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("getChecksum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewTracker(t *testing.T) {
	type args struct {
		root   string
		ignore IgnoreHandler
	}
	tests := []struct {
		name string
		args args
		want *Tracker
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewTracker(tt.args.root, tt.args.ignore); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewTracker() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTracker_Start(t *testing.T) {
	tests := []struct {
		name    string
		t       *Tracker
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.t.Start(); (err != nil) != tt.wantErr {
				t.Errorf("Tracker.Start() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestTracker_Close(t *testing.T) {
	tests := []struct {
		name string
		t    *Tracker
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.t.Close()
		})
	}
}

func TestTracker_addDirs(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name    string
		t       *Tracker
		args    args
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.t.addDirs(tt.args.path); (err != nil) != tt.wantErr {
				t.Errorf("Tracker.addDirs() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_track(t *testing.T) {
	type args struct {
		t *Tracker
	}
	tests := []struct {
		name string
		args args
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			track(tt.args.t)
		})
	}
}

func TestTracker_convert(t *testing.T) {
	type args struct {
		evt    fsnotify.Event
		rename *fsnotify.Event
	}
	tests := []struct {
		name      string
		t         *Tracker
		args      args
		wantEvent Event
		wantErr   bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotEvent, err := tt.t.convert(tt.args.evt, tt.args.rename)
			if (err != nil) != tt.wantErr {
				t.Errorf("Tracker.convert() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotEvent, tt.wantEvent) {
				t.Errorf("Tracker.convert() = %v, want %v", gotEvent, tt.wantEvent)
			}
		})
	}
}

func Test_newEvent(t *testing.T) {
	type args struct {
		root     string
		fullpath string
		op       fsnotify.Op
	}
	tests := []struct {
		name      string
		args      args
		wantEvent Event
		wantErr   bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotEvent, err := newEvent(tt.args.root, tt.args.fullpath, tt.args.op)
			if (err != nil) != tt.wantErr {
				t.Errorf("newEvent() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotEvent, tt.wantEvent) {
				t.Errorf("newEvent() = %v, want %v", gotEvent, tt.wantEvent)
			}
		})
	}
}
