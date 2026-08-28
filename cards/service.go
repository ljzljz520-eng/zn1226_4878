package cards

import (
	"fmt"
	"strconv"
	"strings"
	"time"
	"traininggateway/domain"
	"traininggateway/storage"
)

type Service struct {
	store  *storage.Store
	policy domain.Policy
}

func New(s *storage.Store, p domain.Policy) *Service { return &Service{store: s, policy: p} }
func (s *Service) Generate(prefix string, count, maxUses int, expires int64) ([]domain.AccessCard, error) {
	if !s.policy.PrefixAllowed(prefix) {
		return nil, fmt.Errorf("prefix not allowed")
	}
	if count <= 0 || count > 1000 {
		return nil, fmt.Errorf("invalid count")
	}
	out := make([]domain.AccessCard, 0, count)
	for i := 0; i < count; i++ {
		id := prefix + "-" + strconv.Itoa(i+1)
		c := domain.AccessCard{ID: id, Code: id + "-CODE", Prefix: prefix, Status: domain.CardActive, MaxUses: maxUses, ExpiresAt: expires, CreatedAt: 1}
		if e := s.store.SaveCard(c); e != nil {
			return nil, e
		}
		out = append(out, c)
	}
	return out, nil
}
func (s *Service) Redeem(code, employee string, now int64) (domain.EmployeeSession, error) {
	cards, e := s.store.ListCards()
	if e != nil {
		return domain.EmployeeSession{}, e
	}
	for _, c := range cards {
		if c.Code == code {
			if !c.Usable(timeFromUnix(now)) {
				return domain.EmployeeSession{}, fmt.Errorf("card unavailable")
			}
			c.Uses++
			if c.MaxUses > 0 && c.Uses >= c.MaxUses {
				c.Status = domain.CardExhausted
			}
			if e = s.store.SaveCard(c); e != nil {
				return domain.EmployeeSession{}, e
			}
			x := domain.EmployeeSession{ID: c.ID + "-" + employee, CardID: c.ID, EmployeeID: employee, Active: true, LastChecked: now, CreatedAt: now}
			if e = s.store.SaveSession(x); e != nil {
				return domain.EmployeeSession{}, e
			}
			return x, nil
		}
	}
	return domain.EmployeeSession{}, fmt.Errorf("card not found")
}
func (s *Service) Check(id string, now int64) (bool, error) {
	x, e := s.store.Session(id)
	if e != nil {
		return false, e
	}
	if !x.Valid() || !s.policy.SessionFresh(x.CreatedAt, now) {
		return false, nil
	}
	c, e := s.store.Card(x.CardID)
	if e != nil {
		return false, e
	}
	return c.Usable(timeFromUnix(now)), nil
}
func (s *Service) Block(id string) error {
	c, e := s.store.Card(id)
	if e != nil {
		return e
	}
	c.Status = domain.CardBlocked
	return s.store.SaveCard(c)
}
func (s *Service) Summary(now int64) (domain.CardSummary, error) {
	cs, e := s.store.ListCards()
	if e != nil {
		return domain.CardSummary{}, e
	}
	r := domain.CardSummary{Total: len(cs)}
	for _, c := range cs {
		switch c.Status {
		case domain.CardBlocked:
			r.Blocked++
		case domain.CardExhausted:
			r.Exhausted++
		default:
			if c.ExpiresAt > 0 && c.ExpiresAt <= now {
				r.Expired++
			} else if c.Uses > 0 {
				r.Used++
			} else {
				r.Active++
			}
		}
	}
	return r, nil
}
func NormalizeCode(v string) string  { return strings.ToUpper(strings.TrimSpace(v)) }
func timeFromUnix(v int64) time.Time { return time.Unix(v, 0) }
