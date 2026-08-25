package storage

import (
	"encoding/json"
	"fmt"
	"go.etcd.io/bbolt"
	"traininggateway/domain"
)

func (s *Store) SaveCardAndAudit(c domain.AccessCard, a domain.AuditRecord) error {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.db.Update(func(tx *bbolt.Tx) error {
		if e := put(tx, buckets["cards"], c.ID, c); e != nil {
			return e
		}
		return put(tx, buckets["audits"], a.ID, a)
	})
}
func (s *Store) Count(bucket string) (int, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	b, ok := buckets[bucket]
	if !ok {
		return 0, fmt.Errorf("unknown bucket")
	}
	n := 0
	e := s.db.View(func(tx *bbolt.Tx) error {
		return tx.Bucket(b).ForEach(func(k, v []byte) error {
			if json.Valid(v) {
				n++
			}
			return nil
		})
	})
	return n, e
}
