package admin

import (
	"sort"
	"traininggateway/domain"
	"traininggateway/storage"
)

type Dashboard struct {
	report *Report
	store  *storage.Store
}

func NewDashboard(s *storage.Store) *Dashboard { return &Dashboard{report: New(s), store: s} }
func (d *Dashboard) Snapshot(now int64) (domain.CardSummary, int, int, error) {
	s, e := d.report.CardSummary(now)
	if e != nil {
		return s, 0, 0, e
	}
	a, b, e := d.report.AuditCount()
	return s, a, b, e
}
func (d *Dashboard) RecentAudits(limit int) ([]domain.AuditRecord, error) {
	a, e := d.store.ListAudits()
	if e != nil {
		return nil, e
	}
	sort.Slice(a, func(i, j int) bool { return a[i].CreatedAt > a[j].CreatedAt })
	if limit < 0 {
		limit = 0
	}
	if limit > len(a) {
		limit = len(a)
	}
	return a[:limit], nil
}
func (d *Dashboard) Health(now int64) string {
	s, e := d.report.CardSummary(now)
	if e != nil {
		return "unknown"
	}
	if s.Blocked > 0 {
		return "degraded"
	}
	if s.Active == 0 {
		return "empty"
	}
	return "ready"
}
