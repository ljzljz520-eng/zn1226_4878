package traininggateway

import (
	"path/filepath"
	"testing"
	"traininggateway/cards"
	"traininggateway/domain"
	"traininggateway/storage"
	"traininggateway/video"
	"traininggateway/workflow"
)

func TestGatewayEndToEnd(t *testing.T) {
	s, _ := storage.Open(filepath.Join(t.TempDir(), "db"))
	defer s.Close()
	c := cards.New(s, domain.DefaultPolicy())
	cat := video.New(s)
	cat.Publish(domain.Video{ID: "v", Title: "Intro", URL: "/v"})
	c.Generate("TRN", 1, 0, 100)
	w := workflow.NewCheckout(c, s)
	x, e := w.Enter("TRN-1-CODE", "emp", 1)
	if e != nil {
		t.Fatal(e)
	}
	ok, e := w.Validate(x.ID, 1)
	if e != nil || !ok {
		t.Fatal(e)
	}
	p, e := cat.Browse(domain.CatalogFilter{PublishedOnly: true})
	if e != nil || len(p.Items) != 1 {
		t.Fatal(e)
	}
}
