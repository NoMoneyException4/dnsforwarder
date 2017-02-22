package main

import (
	"errors"
	"net"

	"github.com/miekg/dns"
)

//Limiter Limiter interface
type Limiter interface {
	Limit(dns.ResponseWriter, *dns.Msg) bool
	RemoteIP(dns.ResponseWriter) (net.IP, error)
}

//WhiteListLimiter Limiter with white list
type WhiteListLimiter struct {
	lists []*net.IPNet
}

//NewWhiteListLimiter New WhiteListLimiter with conf
func NewWhiteListLimiter() *WhiteListLimiter {
	lists := make([]*net.IPNet, 0)
	for _, str := range Conf.WhiteList {
		if _, nw, err := net.ParseCIDR(str); err == nil {
			lists = append(lists, nw)
		}
	}
	return &WhiteListLimiter{lists}
}

//Limit limit the request
func (l *WhiteListLimiter) Limit(w dns.ResponseWriter, req *dns.Msg) bool {
	ip, err := l.RemoteIP(w)
	if err != nil {
		return false
	}

	for _, ipNet := range l.lists {
		if ipNet.Contains(ip) {
			return true
		}
	}
	return false
}

RemoteIP get remote ip from response writer
func (l *WhiteListLimiter) RemoteIP(w dns.ResponseWriter) (net.IP, error) {
	switch w.RemoteAddr().(type) {
	case *net.UDPAddr:
		return w.RemoteAddr().(*net.UDPAddr).IP, nil
	case *net.TCPAddr:
		return w.RemoteAddr().(*net.TCPAddr).IP, nil
	default:
		return nil, errors.New("Not a valid IP address: " + w.RemoteAddr().String())
	}
}
