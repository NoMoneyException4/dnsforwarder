package main

import (
	"github.com/miekg/dns"
)

type Record struct {
	Domain  string
	Ttl     int
	Message dns.Msg
}
