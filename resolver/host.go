package resolver

import (
	"net"
	"regexp"
)

type Host struct {
}

func (host *Host) isDomain(domain string) bool {
	if host.isIP(domain) {
		return false
	}
	match, _ := regexp.MatchString(`^([a-zA-Z0-9\*]([a-zA-Z0-9\-]{0,61}[a-zA-Z0-9])?\.)+[a-zA-Z]{2,6}$`, domain)
	return match
}

func (host *Host) isIP(ip string) bool {
	return (net.ParseIP(ip) != nil)
}
