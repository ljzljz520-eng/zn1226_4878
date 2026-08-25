package workflow

import (
	"fmt"
	"traininggateway/domain"
	"traininggateway/storage"
)

type Journal struct{ store *storage.Store }

func NewJournal(s *storage.Store) *Journal { return &Journal{store: s} }
func (j *Journal) Append(id, card, action, detail string, accepted bool, at int64) error {
	if id == "" || action == "" {
		return fmt.Errorf("missing journal fields")
	}
	return j.store.SaveAudit(domain.AuditRecord{ID: id, CardID: card, Action: action, Detail: detail, Accepted: accepted, CreatedAt: at})
}
func (j *Journal) HasAccepted(id string) bool {
	as, e := j.store.ListAudits()
	if e != nil {
		return false
	}
	for _, a := range as {
		if a.ID == id && a.Accepted {
			return true
		}
	}
	return false
}
func (j *Journal) Reject(id, reason string, at int64) error {
	return j.Append(id, "", "reject", reason, false, at)
}
