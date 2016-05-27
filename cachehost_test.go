package main

import (
	"net"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCacheHostGetSet(t *testing.T) {
	expectedIP := "8.8.8.8"
	cacheHost := NewCacheHost()
	record := &Record{
		Domain: "fake.codebear.xyz",
		Addrs:  []net.IP{net.ParseIP(expectedIP)},
		Ttl:    0,
	}
	cacheHost.Set("fake.codebear.xyz", record)
	record, err := cacheHost.Get("fake.codebear.xyz")
	assert.Nil(t, err)
	assert.Equal(t, 1, len(record.Addrs))
	assert.Equal(t, expectedIP, record.Addrs[0].String())
}

func TestCacheHostAll(t *testing.T) {
	expectedIP := "8.8.8.8"
	cacheHost := NewCacheHost()

	for i := 0; i < 5; i++ {
		domain := "fake" + strconv.Itoa(i) + ".codebear.xyz"
		record := &Record{
			Domain: domain,
			Addrs:  []net.IP{net.ParseIP(expectedIP)},
			Ttl:    0,
		}
		cacheHost.Set(domain, record)
	}
	assert.True(t, len(cacheHost.All()) > 0, "Hosts must not be empty")
}

func TestCacheHostClear(t *testing.T) {
	InitLogger()
	expectedIP := "8.8.8.8"
	cacheHost := NewCacheHost()
	record := &Record{
		Domain: "fake.codebear.xyz",
		Addrs:  []net.IP{net.ParseIP(expectedIP)},
		Ttl:    0,
	}
	cacheHost.Set("fake.codebear.xyz", record)
	record, err := cacheHost.Get("fake.codebear.xyz")
	assert.Nil(t, err)
	assert.Equal(t, 1, len(record.Addrs))
	assert.Equal(t, expectedIP, record.Addrs[0].String())
	cacheHost.Clear()
	record, err = cacheHost.Get("fake.codebear.xyz")
	assert.NotNil(t, err)
	assert.Nil(t, record)
}
