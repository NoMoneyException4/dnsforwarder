package main

import (
	"net"
)

//Record Query Record
type Record struct {
	Domain string
	Addrs  []net.IP
	TTL    int
}
