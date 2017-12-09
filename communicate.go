package main

import (
	"encoding/hex"
	"encoding/json"
	"log"
	"net"
)

const (
	multicastAddr   = "224.0.0.1:8000"
	maxDatagramSize = 8192
)

type Communicate interface {
	Hey() int
	Notify(string) int
	ListenNotification()
	SyncDiff()
	Parse(*interface{}) error
}

type Datagram []byte

func (d Datagram) Hey() int {
	n, err := sendUDPMulticast(d)
	if err != nil {
		log.Fatal(err)
	}
	return n
}

func (d Datagram) Notify() int {
	n, err := sendUDPMulticast(d)
	if err != nil {
		log.Fatal(err)
	}
	return n
}

func notificationHandler(src *net.UDPAddr, n int, b []byte) {
	// TODO: convert []byte to Notification
	log.Println(n, "bytes read from", src)
	log.Println(hex.Dump(b[:n]))
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

func sendUDPMulticast(d Datagram) (int, error) {
	addr, err := net.ResolveUDPAddr("udp", multicastAddr)
	if err != nil {
		return 0, err
	}
	c, err := net.DialUDP("udp", nil, addr)
	defer c.Close()
	n, err := c.Write([]byte(d))
	if err != nil {
		return 0, err
	}
	return n, nil
}

func (d Datagram) sendHTTP(dest string) (int, error) {
	return 0, nil
}

func (d Datagram) Parse(v *interface{}) error {
	err := json.Unmarshal([]byte(d), &v)
	if err != nil {
		return err
	}
	return nil
}
