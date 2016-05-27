package main

type ResolvProvider interface {
	Get(domain string) (*Record, error)
	Set(domain string, record *Record) error
	Refresh()
}
