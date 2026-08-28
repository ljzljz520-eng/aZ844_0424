package query

import (
	"coldchain/auth"
	"coldchain/domain"
	"coldchain/store"
	"context"
	"errors"
)

type Service struct {
	Store  *store.Store
	Policy auth.Policy
}

func (s Service) Search(ctx context.Context, actor, q string) ([]domain.Record, error) {
	if ctx == nil {
		return nil, errors.New("context required")
	}
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
	}
	p, e := s.Store.GetProfile(actor)
	if e != nil {
		return nil, e
	}
	if !p.Active {
		return nil, errors.New("inactive")
	}
	records, e := s.Store.QueryRecords(q, "", "")
	if e != nil {
		return nil, e
	}
	out := []domain.Record{}
	for _, r := range records {
		out = append(out, r)
	}
	return out, nil
}
func (s Service) SearchWarehouse(ctx context.Context, actor, warehouse, q string) ([]domain.Record, error) {
	p, e := s.Store.GetProfile(actor)
	if e != nil {
		return nil, e
	}
	if !p.CanAccess(warehouse) {
		return nil, errors.New("forbidden")
	}
	records, e := s.Store.QueryRecords(q, warehouse, "")
	if e != nil {
		return nil, e
	}
	return records, nil
}
