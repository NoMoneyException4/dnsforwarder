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
}
