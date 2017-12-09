package main

import (
	"crypto/sha256"
	"encoding/hex"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/fsnotify/fsnotify"
)

// GetDirs get all directories under the `dir`
func GetDirs(dir string) []string {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}

	paths := []string{dir}
	for _, file := range files {
		name := file.Name()
		if file.IsDir() && strings.HasPrefix(name, ".") == false {
			paths = append(paths, GetDirs(filepath.Join(dir, name))...)
			continue
		}
	}

	return paths
}

// GetItems get all checksums in tree
func GetItems(root string, dir string, ignore IgnoreHandler) ([]Item, error) {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, err
	}
	items := []Item{}
	for _, file := range files {
		name := file.Name()
		path := filepath.Join(dir, name)
		evt, err := newEvent(root, path, fsnotify.Write)
		if err != nil {
			return nil, err
		}
		if ignore(evt) == false {
			if evt.Dir {
				children, err := GetItems(root, path, ignore)
				if err != nil {
					return nil, err
				}
				items = append(items, children...)
			} else {
				reader, err := os.OpenFile(path, os.O_RDONLY, 000)
				if err != nil {
					return nil, err
				}
				b, err := ioutil.ReadAll(reader)
				if err != nil {
					return nil, err
				}
				hash := sha256.Sum256(b)
				items = append(items, Item{
					RelPath:  evt.RelPath,
					Checksum: hex.EncodeToString(hash[:]),
					ModTime:  file.ModTime().UnixNano(),
				})
			}
		}
	}
	return items, err
}

// Event is struct for item change event
type Event struct {
	Op       fsnotify.Op
	FullPath string
	RelPath  string
	Parent   string
	Name     string
	Dir      bool
}

// IgnoreHandler is handler type for decision ignore file or directory
type IgnoreHandler func(Event) bool

// Tracker is watcher for items
type Tracker struct {
	root    string
	ignore  IgnoreHandler
	watcher *fsnotify.Watcher
	Events  chan (Event)
	Errors  chan (error)
}

// NewTracker is constructor for Tracker
func NewTracker(ignore IgnoreHandler) *Tracker {
	return &Tracker{
		ignore: ignore,
		Events: make(chan (Event)),
		Errors: make(chan (error)),
	}
}

// Start to watch items
func (t *Tracker) Start(root string) (err error) {
	if t.watcher == nil {
		if t.watcher, err = fsnotify.NewWatcher(); err != nil {
			return
		}
		if err = t.addDirs(root); err != nil {
			t.Close()
			return
		}
		t.root = root
		go track(t)
	}
	return
}

// Close to watch items
func (t *Tracker) Close() {
	if t.watcher != nil {
		t.watcher.Close()
	}
}

func (t *Tracker) addDirs(path string) error {
	targets := GetDirs(path)
	for _, dir := range targets {
		err := t.watcher.Add(dir)
		if err != nil {
			return err
		}
	}
	return nil
}

func track(t *Tracker) {
	for {
		select {
		case evt := <-t.watcher.Events:
			event, err := t.convert(evt)
			if err != nil {
				t.Errors <- err
			}
			if t.ignore(event) == false {
				if event.Dir && event.Op == fsnotify.Create {
					t.watcher.Add(event.FullPath)
				}
				t.Events <- event
			}
		case err := <-t.watcher.Errors:
			t.Errors <- err
		}
	}
}

func (t *Tracker) convert(evt fsnotify.Event) (event Event, err error) {
	event, err = newEvent(t.root, evt.Name, evt.Op)
	return
}

func newEvent(root, fullpath string, op fsnotify.Op) (event Event, err error) {
	event = Event{
		Op:       op,
		FullPath: fullpath,
	}
	// Get relative path from root
	event.RelPath, err = filepath.Rel(root, event.FullPath)
	if err != nil {
		return
	}
	// Get relative directory path from root
	event.Parent = filepath.Dir(event.RelPath)

	// get name of file or directory
	event.Name = filepath.Base(event.RelPath)
	if event.Name == "." {
		event.Name = strings.Replace(event.FullPath, filepath.Dir(event.FullPath), "", 1)[1:]
	}

	// when remove or rename, cannot read item stat.
	if event.Op == fsnotify.Remove || event.Op == fsnotify.Rename {
		return
	}

	// check item type: directory or file
	info, err := os.Stat(event.FullPath)
	if err != nil {
		return
	}
	event.Dir = info.IsDir()

	return
}
