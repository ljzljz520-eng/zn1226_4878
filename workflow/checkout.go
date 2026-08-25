package workflow

import (
	"fmt"
	"traininggateway/cards"
	"traininggateway/domain"
	"traininggateway/storage"
)

type Checkout struct {
	cards *cards.Service
	store *storage.Store
}

func NewCheckout(c *cards.Service, s *storage.Store) *Checkout { return &Checkout{cards: c, store: s} }
func (w *Checkout) Enter(code, employee string, now int64) (domain.EmployeeSession, error) {
	if code == "" || employee == "" {
		return domain.EmployeeSession{}, fmt.Errorf("missing credentials")
	}
	return w.cards.Redeem(cards.NormalizeCode(code), employee, now)
}
func (w *Checkout) Validate(session string, now int64) (bool, error) {
	return w.cards.Check(session, now)
}
func (w *Checkout) Cancel(session string) error {
	x, e := w.store.Session(session)
	if e != nil {
		return e
	}
	x.Active = false
	return w.store.SaveSession(x)
}
func (w *Checkout) Commit(session string) error {
	x, e := w.store.Session(session)
	if e != nil {
		return e
	}
	if !x.Active {
		return fmt.Errorf("cancelled")
	}
	return w.store.SaveAudit(domain.AuditRecord{ID: session + "-commit", CardID: x.CardID, Action: "commit", Accepted: true, CreatedAt: x.LastChecked})
}
// CancelAndCommit models the "update a record" business chain: the employee
// cancels, then the gateway must refuse the final commit. Cancellation is the
// terminal state — it must not be reverted, and no audit record may be
// persisted afterwards. Commit enforces that by rejecting any session whose
// Active flag is false, so we simply forward to it instead of reactivating.
func (w *Checkout) CancelAndCommit(session string) error {
	if e := w.Cancel(session); e != nil {
		return e
	}
	return w.Commit(session)
}
