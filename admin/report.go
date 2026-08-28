package admin

import (
	"traininggateway/domain"
	"traininggateway/storage"
)

type Report struct{ store *storage.Store }

func New(s *storage.Store) *Report { return &Report{store: s} }
func (r *Report) CardSummary(now int64) (domain.CardSummary, error) {
	cs, e := r.store.ListCards()
	if e != nil {
		return domain.CardSummary{}, e
	}
	out := domain.CardSummary{Total: len(cs)}
	for _, c := range cs {
		if c.Status == domain.CardBlocked {
			out.Blocked++
		} else if c.Status == domain.CardExhausted {
			out.Exhausted++
		} else if c.ExpiresAt > 0 && c.ExpiresAt <= now {
			out.Expired++
		} else if c.Uses > 0 {
			out.Used++
		} else {
			out.Active++
		}
	}
	return out, nil
}
func (r *Report) AuditCount() (int, int, error) {
	as, e := r.store.ListAudits()
	if e != nil {
		return 0, 0, e
	}
	ok, bad := 0, 0
	for _, a := range as {
		if a.Accepted {
			ok++
		} else {
			bad++
		}
	}
	return ok, bad, nil
}
func (r *Report) Record(a domain.AuditRecord) error { return r.store.SaveAudit(a) }
