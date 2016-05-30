package main

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/miekg/dns"
)

//Forwarder Forwarder forward the request to upstreams
type Forwarder struct {
	upstreams map[string][]string
}

//NewForwarder New Forwarder
func NewForwarder() *Forwarder {
	upstreams := parseUpstreams(Conf.Upstreams)
	return &Forwarder{
		upstreams: upstreams,
	}
}

func parseUpstreams(upstreams []string) map[string][]string {
	servers := make(map[string][]string, 0)
	for _, upstream := range upstreams {
		net, server, err := parseUpstream(upstream)
		if err != nil {
			Logger.Error(err)
		}
		servers[net] = append(servers[net], server)
	}
	return servers
}

func parseUpstream(upstream string) (net, server string, err error) {
	upstreamSlice := strings.Split(upstream, "://")
	if len(upstreamSlice) != 2 {
		err = errors.New("A valid server config must contains `://` .")
		return
	}

	net = upstreamSlice[0]
	server = upstreamSlice[1]
	return
}

func lookupWithSpecifyServer(net string, req *dns.Msg, server string, result chan *dns.Msg, err chan error) {
	domain := req.Question[0].Name
	Logger.Debugf("Looking up %s with %s ...\n", domain, server)

	client := &dns.Client{
		Net:          net,
		ReadTimeout:  time.Duration(Conf.Timeout.Forwarder.Read) * time.Millisecond,
		WriteTimeout: time.Duration(Conf.Timeout.Forwarder.Write) * time.Millisecond,
	}

	resp, rtt, exchangeError := client.Exchange(req, server)
	if exchangeError != nil {
		Logger.Error(exchangeError)
		err <- exchangeError
		return
	}
	if resp != nil && resp.Rcode == dns.RcodeServerFailure {
		eStr := fmt.Sprintf("%s failed to get an valid code on %s", domain, server)
		Logger.Warning(eStr)
		err <- errors.New(eStr)
		return
	}

	if resp == nil || resp.Rcode != dns.RcodeSuccess || len(resp.Answer) == 0 {
		eStr := fmt.Sprintf("%s failed to get an valid answer on %s", domain, server)
		Logger.Warning(eStr)
		err <- errors.New(eStr)
		return
	}
	Logger.Debugf("%s resolv on %s (%s) ttl: %d.", domain, server, net, rtt)
	result <- resp
	return
}

//Lookup Lookup the given domain with upstreams
func (f *Forwarder) Lookup(req *dns.Msg, net string) (*dns.Msg, error) {
	result := make(chan *dns.Msg, 1)
	err := make(chan error, 1)

	for _, server := range f.upstreams[net] {
		go lookupWithSpecifyServer(net, req, server, result, err)
	}
	for {
		select {
		case r := <-result:
			return r, nil
		case e := <-err:
			return nil, e
		default:
		}
	}
}
