package domain

type Policy struct {
	MaxSessionAge int64
	CheckInterval int64
	AllowDirect   bool
	Prefixes      []string
}

func DefaultPolicy() Policy {
	return Policy{MaxSessionAge: 3600, CheckInterval: 60, AllowDirect: true, Prefixes: []string{"TRN", "SAFE", "OPS"}}
}
func (p Policy) PrefixAllowed(prefix string) bool {
	for _, v := range p.Prefixes {
		if v == prefix {
			return true
		}
	}
	return false
}
func (p Policy) SessionFresh(created, now int64) bool {
	if created <= 0 || now < created {
		return false
	}
	return now-created <= p.MaxSessionAge
}
