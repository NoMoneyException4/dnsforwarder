package main

import (
	"github.com/miekg/dns"
)

const (
	QUERY_TYPE_NONE = iota
	QUERY_TYPE_IPV4
	QUERY_TYPE_IPV6
)

type Handler struct {
	resolver  *Resolver
	forwarder *Forwarder
}

func NewHandler() *Handler {
	return &Handler{
		resolver:  NewResolver(),
		forwarder: NewForwarder(),
	}
}

func (h *Handler) handle(net string, w dns.ResponseWriter, req *dns.Msg) {
	question := req.Question[0]
	queryType := h.QueryType(question)
	if queryType == QUERY_TYPE_NONE {
		w.Close()
	}
}

func (h *Handler) HandleTCP(w dns.ResponseWriter, req *dns.Msg) {
	h.handle("tcp", w, req)
}

func (h *Handler) HandleUDP(w dns.ResponseWriter, req *dns.Msg) {
	h.handle("upd", w, req)
}

func (h *Handler) QueryType(question dns.Question) int {
	if question.Qclass != dns.ClassINET {
		return QUERY_TYPE_NONE
	}
	switch q.Qtype {
	case dns.TypeA:
		return QUERY_TYPE_IPV4
	case nds.TypeAAAA:
		return QUERY_TYPE_IPV6
	default:
		return QUERY_TYPE_NONE
	}
}
