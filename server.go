package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"

	"github.com/fatih/color"
)

// Server is to handle own file system changes and broadcast a notification to another notifier
type Server struct {
	Hostname         string
	Notifications    chan Notification
	Heys             chan Hey
	Errors           chan error
	localPathResolve LocalPathResolver
	port             int
	conn             *net.UDPConn
	others           map[string]time.Time
}

// NewServer is constructor for creating server
func NewServer(hostname string, port int, resolver LocalPathResolver) *Server {
	return &Server{
		Hostname:         hostname,
		Heys:             make(chan Hey),
		Notifications:    make(chan Notification),
		Errors:           make(chan error),
		localPathResolve: resolver,
		port:             port,
		others:           make(map[string]time.Time),
	}
}

// Start Server
func (s *Server) Start() {
	s.http()
	s.udp()
	s.heartbeat()
}

func (s *Server) http() {
	r := mux.NewRouter()
	r.HandleFunc("/contents", func(w http.ResponseWriter, req *http.Request) {
		path := s.localPathResolve(req.URL.Query().Get("path"))
		if _, err := os.Stat(path); err != nil {
			http.NotFound(w, req)
			return
		}
		http.ServeFile(w, req, path)
	}).Methods("GET")

	r.HandleFunc("/notifications", func(w http.ResponseWriter, req *http.Request) {
		b, err := ioutil.ReadAll(req.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(fmt.Sprintf("Cannot read request body: %s", err)))
			return
		}
		var n Notification
		err = json.Unmarshal(b, &n)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(fmt.Sprintf("Cannot parse request body as JSON: %s", err)))
			return
		}
		ip, _, _ := net.SplitHostPort(req.RemoteAddr)
		n.IP = ip
		s.Notifications <- n
		w.WriteHeader(http.StatusCreated)
	}).Methods("POST")

	go func() {
		err := http.ListenAndServe(fmt.Sprintf(":%d", s.port), r)
		if err != nil {
			color.Yellow("cannot listen http server: %v\n", err)
		}
	}()
}

func (s *Server) udp() {
	addr, err := net.ResolveUDPAddr("udp", fmt.Sprintf("%s:%d", MulticastAddr, s.port))
	if err != nil {
		s.Errors <- err
	}
	s.conn, err = net.ListenMulticastUDP("udp", nil, addr)
	if err != nil {
		s.Errors <- err
	}
	s.conn.SetReadBuffer(MaxDatagramSize)
	go func() {
		for {
			b := make([]byte, MaxDatagramSize)
			n, src, err := s.conn.ReadFromUDP(b)
			if err != nil {
				s.Errors <- err
				continue
			}
			var hey Hey
			err = json.Unmarshal(b[:n], &hey)
			if err != nil {
				s.Errors <- err
			}
			if hey.Hostname == s.Hostname {
				continue
			}
			ip := src.IP.String()
			if _, ok := s.others[ip]; ok == false {
				hey.IP = ip
				s.Heys <- hey
			}
			s.others[ip] = time.Now()
		}
	}()
}

func (s *Server) heartbeat() {
	hey := Hey{
		Hostname: s.Hostname,
	}
	go func() {
		for {
			log.Print(".")
			b, err := json.Marshal(hey)
			if err != nil {
				s.Errors <- err
				continue
			}
			_, err = Multicast(s.port, b)
			if err != nil {
				s.Errors <- err
			}
			time.Sleep(10 * time.Second)
		}
	}()
}

// Others manages another notifier
func (s *Server) Others() []string {
	pivot := time.Now().Add(-15 * time.Second)
	ret := make([]string, 0, len(s.others))
	for ip, last := range s.others {
		if last.After(pivot) {
			ret = append(ret, ip)
		}
	}
	return ret
}

// Close is called at exiting this application
func (s *Server) Close() {
	s.conn.Close()
}
