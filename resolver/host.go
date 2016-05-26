package resolver

type Hosts struct {
	hosts map[string]string
}

func (h *Hosts) Get(domain string) (error, Record) {

}

func (h *Hosts) Set(domain string, record Record) error {
	panic("Not implement yet.")
	return nil
}
