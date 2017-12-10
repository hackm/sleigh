package main

import (
	"net"
	"regexp"
	"testing"
)

func initializeConn() (conn *Conn) {
	defaultPort := uint(8986)
	conn = NewConn(defaultPort)
	return
}

func TestConnSerialize(t *testing.T) {
	conn := initializeConn()
	re := regexp.MustCompile("[0-9]{1-3}\\.[0-9]{1-3}\\.[0-9]{1-3}\\.[0-9]{1-3}:[0-9]{1-4}")
	go func() {
		if re.Match([]byte(conn.ListenAddr)) {
			t.Error("ListenAddr is invalid")
		}
		srcAddr := net.UDPAddr{}
		conn.Datagram <- Datagram{
			SrcAddr: srcAddr,
			Payload: []byte("hello world\n"),
		}
	}()
	<-conn.Datagram
}

func TestConnHey(t *testing.T) {
	conn := initializeConn()
	hey := Hey{}
	err := conn.Hey(hey)
	if err != nil {
		t.Error("network error")
	}
}

func TestConnNotify(t *testing.T) {
	conn := initializeConn()
	notification := Notification{}
	err := conn.Notify(notification)
	if err != nil {
		t.Error("network error")
	}
}
