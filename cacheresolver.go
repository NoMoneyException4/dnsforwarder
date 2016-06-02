package main

import (
	"errors"
	"time"

	"github.com/codebear4/ttlcache"
)

//CacheResolver In-memory based host resolver
type CacheResolver struct {
	store *ttlcache.Cache
}

//NewCacheResolver New CacheResolver
func NewCacheResolver() *CacheResolver {
	host := &CacheResolver{store: ttlcache.NewCache()}
	return host
}

//Get Get host from cache
func (host *CacheResolver) Get(domain string) (*Record, error) {
	record, ok := host.store.Get(domain)
	if !ok {
		return nil, errors.New("Not found.")
	}
	Logger.Debugf("[HitCache] Get record of domain %s from cache.", domain)
	return record.(*Record), nil
}

//All Get all hosts
func (host *CacheResolver) All() []*Record {
	allKeys := host.store.Items()
	all := make([]*Record, len(allKeys))
	for key := range allKeys {
		record, err := host.Get(key)
		if err == nil {
			all = append(all, record)
		}
	}
	return all
}

//Clear Purge hosts
func (host *CacheResolver) Clear() {
	host.store = ttlcache.NewCache()
}

//Set Set host with given Record
func (host *CacheResolver) Set(domain string, record *Record) error {
	host.store.SetWithTTL(domain, record, time.Duration(record.TTL)*time.Second)
	return nil
}
