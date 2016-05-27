package main

import (
	"flag"
	"os"
	"os/signal"
	"time"
)

var (
	confFilePath string
	listenHost   string
	listenPort   string
)

func initial() {
	flag.StringVar(&confFilePath, "c", "dnsforwarder.yml", "Configuration file path.")
	flag.StringVar(&listenHost, "h", "", "Listening host.")
	flag.StringVar(&listenPort, "p", "53", "Listening port.")
	flag.Parse()

	LoadConf(confFilePath)
	InitLogger()
}

func main() {
	initial()
	server := NewServer(listenHost, listenPort, time.Duration(Conf.Timeout.Read)*time.Millisecond, time.Duration(Conf.Timeout.Write)*time.Millisecond)
	server.Listen()

	sig := make(chan os.Signal)
	signal.Notify(sig, os.Interrupt)

forever:
	for {
		select {
		case <-sig:
			Logger.Info("Signal received, stopping.")
			// TODO: save cache to local file
			break forever
		}
	}
}
