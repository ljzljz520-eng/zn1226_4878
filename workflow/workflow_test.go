package workflow

import (
	"path/filepath"
	"testing"
	"traininggateway/cards"
	"traininggateway/domain"
	"traininggateway/storage"
)

func setup(t *testing.T) (*Checkout, *storage.Store) {
	s, _ := storage.Open(filepath.Join(t.TempDir(), "db"))
	c := cards.New(s, domain.DefaultPolicy())
	c.Generate("TRN", 1, 0, 100)
	return NewCheckout(c, s), s
}
func TestWorkflowOne(t *testing.T) {
	w, s := setup(t)
	defer s.Close()
	x, e := w.Enter("TRN-1-CODE", "e", 1)
	if e != nil {
		t.Fatal(e)
	}
	ok, e := w.Validate(x.ID, 1)
	if e != nil || !ok {
		t.Fatal(e)
	}
}
func TestWorkflowTwo(t *testing.T) {
	w, s := setup(t)
	defer s.Close()
	x, _ := w.Enter("TRN-1-CODE", "e", 1)
	if e := w.Cancel(x.ID); e != nil {
		t.Fatal(e)
	}
}
func TestWorkflowThree(t *testing.T) {
	w, s := setup(t)
	defer s.Close()
	x, _ := w.Enter("TRN-1-CODE", "e", 1)
	if e := w.Commit(x.ID); e != nil {
		t.Fatal(e)
	}
}
func TestBusinessChain23(t *testing.T) {
	w, s := setup(t)
	defer s.Close()
	x, _ := w.Enter("TRN-1-CODE", "e", 1)
	if e := w.CancelAndCommit(x.ID); e == nil {
		t.Fatal("cancelled operation was committed")
	}
}
