package main

import (
	"flag"
	"fmt"

	. "github.com/codebear4/dnsforwarder/conf"
	. "github.com/codebear4/dnsforwarder/logger"
	"github.com/codebear4/dnsforwarder/resolver"
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
	fileHost := resolver.NewFileHost()
	err, record := fileHost.Get("www.sipin.frank")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%#v", record)
}
