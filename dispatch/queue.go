package dispatch

import (
	"coldchain/domain"
	"errors"
	"sync"
)

type Queue struct {
	mu    sync.Mutex
	items []domain.Record
}

func (q *Queue) Push(r domain.Record) error {
	q.mu.Lock()
	defer q.mu.Unlock()
	if r.ID == "" {
		return errors.New("id")
	}
	q.items = append(q.items, r)
	return nil
}
func (q *Queue) Pop() (domain.Record, bool) {
	q.mu.Lock()
	defer q.mu.Unlock()
	if len(q.items) == 0 {
		return domain.Record{}, false
	}
	r := q.items[0]
	q.items = q.items[1:]
	return r, true
}
func (q *Queue) Len() int { q.mu.Lock(); defer q.mu.Unlock(); return len(q.items) }
func (q *Queue) Drain() []domain.Record {
	q.mu.Lock()
	defer q.mu.Unlock()
	out := append([]domain.Record{}, q.items...)
	q.items = nil
	return out
}
func (q *Queue) Contains(id string) bool {
	q.mu.Lock()
	defer q.mu.Unlock()
	for _, r := range q.items {
		if r.ID == id {
			return true
		}
	}
	return false
}
