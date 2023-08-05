package main

import (
	"log"
	"net"
	"time"

	"github.com/miekg/dns"
)

// Server Server
type Server struct {
	Host         string
	Port         string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

// NewServer New Server
func NewServer(host, port string, rt, wt time.Duration) *Server {
	return &Server{
		Host:         host,
		Port:         port,
		ReadTimeout:  rt,
		WriteTimeout: wt,
	}
}

// Addr Return the addr that server is listening at
func (s *Server) Addr() string {
	return net.JoinHostPort(s.Host, s.Port)
}

// Listen Server start listen
func (s *Server) Listen() {
	handler := NewHandler()

	tcpHandler := dns.NewServeMux()
	tcpHandler.HandleFunc(".", handler.HandleTCP)

	udpHandler := dns.NewServeMux()
	udpHandler.HandleFunc(".", handler.HandleUDP)

	tcpServer := &dns.Server{
		Addr:         s.Addr(),
		Net:          "tcp",
		Handler:      tcpHandler,
		ReadTimeout:  s.ReadTimeout,
		WriteTimeout: s.WriteTimeout,
	}

	udpServer := &dns.Server{
		Addr:         s.Addr(),
		Net:          "udp",
		Handler:      udpHandler,
		UDPSize:      65535,
		ReadTimeout:  s.ReadTimeout,
		WriteTimeout: s.WriteTimeout,
	}

	go s.start(udpServer)
	go s.start(tcpServer)
}

func (s *Server) start(ds *dns.Server) {
	Logger.Infof("Start %s server listening on %s", ds.Net, s.Addr())
	err := ds.ListenAndServe()
	if err != nil {
		log.Fatalf("Start %s server failed:%s", ds.Net, err.Error())
	}
}
