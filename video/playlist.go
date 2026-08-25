package video

import (
	"fmt"
	"sort"
	"traininggateway/domain"
)

type Playlist struct {
	ID, Name string
	VideoIDs []string
}

func NewPlaylist(id, name string) *Playlist {
	return &Playlist{ID: id, Name: name, VideoIDs: []string{}}
}
func (p *Playlist) Add(id string) error {
	if id == "" {
		return fmt.Errorf("empty video")
	}
	for _, v := range p.VideoIDs {
		if v == id {
			return fmt.Errorf("duplicate video")
		}
	}
	p.VideoIDs = append(p.VideoIDs, id)
	return nil
}
func (p *Playlist) Remove(id string) bool {
	for i, v := range p.VideoIDs {
		if v == id {
			p.VideoIDs = append(p.VideoIDs[:i], p.VideoIDs[i+1:]...)
			return true
		}
	}
	return false
}
func (p Playlist) Ordered() []string { return append([]string(nil), p.VideoIDs...) }
func RankVideos(vs []domain.Video, query string) []domain.Video {
	out := append([]domain.Video(nil), vs...)
	sort.SliceStable(out, func(i, j int) bool {
		if query != "" {
			ai := out[i].Title == query
			aj := out[j].Title == query
			if ai != aj {
				return ai
			}
		}
		return out[i].Title < out[j].Title
	})
	return out
}
