package archive

import (
	"coldchain/domain"
	"coldchain/store"
	"errors"
	"time"
)

type Ledger struct{ Store *store.Store }

func (l Ledger) Record(a domain.Audit) error {
	if a.At.IsZero() {
		a.At = time.Now().UTC()
	}
	if e := a.Validate(); e != nil {
		return e
	}
	return l.Store.SaveAudit(a)
}
func (l Ledger) Close(r domain.Record, actor string) error {
	if !r.IsClosed() {
		return errors.New("record must be archived")
	}
	return l.Record(domain.Audit{ID: r.ID + "-close", Actor: actor, Action: "close", Resource: r.ID, WarehouseID: r.WarehouseID, Allowed: true})
}
