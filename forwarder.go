package main

//Forwarder Forwarder forward the request to upstreams
type Forwarder struct {
}

//NewForwarder New Forwarder
func NewForwarder() *Forwarder {
	return &Forwarder{}
}

//Lookup Lookup the given domain with upstreams
func (f *Forwarder) Lookup(domain string) (*Record, error) {
	return nil, nil
}
