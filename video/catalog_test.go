package video

import (
	"path/filepath"
	"testing"
	"traininggateway/domain"
	"traininggateway/storage"
)

func TestBrowsePublished(t *testing.T) {
	s, _ := storage.Open(filepath.Join(t.TempDir(), "db"))
	defer s.Close()
	c := New(s)
	c.Add(domain.Video{ID: "v1", Title: "Go", URL: "/go"})
	c.Publish(domain.Video{ID: "v2", Title: "SQL", URL: "/sql"})
	p, e := c.Browse(PublishedOnly())
	if e != nil || len(p.Items) != 1 {
		t.Fatal(e, len(p.Items))
	}
}
