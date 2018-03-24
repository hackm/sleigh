package main

import (
	"reflect"
	"testing"
)

func TestNewServer(t *testing.T) {
	type args struct {
		hostname string
		port     int
		resolver LocalPathResolver
	}
	tests := []struct {
		name string
		args args
		want *Server
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewServer(tt.args.hostname, tt.args.port, tt.args.resolver); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewServer() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestServer_Start(t *testing.T) {
	tests := []struct {
		name string
		s    *Server
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.s.Start()
		})
	}
}

func TestServer_http(t *testing.T) {
	tests := []struct {
		name string
		s    *Server
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.s.http()
		})
	}
}

func TestServer_udp(t *testing.T) {
	tests := []struct {
		name string
		s    *Server
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.s.udp()
		})
	}
}

func TestServer_heartbeat(t *testing.T) {
	tests := []struct {
		name string
		s    *Server
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.s.heartbeat()
		})
	}
}

func TestServer_Others(t *testing.T) {
	tests := []struct {
		name string
		s    *Server
		want []string
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.Others(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Server.Others() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestServer_Close(t *testing.T) {
	tests := []struct {
		name string
		s    *Server
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.s.Close()
		})
	}
}
