package main

import (
	"encoding/hex"
	"encoding/json"
	"log"
	"net"
	"os"
	"path/filepath"
	"strings"

	"github.com/fsnotify/fsnotify"
)

const (
	multicastAddr   = "224.0.0.1:8986"
	maxDatagramSize = 8192
)

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

// Comm is notifier deamon
type Comm struct {
	listenAddr string
	Events     chan (Event)
	Errors     chan (error)
}

// NewComm is constructor for Comm
func NewComm() *Comm {
	return &Comm{
		listenAddr: multicastAddr,
		Events:     make(chan (Event)),
		Errors:     make(chan (error)),
	}
}

// Start to listen Hey and Diff Event
func (c *Comm) Start() (err error) {
	if addr, err = net.ResolveUDPAddr("udp", a); err != nil {
		return
	}
	l, err := net.ListenMulticastUDP("udp", nil, addr)
	defer l.Close()
	if err != nil {
		return
	}
	l.SetReadBuffer(maxDatagramSize)
	for {
		b := make([]byte, maxDatagramSize)
		n, src, err := l.ReadFromUDP(b)
		if err != nil {
			log.Fatal("ReadFromUDP failed:", err)
		}
		commHandler(src, n, b)
	}
}

func commHandler(src *net.UDPAddr, n int, b []byte) {
	// deamon should listen Hey to return notification
	// deamon should listen Notification to check diff
	// deamon should listen File event to notify
	json.Unmarshal()
	log.Println(n, "bytes read from", src)
	log.Println(hex.Dump(b[:n]))
}

func (c *Comm) Close() error {

}

func (h Hey) Send() int {
	n, err := sendUDPMulticast(h)
	if err != nil {
		log.Fatal(err)
	}
	return n
}

// Start to watch items
func (t *Tracker) Start() (err error) {
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

func (d Datagram) Notify() int {
	n, err := sendUDPMulticast(d)
	if err != nil {
		log.Fatal(err)
	}
	return n
}

func ListenNotification(a string, h func(*net.UDPAddr, int, []byte)) {
	addr, err := net.ResolveUDPAddr("udp", a)
	if err != nil {
		log.Fatal(err)
	}
	l, err := net.ListenMulticastUDP("udp", nil, addr)
	l.SetReadBuffer(maxDatagramSize)
	for {
		b := make([]byte, maxDatagramSize)
		n, src, err := l.ReadFromUDP(b)
		if err != nil {
			log.Fatal("ReadFromUDP failed:", err)
		}
		h(src, n, b)
	}
}

func sendUDPMulticast(b []byte) (int, error) {
	addr, err := net.ResolveUDPAddr("udp", multicastAddr)
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
