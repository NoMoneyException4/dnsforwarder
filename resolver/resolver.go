package resolver

type ResolvProvider interface {
	Get(domain string) (error, Record)
	Set(domain string, record Record) error
}
