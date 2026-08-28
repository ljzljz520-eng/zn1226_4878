package storage

import (
	"path/filepath"
	"testing"
	"traininggateway/domain"
)

func TestPersistenceSurvivesReopen(t *testing.T) {
	p := filepath.Join(t.TempDir(), "db")
	s, e := Open(p)
	if e != nil {
		t.Fatal(e)
	}
	if e = s.SaveCard(domain.AccessCard{ID: "A", Code: "A"}); e != nil {
		t.Fatal(e)
	}
	s.Close()
	s, e = Open(p)
	if e != nil {
		t.Fatal(e)
	}
	defer s.Close()
	if _, e = s.Card("A"); e != nil {
		t.Fatal(e)
	}
}
