package cards

import (
	"path/filepath"
	"testing"
	"traininggateway/domain"
	"traininggateway/storage"
)

func TestGenerateAndRedeem(t *testing.T) {
	s, _ := storage.Open(filepath.Join(t.TempDir(), "db"))
	defer s.Close()
	c := New(s, domain.DefaultPolicy())
	xs, e := c.Generate("TRN", 2, 1, 100)
	if e != nil || len(xs) != 2 {
		t.Fatal(e)
	}
	if _, e = c.Redeem(xs[0].Code, "e1", 1); e != nil {
		t.Fatal(e)
	}
}
