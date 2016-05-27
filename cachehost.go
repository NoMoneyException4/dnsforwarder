package main

import (
	"encoding/json"
	"errors"
	"time"

	"github.com/codebear4/ttlcache"
)

//CacheHost In-memory based host cache
type CacheHost struct {
	store *ttlcache.Cache
}

//NewCacheHost New CacheHost
func NewCacheHost() *CacheHost {
	return &CacheHost{store: ttlcache.NewCache()}
}

//Get Get host from cache
func (host *CacheHost) Get(domain string) (*Record, error) {
	recordString, ok := host.store.Get(domain)
	if !ok {
		Logger.Debugf("Domain %s not found in the memory cache.", domain)
		return nil, errors.New("Not found.")
	}
	var record Record
	err := json.Unmarshal([]byte(recordString.(string)), &record)
	if err != nil {
		Logger.Debugf("Domain record %s unmarshal failed.", domain)
		return nil, err
	}
	return &record, nil
}

//All Get all hosts
func (host *CacheHost) All() []*Record {
	allKeys := host.store.Items()
	all := make([]*Record, len(allKeys))
	for key, _ := range allKeys {
		record, err := host.Get(key)
		if err == nil {
			all = append(all, record)
		}
	}
	return all
}

//Clear Purge hosts
func (host *CacheHost) Clear() {
	host.store = ttlcache.NewCache()
}

//Set Set host with given Record
func (host *CacheHost) Set(domain string, record *Record) error {
	recordBytes, err := json.Marshal(record)
	if err != nil {
		Logger.Debugf("Domain record %s set failed.%#v", domain, record)
		return err
	}

	host.store.SetWithTTL(domain, string(recordBytes), time.Duration(record.Ttl)*time.Second)
	return nil
}

//Refresh Refresh the cached records
func (host *CacheHost) Refresh() {
	panic("Not implement yet.")
}
