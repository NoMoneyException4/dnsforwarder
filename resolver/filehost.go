package resolver

import (
	"bufio"
	"errors"
	"net"
	"os"
	"strings"

	. "github.com/codebear4/dnsforwarder/conf"
	. "github.com/codebear4/dnsforwarder/logger"
)

type FileHost struct {
	files []string
	hosts map[string][]net.IP
	Host
}

func NewFileHost() *FileHost {
	host := &FileHost{
		files: Conf.Hosts.Resolvs,
		hosts: make(map[string][]net.IP, 0),
	}
	host.Refresh()

	return host
}

func (host *FileHost) Get(domain string) (error, *Record) {
	addrs, ok := host.hosts[domain]
	if ok {
		return nil, &Record{Domain: domain, Ttl: 0, Addrs: addrs}
	}
	Logger.Debugf("Domain %s not found in hosts files.", domain)
	return errors.New("Not found."), nil
}

func (host *FileHost) Set(domain string, record *Record) error {
	host.hosts[domain] = append(host.hosts[domain], record.Addrs...)
	return nil
}

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
