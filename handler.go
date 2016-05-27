package main

import (
	"fmt"

	"github.com/miekg/dns"
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
	fmt.Printf("%#v\n", req)
}

func (h *Handler) HandleTCP(w dns.ResponseWriter, req *dns.Msg) {
	h.handle("tcp", w, req)
}

func (h *Handler) HandleUDP(w dns.ResponseWriter, req *dns.Msg) {
	h.handle("upd", w, req)
}
