package video

import (
	"sort"
	"strings"
	"traininggateway/domain"
	"traininggateway/storage"
)

type Catalog struct{ store *storage.Store }

func New(s *storage.Store) *Catalog             { return &Catalog{store: s} }
func (c *Catalog) Publish(v domain.Video) error { v.Published = true; return c.store.SaveVideo(v) }
func (c *Catalog) Add(v domain.Video) error     { return c.store.SaveVideo(v) }
func (c *Catalog) Browse(f domain.CatalogFilter) (domain.Page[domain.Video], error) {
	all, e := c.store.ListVideos()
	if e != nil {
		return domain.Page[domain.Video]{}, e
	}
	out := []domain.Video{}
	for _, v := range all {
		if f.PublishedOnly && !v.Published {
			continue
		}
		if f.Category != "" && v.Category != f.Category {
			continue
		}
		if f.Query != "" && !strings.Contains(strings.ToLower(v.Title), strings.ToLower(f.Query)) {
			continue
		}
		out = append(out, v)
	}
	sort.Slice(out, func(i, j int) bool { return out[i].Title < out[j].Title })
	return domain.Page[domain.Video]{Items: out, Total: len(out), Limit: len(out)}, nil
}
func (c *Catalog) Direct(id string, allowed bool) (domain.Video, error) {
	if !allowed {
		return domain.Video{}, ErrUnauthorized{}
	}
	all, e := c.store.ListVideos()
	if e != nil {
		return domain.Video{}, e
	}
	for _, v := range all {
		if v.ID == id && v.Ready() {
			return v, nil
		}
	}
	return domain.Video{}, ErrMissing{}
}

type ErrUnauthorized struct{}

func (ErrUnauthorized) Error() string { return "unauthorized" }

type ErrMissing struct{}

func (ErrMissing) Error() string { return "video not found" }
