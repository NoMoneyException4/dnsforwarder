package main

import (
	"errors"
	"fmt"
)

//ResolvProvider Interface for host providers
type ResolvProvider interface {
	Get(domain string) (*Record, error)
	Set(domain string, record *Record) error
	All() []*Record
	Clear()
	Refresh()
}

//Resolver Resolver struct
type Resolver struct {
	providers []ResolvProvider
}

//NewResolver New Resolver
func NewResolver() *Resolver {
	var providers []ResolvProvider
	if Conf.Hosts.Enable {
		fileHostProvider := NewFileHost()
		providers = append(providers, fileHostProvider)
	}

	if Conf.Cache.Enable {
		cacheHostProvider := NewCacheHost()
		providers = append(providers, cacheHostProvider)
	}
	return &Resolver{providers: providers}
}

//Lookup Lookup the given domain with providers
func (r *Resolver) Lookup(domain string) (*Record, error) {
	for _, provider := range r.providers {
		record, err := provider.Get(domain)
		if err == nil {
			return record, nil
		}
	}
	return nil, errors.New("Cannot resolv the given domain by resolv providers.")
}
