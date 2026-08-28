package admin

import (
	"path/filepath"
	"testing"
	"traininggateway/domain"
	"traininggateway/storage"
)

func TestReport(t *testing.T) {
	s, _ := storage.Open(filepath.Join(t.TempDir(), "db"))
	defer s.Close()
	s.SaveCard(domain.AccessCard{ID: "b", Status: domain.CardBlocked})
	r := New(s)
	x, e := r.CardSummary(1)
	if e != nil || x.Blocked != 1 {
		t.Fatal(e)
	}
}
