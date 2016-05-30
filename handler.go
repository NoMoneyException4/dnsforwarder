package main

import (
	"net"

	"github.com/miekg/dns"
)

//Handler Handle all the queries
type Handler struct {
	cacheResolver *CacheResolver
	fileResolver  *FileResolver
	forwarder     *Forwarder
}

//NewHandler New Handler
func NewHandler() *Handler {
	return &Handler{
		cacheResolver: NewCacheResolver(),
		fileResolver:  NewFileResolver(),
		forwarder:     NewForwarder(),
	}
}

func (h *Handler) handle(net string, w dns.ResponseWriter, req *dns.Msg) {
	question := req.Question[0]
	Logger.Infof("Query %s for %s.", question.Name, dns.Type(question.Qtype).String())

	switch question.Qtype {
	case dns.TypeA, dns.TypeAAAA:
		if addrs, err := h.fileResolver.Get(question.Name); err == nil {
			m := new(dns.Msg)
			m.SetReply(req)
			header := h.buildRRHeader(question.Name, question.Qtype, Conf.Cache.TTL)
			for _, addr := range addrs {
				answer := h.buildAnswer(header, addr)
				m.Answer = append(m.Answer, answer)
			}
			w.WriteMsg(m)
			return
		}
		fallthrough
	default:
		if record, err := h.cacheResolver.Get(question.Name); err == nil {
			m := new(dns.Msg)
			m.SetReply(req)
			for _, answer := range record.Answers {
				m.Answer = append(m.Answer, answer)
			}
			w.WriteMsg(m)
			return
		}
		if msg, err := h.forwarder.Lookup(req, net); err == nil {
			if len(msg.Answer) > 0 {
				h.cacheResolver.Set(question.Name, &Record{
					Domain:  question.Name,
					Type:    question.Qtype,
					TTL:     msg.Answer[0].Header().Ttl,
					Answers: msg.Answer,
				})
				Logger.Infof("Domain %s cached.", question.Name)
			}
			w.WriteMsg(msg)
			return
		} else {
			// Logger.Error(err)
			dns.HandleFailed(w, req)
		}
	}
}

//HandleTCP Handle TCP conn
func (h *Handler) HandleTCP(w dns.ResponseWriter, req *dns.Msg) {
	h.handle("tcp", w, req)
}

//HandleUDP Handle UDP conn
func (h *Handler) HandleUDP(w dns.ResponseWriter, req *dns.Msg) {
	h.handle("udp", w, req)
}

func (h *Handler) buildRRHeader(name string, qtype uint16, ttl uint32) dns.RR_Header {
	return dns.RR_Header{
		Name:   name,
		Rrtype: qtype,
		Class:  dns.ClassINET,
		Ttl:    ttl,
	}
}

func (h *Handler) buildAnswer(header dns.RR_Header, target string) dns.RR {
	switch header.Rrtype {
	case dns.TypeA:
		return &dns.A{header, net.ParseIP(target).To4()}
	case dns.TypeAAAA:
		return &dns.A{header, net.ParseIP(target).To16()}
	default:
		Logger.Errorf("Unsupport query: %#v", header)
		return &dns.A{}
	}
}
