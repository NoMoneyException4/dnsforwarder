package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var (
	version string

	confFilePath string
	listenHost   string
	listenPort   string
	debug        bool
)

func initial() {
	flag.StringVar(&confFilePath, "c", "dnsforwarder.yml", "Configuration file path.")
	flag.StringVar(&listenHost, "h", "", "Listening host.")
	flag.StringVar(&listenPort, "p", "53", "Listening port.")
	flag.BoolVar(&debug, "d", false, "Debug Mode")
	flag.Usage = func() {
		_, err := fmt.Fprintf(os.Stderr, "DnsForwarder version %s\n", version)
		if err != nil {
			Logger.Error(err)
		}
		flag.PrintDefaults()
	}
	flag.Parse()

	LoadConf(confFilePath)
	InitLogger()
}

func main() {
	initial()

	if debug {
		go func() {
			log.Println(http.ListenAndServe("localhost:6060", nil))
		}()
	}

	server := NewServer(
		listenHost,
		listenPort,
		time.Duration(Conf.Timeout.Server.Read)*time.Millisecond,
		time.Duration(Conf.Timeout.Server.Write)*time.Millisecond,
	)
	server.Listen()

	sig := make(chan os.Signal)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)

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
