package main

import (
	"encoding/json"
	"net"
)

const (
	multicastAddr   = "224.0.0.1"
	maxDatagramSize = 8192
)

// Datagram is received by Conn listener
type Datagram struct {
	SrcAddr net.UDPAddr
	Payload []byte
}

// Conn is process multicast deamon listener
type Conn struct {
	ListenAddr string
	Datagram   chan (Datagram)
	Errors     chan (error)
}

// NewConn is constructor for Conn
func NewConn(listenPort string) *Conn {
	return &Conn{
		ListenAddr: multicastAddr + ":" + listenPort,
		Datagram:   make(chan (Datagram)),
		Errors:     make(chan (error)),
	}
}

// Listen to receive Datagram
func (conn *Conn) Listen() {
	addr, err := net.ResolveUDPAddr("udp", conn.ListenAddr)
	if err != nil {
		conn.Errors <- err
	}
	l, err := net.ListenMulticastUDP("udp", nil, addr)
	defer l.Close()
	if err != nil {
		conn.Errors <- err
	}
	l.SetReadBuffer(maxDatagramSize)
	for {
		b := make([]byte, maxDatagramSize)
		n, src, err := l.ReadFromUDP(b)
		if err != nil {
			conn.Errors <- err
		}
		conn.Datagram <- Datagram{
			SrcAddr: *src,
			Payload: b[:n],
		}
	}
}

// Hey is sending Hey
func (conn *Conn) Hey(hey Hey) (err error) {
	b, err := json.Marshal(hey)
	if err != nil {
		return
	}
	_, err = conn.sendUDPMulticast(b)
	if err != nil {
		return
	}
	return
}

// Notify is sending Notification
func (conn *Conn) Notify(notification Notification) (err error) {
	b, err := json.Marshal(notification)
	if err != nil {
		return
	}
	_, err = conn.sendUDPMulticast(b)
	if err != nil {
		return
	}
	return
}

func (conn *Conn) sendUDPMulticast(b []byte) (int, error) {
	addr, err := net.ResolveUDPAddr("udp", conn.ListenAddr)
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
