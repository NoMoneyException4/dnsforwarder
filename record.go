package main

import (
	"github.com/miekg/dns"
)

//Record Cache record
type Record struct {
	Domain  string
	Type    uint16
	TTL     uint32
	Answers []dns.RR
}
