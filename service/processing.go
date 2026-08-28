package service

import (
	"coldchain/domain"
	"errors"
	"time"
)

func ProcessRecord(s RecordStore, id, actor string) error {
	r, e := s.GetRecord(id)
	if e != nil {
		return e
	}
	if e = domain.Transition(&r, "processing"); e != nil {
		return e
	}
	if e = s.SaveRecord(r); e != nil {
		return e
	}
	return s.SaveEvent(domain.Event{ID: id + "-processing", RecordID: id, Kind: "processing", Actor: actor, At: time.Now().UTC()})
}
func MarkReady(s RecordStore, id, actor string) error {
	r, e := s.GetRecord(id)
	if e != nil {
		return e
	}
	if r.Weight == 0 {
		return errors.New("weight required")
	}
	if e = domain.Transition(&r, "ready"); e != nil {
		return e
	}
	if e = s.SaveRecord(r); e != nil {
		return e
	}
	return s.SaveEvent(domain.Event{ID: id + "-ready", RecordID: id, Kind: "ready", Actor: actor, At: time.Now().UTC()})
}
func Cancel(s RecordStore, id, actor string) error {
	r, e := s.GetRecord(id)
	if e != nil {
		return e
	}
	if r.IsClosed() {
		return errors.New("already closed")
	}
	r.Status = "cancelled"
	r.UpdatedAt = time.Now().UTC()
	if e = s.SaveRecord(r); e != nil {
		return e
	}
	return s.SaveEvent(domain.Event{ID: id + "-cancelled", RecordID: id, Kind: "cancelled", Actor: actor, At: r.UpdatedAt})
}
