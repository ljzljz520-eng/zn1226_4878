package web

import (
	"net/http/httptest"
	"path/filepath"
	"testing"
	"traininggateway/cards"
	"traininggateway/domain"
	"traininggateway/storage"
	"traininggateway/video"
	"traininggateway/workflow"
)

func TestHandlerRedeem(t *testing.T) {
	s, _ := storage.Open(filepath.Join(t.TempDir(), "db"))
	defer s.Close()
	c := cards.New(s, domain.DefaultPolicy())
	c.Generate("TRN", 1, 0, 100)
	h := New(c, video.New(s), workflow.NewCheckout(c, s)).Handler()
	r := httptest.NewRequest("GET", "/redeem?code=TRN-1-CODE&employee=e", nil)
	w := httptest.NewRecorder()
	h.ServeHTTP(w, r)
	if w.Code != 200 {
		t.Fatal(w.Code)
	}
}
