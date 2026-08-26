package store

import (
	"coldchain/domain"
	"encoding/json"
	"errors"
	"fmt"
	"go.etcd.io/bbolt"
	"path/filepath"
	"sync"
	"time"
)

var buckets = map[string][]byte{"records": []byte("records"), "profiles": []byte("profiles"), "events": []byte("events"), "audits": []byte("audits")}

type Store struct {
	db   *bbolt.DB
	mu   sync.RWMutex
	path string
}

func Open(path string) (*Store, error) {
	db, err := bbolt.Open(filepath.Clean(path), 0600, bbolt.DefaultOptions)
	if err != nil {
		return nil, err
	}
	s := &Store{db: db, path: path}
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
func (s *Store) Path() string { return s.path }
func put(tx *bbolt.Tx, b []byte, key string, v any) error {
	data, e := json.Marshal(v)
	if e != nil {
		return e
	}
	return tx.Bucket(b).Put([]byte(key), data)
}
func get(tx *bbolt.Tx, b []byte, key string, v any) error {
	data := tx.Bucket(b).Get([]byte(key))
	if data == nil {
		return errors.New("not found")
	}
	return json.Unmarshal(data, v)
}
func (s *Store) SaveRecord(r domain.Record) error {
	s.mu.RLock()
	defer s.mu.RUnlock()
	if s.db == nil {
		return errors.New("closed")
	}
	return s.db.Update(func(tx *bbolt.Tx) error { return put(tx, buckets["records"], r.ID, r) })
}
func (s *Store) GetRecord(id string) (domain.Record, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	var r domain.Record
	if s.db == nil {
		return r, errors.New("closed")
	}
	e := s.db.View(func(tx *bbolt.Tx) error { return get(tx, buckets["records"], id, &r) })
	return r, e
}
func (s *Store) SaveProfile(p domain.Profile) error {
	s.mu.RLock()
	defer s.mu.RUnlock()
	if s.db == nil {
		return errors.New("closed")
	}
	return s.db.Update(func(tx *bbolt.Tx) error { return put(tx, buckets["profiles"], p.ID, p) })
}
func (s *Store) GetProfile(id string) (domain.Profile, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	var p domain.Profile
	if s.db == nil {
		return p, errors.New("closed")
	}
	e := s.db.View(func(tx *bbolt.Tx) error { return get(tx, buckets["profiles"], id, &p) })
	return p, e
}
func (s *Store) SaveEvent(e domain.Event) error {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.db.Update(func(tx *bbolt.Tx) error { return put(tx, buckets["events"], e.ID, e) })
}
func (s *Store) SaveAudit(a domain.Audit) error {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.db.Update(func(tx *bbolt.Tx) error { return put(tx, buckets["audits"], a.ID, a) })
}
func (s *Store) ListRecords() ([]domain.Record, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	out := []domain.Record{}
	if s.db == nil {
		return out, errors.New("closed")
	}
	e := s.db.View(func(tx *bbolt.Tx) error {
		return tx.Bucket(buckets["records"]).ForEach(func(_, v []byte) error {
			var r domain.Record
			if e := json.Unmarshal(v, &r); e != nil {
				return e
			}
			out = append(out, r)
			return nil
		})
	})
	return out, e
}
func (s *Store) Health() error {
	s.mu.RLock()
	defer s.mu.RUnlock()
	if s.db == nil {
		return errors.New("closed")
	}
	return s.db.View(func(*bbolt.Tx) error { return nil })
}
func (s *Store) Snapshot() string {
	return fmt.Sprintf("%s@%s", s.path, time.Now().UTC().Format(time.RFC3339))
}
