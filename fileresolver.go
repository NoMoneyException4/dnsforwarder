package main

import (
	"bufio"
	"errors"
	"net"
	"os"
	"strings"

	"github.com/miekg/dns"
)

//FileResolver File based host resolver
type FileResolver struct {
	files []string
	hosts map[string][]string
	Host
}

//NewFileResolver New FileResolver
func NewFileResolver() *FileResolver {
	host := &FileResolver{
		files: Conf.Hosts.Resolvs,
		hosts: make(map[string][]string, 0),
	}
	host.Refresh()

	return host
}

//Get Get host from cache
func (host *FileResolver) Get(domain string) ([]string, error) {
	addrs, ok := host.hosts[domain]
	if ok {
		Logger.Debugf("[HitHost] Domain %s found in hosts files.", domain)
		return addrs, nil
	}
	return nil, errors.New("Not found.")
}

//Refresh Refresh the cached records
func (host *FileResolver) Refresh() {
	for _, file := range host.files {
		buf, err := os.OpenFile(file, os.O_RDONLY, 0777)
		if err != nil {
			Logger.Warningf("Update hosts records from file %s failed by %s.", file, err)
			return
		}
		defer buf.Close()

		scanner := bufio.NewScanner(buf)
		for scanner.Scan() {
			line := scanner.Text()
			line = strings.TrimSpace(line)
			if strings.HasPrefix(line, "#") || line == "" {
				continue
			}

			sli := strings.Split(line, " ")
			if len(sli) == 1 {
				sli = strings.Split(line, "\t")
			}

			if len(sli) < 2 {
				continue
			}

			domain := sli[len(sli)-1]
			ipString := sli[0]
			if !host.isDomain(domain) {
				Logger.Debugf("Cannot parse an invalid domain: `%s` from %s .", domain, file)
				continue
			}

			if !host.isIP(ipString) {
				Logger.Debugf("Cannot parse an invalid ip: `%s` from %s .", ipString, file)
				continue
			}

			domain = strings.ToLower(domain)
			domain = dns.Fqdn(domain)
			ip := net.ParseIP(ipString)
			if ip == nil {
				continue
			}

			host.hosts[domain] = append(host.hosts[domain], ipString)
		}
		Logger.Debugf("Cached hosts records from %s .", file)
	}
}
