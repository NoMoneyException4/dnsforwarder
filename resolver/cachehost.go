package resolver

import (
	"encoding/json"
	"errors"
	"time"

	"github.com/codebear4/ttlcache"
)

type CacheHost struct {
	store *ttlcache.Cache
}

func NewCacheHost() *CacheHost {
	return &CacheHost{store: ttlcache.NewCache()}
}

func (host *CacheHost) Get(domain string) (error, *Record) {
	recordString, ok := host.store.Get(domain)
	if !ok {
		return errors.New("Not found."), nil
	}
	var record Record
	err := json.Unmarshal([]byte(recordString.(string)), &record)
	if err != nil {
		return err, nil
	}
	return nil, &record
}

func (host *CacheHost) Set(domain string, record Record) error {
	recordBytes, err := json.Marshal(record)
	if err != nil {
		return err
	}

	host.store.SetWithTTL(domain, string(recordBytes), time.Duration(record.Ttl)*time.Second)
	return nil
}

func (host *CacheHost) Refresh() {
	panic("Not implement yet.")
}
