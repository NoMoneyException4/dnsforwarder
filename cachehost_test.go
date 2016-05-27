package main

import (
	"net"
	"testing"
)

func TestCacheHostGet(t *testing.T) {
	expectedIP := "8.8.8.8"
	cacheHost := NewCacheHost()
	record := &Record{
		Domain: "fake.codebear.xyz",
		Addrs:  []net.IP{net.ParseIP(expectedIP)},
		Ttl:    0,
	}
	cacheHost.Set("fake.codebear.xyz", record)
	record, err := cacheHost.Get("fake.codebear.xyz")
	if err != nil {
		t.Fatal(err)
	}
	if len(record.Addrs) != 1 {
		t.Fatalf("Wrong count of Addrs")
	}
	ip := record.Addrs[0]
	if expectedIP != ip.String() {
		t.Fatalf("Wrong ip addr")
	}
}
