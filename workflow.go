package coldchain

import (
	"coldchain/domain"
	"coldchain/query"
	"coldchain/service"
	"coldchain/store"
	"context"
)

type Workflow struct {
	Store  *store.Store
	Intake service.Intake
	Query  query.Service
}

func NewWorkflow(s *store.Store) Workflow {
	return Workflow{Store: s, Intake: service.Intake{Store: s}, Query: query.Service{Store: s}}
}
func (w Workflow) ReceiveProcessArchive(ctx context.Context, r domain.Record, actor string) error {
	if e := w.Intake.Register(r, actor); e != nil {
		return e
	}
	if e := service.ProcessRecord(w.Store, r.ID, actor); e != nil {
		return e
	}
	if e := service.MarkReady(w.Store, r.ID, actor); e != nil {
		return e
	}
	return service.ArchiveRecord(w.Store, r.ID, actor)
}
