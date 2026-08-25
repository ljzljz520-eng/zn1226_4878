package storage

func (s *Store) Healthy() bool { s.mu.RLock(); defer s.mu.RUnlock(); return s.db != nil }
func (s *Store) Flush() error {
	if !s.Healthy() {
		return nil
	}
	return s.db.Sync()
}
