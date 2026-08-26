package auth

import (
	"sync"
	"time"
)

type Session struct {
	ID, Actor string
	ExpiresAt time.Time
}
type Sessions struct {
	mu    sync.RWMutex
	items map[string]Session
}

func NewSessions() *Sessions      { return &Sessions{items: map[string]Session{}} }
func (s *Sessions) Put(v Session) { s.mu.Lock(); defer s.mu.Unlock(); s.items[v.ID] = v }
func (s *Sessions) Get(id string) (Session, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	v, ok := s.items[id]
	if !ok || time.Now().After(v.ExpiresAt) {
		return Session{}, false
	}
	return v, true
}
func (s *Sessions) Revoke(id string) { s.mu.Lock(); defer s.mu.Unlock(); delete(s.items, id) }
