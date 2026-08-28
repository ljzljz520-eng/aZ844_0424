package service

import (
	"coldchain/domain"
	"errors"
	"time"
)

func ArchiveRecord(s RecordStore, id, actor string) error {
	r, e := s.GetRecord(id)
	if e != nil {
		return e
	}
	if r.Status != "ready" && r.Status != "dispatched" {
		return errors.New("record not ready")
	}
	if e = domain.Transition(&r, "archived"); e != nil {
		return e
	}
	if e = s.SaveRecord(r); e != nil {
		return e
	}
	return s.SaveEvent(domain.Event{ID: id + "-archived", RecordID: id, Kind: "archived", Actor: actor, At: time.Now().UTC()})
}
func Dispatch(s RecordStore, id, actor string) error {
	r, e := s.GetRecord(id)
	if e != nil {
		return e
	}
	if e = domain.Transition(&r, "dispatched"); e != nil {
		return e
	}
	if e = s.SaveRecord(r); e != nil {
		return e
	}
	return s.SaveEvent(domain.Event{ID: id + "-dispatched", RecordID: id, Kind: "dispatched", Actor: actor, At: time.Now().UTC()})
}
