package domain

import (
	"errors"
	"fmt"
	"strings"
	"time"
)

type Record struct {
	ID, BatchID, WarehouseID, Zone string
	Status                         string
	Weight                         float64
	CreatedAt, UpdatedAt           time.Time
	Owner                          string
}
type Profile struct {
	ID, Name        string
	PermissionScope []string
	Role            string
	Active          bool
}
type Event struct {
	ID, RecordID, Kind, Actor, Detail string
	At                                time.Time
}
type Audit struct {
	ID, Actor, Action, Resource, WarehouseID string
	Allowed                                  bool
	At                                       time.Time
}

func (r Record) Validate() error {
	if strings.TrimSpace(r.ID) == "" {
		return errors.New("record id required")
	}
	if r.BatchID == "" {
		return errors.New("batch required")
	}
	if r.WarehouseID == "" {
		return errors.New("warehouse required")
	}
	if r.Zone == "" {
		return errors.New("zone required")
	}
	if r.Weight < 0 {
		return errors.New("weight negative")
	}
	return nil
}
func (r Record) Normalize() Record {
	r.ID = strings.TrimSpace(r.ID)
	r.BatchID = strings.TrimSpace(r.BatchID)
	r.WarehouseID = strings.ToUpper(strings.TrimSpace(r.WarehouseID))
	r.Zone = strings.ToUpper(strings.TrimSpace(r.Zone))
	if r.Status == "" {
		r.Status = "received"
	}
	return r
}
func (r Record) IsClosed() bool { return r.Status == "archived" || r.Status == "cancelled" }
func (r Record) Summary() string {
	return fmt.Sprintf("%s/%s %s %.2fkg", r.WarehouseID, r.BatchID, r.Zone, r.Weight)
}
func (p Profile) CanAccess(w string) bool {
	if !p.Active {
		return false
	}
	if p.Role == "admin" {
		return true
	}
	for _, s := range p.PermissionScope {
		if s == w {
			return true
		}
	}
	return false
}
func (p Profile) Validate() error {
	if p.ID == "" || p.Name == "" {
		return errors.New("profile identity required")
	}
	if p.Role == "" {
		return errors.New("role required")
	}
	return nil
}
func (e Event) Validate() error {
	if e.ID == "" || e.RecordID == "" || e.Kind == "" {
		return errors.New("event fields required")
	}
	return nil
}
func (a Audit) Validate() error {
	if a.ID == "" || a.Actor == "" || a.Action == "" {
		return errors.New("audit fields required")
	}
	return nil
}
