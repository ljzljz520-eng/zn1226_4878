package cards

import (
	"fmt"
	"traininggateway/domain"
	"traininggateway/storage"
)

type Lifecycle struct{ store *storage.Store }

func NewLifecycle(s *storage.Store) *Lifecycle { return &Lifecycle{store: s} }
func (l *Lifecycle) Expire(now int64) (int, error) {
	cs, e := l.store.ListCards()
	if e != nil {
		return 0, e
	}
	n := 0
	for _, c := range cs {
		if c.Status == domain.CardActive && c.ExpiresAt > 0 && c.ExpiresAt <= now {
			c.Status = domain.CardExpired
			if e = l.store.SaveCard(c); e != nil {
				return n, e
			}
			n++
		}
	}
	return n, nil
}
func (l *Lifecycle) Restore(id string) error {
	c, e := l.store.Card(id)
	if e != nil {
		return e
	}
	if c.Status != domain.CardBlocked {
		return fmt.Errorf("only blocked cards can be restored")
	}
	c.Status = domain.CardActive
	return l.store.SaveCard(c)
}
func (l *Lifecycle) Revoke(id string) error {
	c, e := l.store.Card(id)
	if e != nil {
		return e
	}
	if c.Status == domain.CardExhausted {
		return fmt.Errorf("exhausted card")
	}
	c.Status = domain.CardBlocked
	return l.store.SaveCard(c)
}
func (l *Lifecycle) Eligible(id string, now int64) bool {
	c, e := l.store.Card(id)
	if e != nil {
		return false
	}
	return c.Usable(timeFromUnix(now))
}
