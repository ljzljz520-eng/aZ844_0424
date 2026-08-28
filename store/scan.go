package store

import (
	"coldchain/domain"
	"errors"
	"go.etcd.io/bbolt"
)

func (s *Store) QueryRecords(q string, zone string, status string) ([]domain.Record, error) {
	rs, e := s.ListRecords()
	if e != nil {
		return nil, e
	}
	q = normalizeQuery(q)
	out := []domain.Record{}
	for _, r := range filterRecords(rs, zone, status) {
		if matchesBatch(r, q) {
			out = append(out, r)
		}
	}
	return out, nil
}
func (s *Store) DeleteRecord(id string) error {
	s.mu.RLock()
	defer s.mu.RUnlock()
	if s.db == nil {
		return errors.New("closed")
	}
	return s.db.Update(func(tx *bbolt.Tx) error { return tx.Bucket(buckets["records"]).Delete([]byte(id)) })
}
