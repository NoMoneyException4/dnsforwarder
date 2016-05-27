package main

import (
	"net"
	"time"

	"github.com/miekg/dns"
)

type Server struct {
	Host         string
	Port         string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

func (s *Server) Addr() string {
	return net.JoinHostPort(s.Host, s.Port)
}

func (s *Server) Listen() {

}

func (s *Server) start(ds *dns.Server) {
	Logger.Info("Start %s server listening on %s", ds.Net, s.Addr())
	err := ds.ListenAndServe()
	if err != nil {
		Logger.Error("Start %s server failed:%s", ds.Net, err.Error())
	}
}
