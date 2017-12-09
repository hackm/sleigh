package main

import (
	"encoding/json"
	"fmt"
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
	listener   *net.UDPConn
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
	fmt.Println(conn.ListenAddr)
	addr, err := net.ResolveUDPAddr("udp", conn.ListenAddr)
	if err != nil {
		conn.Errors <- err
	}
	conn.listener, err = net.ListenMulticastUDP("udp", nil, addr)
	if err != nil {
		conn.Errors <- err
	}
	conn.listener.SetReadBuffer(maxDatagramSize)
	go connDeamon(conn)
}

func connDeamon(conn *Conn) {
	for {
		b := make([]byte, maxDatagramSize)
		n, src, err := conn.listener.ReadFromUDP(b)
		fmt.Printf("%d\n", n)
		if err != nil {
			conn.Errors <- err
			continue
		}
		conn.Datagram <- Datagram{
			SrcAddr: *src,
			Payload: b[:n],
		}
	}
}

func (conn *Conn) Close() {
	conn.listener.Close()
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
