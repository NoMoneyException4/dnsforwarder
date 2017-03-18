package main

import (
	"errors"
	"strings"
	"sync"
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
		err = errors.New("a valid server config must contains `://` ")
		return
	}

	net = upstreamSlice[0]
	server = upstreamSlice[1]
	return
}

//Lookup Lookup the given domain with upstreams
func (f *Forwarder) Lookup(req *dns.Msg, net string) (*dns.Msg, error) {
	var wg sync.WaitGroup
	lookup := func(net string, req *dns.Msg, server string, result chan *dns.Msg, err chan error) {
		defer wg.Done()
		domain := req.Question[0].Name

		client := &dns.Client{
			Net:          net,
			DialTimeout:  time.Duration(Conf.Timeout.Forwarder.Read) * time.Millisecond * 10,
			WriteTimeout: time.Duration(Conf.Timeout.Forwarder.Write) * time.Millisecond * 10,
		}

		resp, _, lookupError := client.Exchange(req, server)
		if lookupError != nil {
			err <- lookupError
			return
		}
		if resp != nil && resp.Rcode == dns.RcodeServerFailure {
			Logger.Warningf("%s failed to get an valid record from upstream %s", domain, server)
			err <- errors.New("upstream server failure")
			return
		}

		result <- resp
		return
	}

	result := make(chan *dns.Msg, len(f.upstreams[net]))
	err := make(chan error, len(f.upstreams[net]))
	for _, server := range f.upstreams[net] {
		wg.Add(1)
		go lookup(net, req, server, result, err)
	}

	wg.Wait()
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
