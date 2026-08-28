package cards

import (
	"path/filepath"
	"testing"
	"traininggateway/domain"
	"traininggateway/storage"
)

func TestSummary(t *testing.T) {
	s, _ := storage.Open(filepath.Join(t.TempDir(), "db"))
	defer s.Close()
	c := New(s, domain.DefaultPolicy())
	c.Generate("TRN", 1, 0, 2)
	x, e := c.Summary(3)
	if e != nil || x.Expired != 1 {
		t.Fatalf("%+v %v", x, e)
	}
}
