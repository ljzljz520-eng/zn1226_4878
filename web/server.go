package web

import (
	"encoding/json"
	"net/http"
	"traininggateway/cards"
	"traininggateway/domain"
	"traininggateway/video"
	"traininggateway/workflow"
)

type Server struct {
	checkout *workflow.Checkout
	catalog  *video.Catalog
	cards    *cards.Service
}

func New(c *cards.Service, cat *video.Catalog, w *workflow.Checkout) *Server {
	return &Server{cards: c, catalog: cat, checkout: w}
}
func (s *Server) Handler() http.Handler {
	m := http.NewServeMux()
	m.HandleFunc("/redeem", s.redeem)
	m.HandleFunc("/catalog", s.catalogHandler)
	m.HandleFunc("/video", s.videoHandler)
	return m
}
func (s *Server) redeem(w http.ResponseWriter, r *http.Request) {
	x, e := s.checkout.Enter(r.URL.Query().Get("code"), r.URL.Query().Get("employee"), 1)
	if e != nil {
		http.Error(w, e.Error(), 403)
		return
	}
	json.NewEncoder(w).Encode(x)
}
func (s *Server) catalogHandler(w http.ResponseWriter, r *http.Request) {
	p, e := s.catalog.Browse(domain.CatalogFilter{PublishedOnly: true, Query: r.URL.Query().Get("q")})
	if e != nil {
		http.Error(w, e.Error(), 500)
		return
	}
	json.NewEncoder(w).Encode(p)
}
func (s *Server) videoHandler(w http.ResponseWriter, r *http.Request) {
	ok, _ := s.checkout.Validate(r.URL.Query().Get("session"), 1)
	v, e := s.catalog.Direct(r.URL.Query().Get("id"), ok)
	if e != nil {
		http.Error(w, e.Error(), 403)
		return
	}
	json.NewEncoder(w).Encode(v)
}
