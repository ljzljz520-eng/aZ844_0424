package service

import (
	"coldchain/domain"
	"errors"
	"time"
)

type RecordStore interface {
	SaveRecord(domain.Record) error
	GetRecord(string) (domain.Record, error)
	SaveEvent(domain.Event) error
	SaveAudit(domain.Audit) error
}
type Intake struct {
	Store RecordStore
	Clock func() time.Time
}

func (i Intake) Register(r domain.Record, actor string) error {
	r = r.Normalize()
	if e := r.Validate(); e != nil {
		return e
	}
	if i.Clock == nil {
		i.Clock = time.Now
	}
	r.CreatedAt = i.Clock().UTC()
	r.UpdatedAt = r.CreatedAt
	if e := i.Store.SaveRecord(r); e != nil {
		return e
	}
	return i.Store.SaveEvent(domain.Event{ID: r.ID + "-received", RecordID: r.ID, Kind: "received", Actor: actor, At: r.CreatedAt})
}
func (i Intake) UpdateWeight(id string, w float64, actor string) error {
	r, e := i.Store.GetRecord(id)
	if e != nil {
		return e
	}
	if r.IsClosed() {
		return errors.New("record closed")
	}
	if w < 0 {
		return errors.New("negative weight")
	}
	r.Weight = w
	r.UpdatedAt = time.Now().UTC()
	if e = i.Store.SaveRecord(r); e != nil {
		return e
	}
	return i.Store.SaveEvent(domain.Event{ID: id + "-weight", RecordID: id, Kind: "weight_updated", Actor: actor, At: r.UpdatedAt})
}
