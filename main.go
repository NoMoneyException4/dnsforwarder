package main

import (
	"flag"
	"fmt"
	"net"
)

func initial() {
	var confFilePath string

	flag.StringVar(&confFilePath, "c", "dnsforwarder.yml", "Configuration file path.")
	flag.Parse()

	LoadConf(confFilePath)
	InitLogger()
}

func main() {
	initial()
	cacheHost := NewCacheHost()
	record := &Record{
		Domain: "www.sipin.frank",
		Addrs:  []net.IP{net.ParseIP("119.29.29.29")},
		Ttl:    0,
	}
	cacheHost.Set("www.sipin.frank", record)
	record, err := cacheHost.Get("www.sipin.frank")
	if err != nil {
		panic(err)
	}
	for _, ip := range record.Addrs {
		fmt.Println(ip.String())
	}
}
