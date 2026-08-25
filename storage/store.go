package storage

import (
	"encoding/json"
	"fmt"
	"go.etcd.io/bbolt"
	"path/filepath"
	"sync"
	"traininggateway/domain"
)

var buckets = map[string][]byte{"cards": []byte("cards"), "videos": []byte("videos"), "sessions": []byte("sessions"), "audits": []byte("audits")}

type Store struct {
	db *bbolt.DB
	mu sync.RWMutex
}

func Open(path string) (*Store, error) {
	db, err := bbolt.Open(filepath.Clean(path), 0600, nil)
	if err != nil {
		return nil, err
	}
	s := &Store{db: db}
	err = db.Update(func(tx *bbolt.Tx) error {
		for _, b := range buckets {
			if _, e := tx.CreateBucketIfNotExists(b); e != nil {
				return e
			}
		}
		return nil
	})
	if err != nil {
		db.Close()
		return nil, err
	}
	return s, nil
}
func (s *Store) Close() error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.db == nil {
		return nil
	}
	err := s.db.Close()
	s.db = nil
	return err
}
func put[T any](tx *bbolt.Tx, b []byte, key string, v T) error {
	raw, e := json.Marshal(v)
	if e != nil {
		return e
	}
	return tx.Bucket(b).Put([]byte(key), raw)
}
func get[T any](tx *bbolt.Tx, b []byte, key string, out *T) error {
	raw := tx.Bucket(b).Get([]byte(key))
	if raw == nil {
		return fmt.Errorf("not found")
	}
	return json.Unmarshal(raw, out)
}
func (s *Store) SaveCard(c domain.AccessCard) error {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.db.Update(func(tx *bbolt.Tx) error { return put(tx, buckets["cards"], c.ID, c) })
}
func (s *Store) Card(id string) (domain.AccessCard, error) {
	var c domain.AccessCard
	s.mu.RLock()
	defer s.mu.RUnlock()
	e := s.db.View(func(tx *bbolt.Tx) error { return get(tx, buckets["cards"], id, &c) })
	return c, e
}
func (s *Store) ListCards() ([]domain.AccessCard, error) {
	out := []domain.AccessCard{}
	s.mu.RLock()
	defer s.mu.RUnlock()
	e := s.db.View(func(tx *bbolt.Tx) error {
		return tx.Bucket(buckets["cards"]).ForEach(func(_, v []byte) error {
			var c domain.AccessCard
			if e := json.Unmarshal(v, &c); e != nil {
				return e
			}
			out = append(out, c)
			return nil
		})
	})
	return out, e
}
func (s *Store) SaveVideo(v domain.Video) error {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.db.Update(func(tx *bbolt.Tx) error { return put(tx, buckets["videos"], v.ID, v) })
}
func (s *Store) ListVideos() ([]domain.Video, error) {
	out := []domain.Video{}
	s.mu.RLock()
	defer s.mu.RUnlock()
	e := s.db.View(func(tx *bbolt.Tx) error {
		return tx.Bucket(buckets["videos"]).ForEach(func(_, v []byte) error {
			var x domain.Video
			if e := json.Unmarshal(v, &x); e != nil {
				return e
			}
			out = append(out, x)
			return nil
		})
	})
	return out, e
}
func (s *Store) SaveSession(x domain.EmployeeSession) error {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.db.Update(func(tx *bbolt.Tx) error { return put(tx, buckets["sessions"], x.ID, x) })
}
func (s *Store) Session(id string) (domain.EmployeeSession, error) {
	var x domain.EmployeeSession
	s.mu.RLock()
	defer s.mu.RUnlock()
	e := s.db.View(func(tx *bbolt.Tx) error { return get(tx, buckets["sessions"], id, &x) })
	return x, e
}
func (s *Store) SaveAudit(a domain.AuditRecord) error {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.db.Update(func(tx *bbolt.Tx) error { return put(tx, buckets["audits"], a.ID, a) })
}
func (s *Store) ListAudits() ([]domain.AuditRecord, error) {
	out := []domain.AuditRecord{}
	s.mu.RLock()
	defer s.mu.RUnlock()
	e := s.db.View(func(tx *bbolt.Tx) error {
		return tx.Bucket(buckets["audits"]).ForEach(func(_, v []byte) error {
			var a domain.AuditRecord
			if e := json.Unmarshal(v, &a); e != nil {
				return e
			}
			out = append(out, a)
			return nil
		})
	})
	return out, e
}
