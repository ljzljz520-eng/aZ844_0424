package auth

import (
	"coldchain/domain"
	"errors"
)

type Request struct {
	Actor     string
	Warehouse string
	Action    string
}
type Policy struct {
	Store interface {
		GetProfile(string) (domain.Profile, error)
	}
}

func (p Policy) Authorize(req Request) (domain.Profile, error) {
	if req.Actor == "" || req.Warehouse == "" {
		return domain.Profile{}, errors.New("actor and warehouse required")
	}
	prof, e := p.Store.GetProfile(req.Actor)
	if e != nil {
		return prof, e
	}
	if !prof.CanAccess(req.Warehouse) {
		return prof, errors.New("forbidden")
	}
	return prof, nil
}
func (p Policy) CanRead(prof domain.Profile, warehouse string) bool { return prof.CanAccess(warehouse) }
func Scope(profile domain.Profile) []string {
	out := append([]string{}, profile.PermissionScope...)
	return out
}
