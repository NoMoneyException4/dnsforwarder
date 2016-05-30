package main

import (
	"net"

	"github.com/miekg/dns"
)

const (
	//QTNone None ip query, etc: mx
	QTNone = 0
	//QTIPV4 IPv4 query
	QTIPV4 = dns.TypeA
	//QTIPV6 IPv6 query
	QTIPV6 = dns.TypeAAAA
)

//Handler Handle all the queries
type Handler struct {
	resolver  *Resolver
	forwarder *Forwarder
}

//NewHandler New Handler
func NewHandler() *Handler {
	return &Handler{
		resolver:  NewResolver(),
		forwarder: NewForwarder(),
	}
}

func (h *Handler) handle(net string, w dns.ResponseWriter, req *dns.Msg) {
	Logger.Infof("Handling query: %s.", req.Question[0].String())

	question := req.Question[0]
	queryType := h.queryType(question)
	if queryType == QTNone {
		Logger.Error("Invalid query type: QTNone")
		m := new(dns.Msg)
		m.SetRcode(req, dns.RcodeNotImplemented)
		w.WriteMsg(m)
	}

	//Lookup with resolver
	record, err := h.resolver.Lookup(question.Name)
	if err == nil {
		m := new(dns.Msg)
		m.SetReply(req)

		header := h.buildRRHeader(question.Name, int(queryType), Conf.Cache.TTL)
		for _, ip := range record.Addrs {
			answer := h.buildAnswer(header, ip)
			m.Answer = append(m.Answer, answer)
		}
		w.WriteMsg(m)
		return
	}

	//Lookup with forwarder
	resp, err := h.forwarder.Lookup(req, net)
	if resp == nil {
		m := new(dns.Msg)
		m.SetRcode(req, dns.RcodeServerFailure)
		w.WriteMsg(m)
		return
	}

	record = &Record{
		Domain: question.Name,
		TTL:    0,
	}
	for _, answer := range resp.Answer {
		if int(answer.Header().Ttl) > record.TTL {
			record.TTL = int(answer.Header().Ttl)
		}
		switch answer.(type) {
		case *dns.A:
			record.Addrs = append(record.Addrs, answer.(*dns.A).A)
		case *dns.AAAA:
			record.Addrs = append(record.Addrs, answer.(*dns.AAAA).AAAA)
		default:
			Logger.Errorf("Unsupport type: %#v", answer)
		}
	}

	h.resolver.Persist(question.Name, record)
	w.WriteMsg(resp)
	return
}

//HandleTCP Handle TCP conn
func (h *Handler) HandleTCP(w dns.ResponseWriter, req *dns.Msg) {
	h.handle("tcp", w, req)
}

//HandleUDP Handle UDP conn
func (h *Handler) HandleUDP(w dns.ResponseWriter, req *dns.Msg) {
	h.handle("udp", w, req)
}

func (h *Handler) queryType(question dns.Question) uint16 {
	if question.Qclass != dns.ClassINET {
		return QTNone
	}
	switch question.Qtype {
	case dns.TypeA:
		return QTIPV4
	case dns.TypeAAAA:
		return QTIPV6
	default:
		return QTNone
	}
}

func (h *Handler) buildRRHeader(name string, queryType, ttl int) dns.RR_Header {
	return dns.RR_Header{
		Name:   name,
		Rrtype: uint16(queryType),
		Class:  dns.ClassINET,
		Ttl:    uint32(ttl),
	}
}

func (h *Handler) buildAnswer(header dns.RR_Header, ip net.IP) dns.RR {
	if header.Rrtype == QTIPV4 {
		return &dns.A{header, ip.To4()}
	} else if header.Rrtype == QTIPV6 {
		return &dns.AAAA{header, ip.To16()}
	} else {
		Logger.Errorf("Unsupport query: %#v", header)
		return &dns.A{}
	}
}
