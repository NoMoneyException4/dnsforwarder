package main

import (
	"net"
	"strconv"
	"testing"

	"github.com/miekg/dns"
	"github.com/stretchr/testify/assert"
)

func TestCacheHostGetSet(t *testing.T) {
	InitLogger()
	domain := dns.Fqdn("fake.codebear.xyz")
	expectedIP := "8.8.8.8"
	cacheHost := NewCacheHost()
	record := &Record{
		Domain: domain,
		Addrs:  []net.IP{net.ParseIP(expectedIP)},
		TTL:    0,
	}
	cacheHost.Set(domain, record)
	record, err := cacheHost.Get(domain)
	assert.Nil(t, err)
	assert.Equal(t, 1, len(record.Addrs))
	assert.Equal(t, expectedIP, record.Addrs[0].String())
}

func TestCacheHostAll(t *testing.T) {
	InitLogger()
	expectedIP := "8.8.8.8"
	cacheHost := NewCacheHost()

	for i := 0; i < 5; i++ {
		domain := dns.Fqdn("fake" + strconv.Itoa(i) + ".codebear.xyz")
		record := &Record{
			Domain: domain,
			Addrs:  []net.IP{net.ParseIP(expectedIP)},
			TTL:    0,
		}
		cacheHost.Set(domain, record)
	}
	assert.True(t, len(cacheHost.All()) > 0, "Hosts must not be empty")
}

func TestCacheHostClear(t *testing.T) {
	InitLogger()
	domain := dns.Fqdn("fake.codebear.xyz")
	expectedIP := "8.8.8.8"
	cacheHost := NewCacheHost()
	record := &Record{
		Domain: domain,
		Addrs:  []net.IP{net.ParseIP(expectedIP)},
		TTL:    0,
	}
	cacheHost.Set(domain, record)
	record, err := cacheHost.Get(domain)
	assert.Nil(t, err)
	assert.Equal(t, 1, len(record.Addrs))
	assert.Equal(t, expectedIP, record.Addrs[0].String())
	cacheHost.Clear()
	record, err = cacheHost.Get(domain)
	assert.NotNil(t, err)
	assert.Nil(t, record)
}
