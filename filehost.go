package main

import (
	"bufio"
	"errors"
	"net"
	"os"
	"strings"
)

//FileHost File based host resolver
type FileHost struct {
	files []string
	hosts map[string][]net.IP
	Host
}

//NewFileHost New FileHost
func NewFileHost() *FileHost {
	host := &FileHost{
		files: Conf.Hosts.Resolvs,
		hosts: make(map[string][]net.IP, 0),
	}
	host.Refresh()

	return host
}

//Get Get host from cache
func (host *FileHost) Get(domain string) (*Record, error) {
	addrs, ok := host.hosts[domain]
	if ok {
		return &Record{Domain: domain, Ttl: 0, Addrs: addrs}, nil
	}
	Logger.Debugf("Domain %s not found in hosts files.", domain)
	return nil, errors.New("Not found.")
}

//Set Set host with given Record
func (host *FileHost) Set(domain string, record *Record) error {
	host.hosts[domain] = append(host.hosts[domain], record.Addrs...)
	return nil
}

//All Get all hosts
func (host *FileHost) All() []*Record {
	all := make([]*Record, len(host.hosts))
	for key, _ := range host.hosts {
		record, err := host.Get(key)
		if err == nil {
			all = append(all, record)
		}
	}
	return all
}

//Clear Clear all hosts
func (host *FileHost) Clear() {
	host.hosts = make(map[string][]net.IP, 0)
}

//Refresh Refresh the cached records
func (host *FileHost) Refresh() {
	for _, file := range host.files {
		buf, err := os.OpenFile(file, os.O_RDONLY, 0777)
		if err != nil {
			Logger.Warningf("Update hosts records from file %s failed %s.", file, err)
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
				Logger.Debugf("Cannot parse a invalid domain: `%s` from %s.", domain, file)
				continue
			}

			if !host.isIP(ipString) {
				Logger.Debugf("Cannot parse a  invalid ip: `%s` from %s", ipString, file)
				continue
			}

			domain = strings.ToLower(domain)
			ip := net.ParseIP(ipString)
			if ip == nil {
				continue
			}

			host.hosts[domain] = append(host.hosts[domain], ip)
		}
		Logger.Debugf("Updated hosts records from %s successful.", file)
	}
}
