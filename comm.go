package main

import (
	"encoding/hex"
	"encoding/json"
	"log"
	"net"
)

const (
	multicastAddr   = "224.0.0.1:8986"
	maxDatagramSize = 8192
)

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
	addr, err := net.ResolveUDPAddr("udp", c.listenAddr)
	if err != nil {
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
	log.Println(n, "bytes read from", src)
	log.Println(hex.Dump(b[:n]))
}

func (c *Comm) Close() {
	return
}

func (hey Hey) Send() {
	var err error
	b, err := json.Marshal(hey)
	if err != nil {
		log.Fatal(err)
	}
	_, err = sendUDPMulticast(b)
	if err != nil {
		log.Fatal(err)
	}
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
