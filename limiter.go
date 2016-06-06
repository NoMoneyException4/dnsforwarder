package main

import (
	"net"

	"github.com/miekg/dns"
)

//Limiter Limiter interface
type Limiter interface {
	Limit(dns.ResponseWriter, *dns.Msg) bool
}

//WhiteListLimiter Limiter with white list
type WhiteListLimiter struct {
	lists []*net.IPNet
}

//NewWhiteListLimiter New WhiteListLimiter with conf
func NewWhiteListLimiter() *WhiteListLimiter {
	lists := make([]*net.IPNet)
	for _, str := range Conf.WhiteList {
		if _, nw, err := net.ParseCIDR(str); err == nil {
			lists = append(lists, nw)
		}
	}
	return &WhiteListLimiter{lists}
}

//Limit limit the request
func (l *WhiteListLimiter) Limit(w dns.ResponseWriter, req *dns.Msg) {
	for _, ipNet := range l.lists {
		if ipNet.Contains(w.RemoteAddr().(net.IPAddr).IP) {
			return true
		}
	}
	return false
}
