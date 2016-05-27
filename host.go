package main

import (
	"net"
	"regexp"
)

//Host Base Host
type Host struct {
}

//isDomain Determine the given host is valid
func (host *Host) isDomain(domain string) bool {
	if host.isIP(domain) {
		return false
	}
	match, _ := regexp.MatchString(`^([a-zA-Z0-9\*]([a-zA-Z0-9\-]{0,61}[a-zA-Z0-9])?\.)+[a-zA-Z]{2,6}$`, domain)
	return match
}

//isIP Determine the given ip is valid
func (host *Host) isIP(ip string) bool {
	return (net.ParseIP(ip) != nil)
}
