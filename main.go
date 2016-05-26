package main

import (
	"flag"

	. "github.com/codebear4/dnsforwarder/conf"
	. "github.com/codebear4/dnsforwarder/logger"
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
	Logger.Error(Conf.Loggers.File.Enable)
}
