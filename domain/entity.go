package domain

import "time"

type AccessCard struct {
	ID, Code, Prefix, Status string
	Uses, MaxUses            int
	ExpiresAt                int64
	CreatedAt                int64
}
type Video struct {
	ID, Title, Description, URL, Category string
	DurationSeconds                       int
	Published                             bool
}
type EmployeeSession struct {
	ID, CardID, EmployeeID string
	Active                 bool
	LastChecked            int64
	CreatedAt              int64
}
type AuditRecord struct {
	ID, CardID, Action, Detail string
	Accepted                   bool
	CreatedAt                  int64
}

const (
	CardActive    = "active"
	CardExpired   = "expired"
	CardBlocked   = "blocked"
	CardExhausted = "exhausted"
)

func (c AccessCard) Usable(now time.Time) bool {
	if c.Status != CardActive {
		return false
	}
	if c.ExpiresAt > 0 && now.Unix() >= c.ExpiresAt {
		return false
	}
	if c.MaxUses > 0 && c.Uses >= c.MaxUses {
		return false
	}
	return true
}
func (s EmployeeSession) Valid() bool { return s.Active && s.CardID != "" && s.EmployeeID != "" }
func (v Video) Ready() bool           { return v.Published && v.ID != "" && v.URL != "" }
