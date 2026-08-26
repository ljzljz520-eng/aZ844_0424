package dispatch

import (
	"coldchain/domain"
	"fmt"
	"sort"
	"time"
)

type Slot struct {
	ID, Warehouse, Route string
	Start, End           time.Time
	Capacity             float64
	Used                 float64
}
type Planner struct{ Slots []Slot }

func PlanRecord(r domain.Record, slots []Slot) (Slot, error) {
	p := Planner{Slots: slots}
	a := p.Available(r.WarehouseID, r.Weight)
	if len(a) == 0 {
		return Slot{}, fmt.Errorf("no capacity")
	}
	return a[0], nil
}

func (p *Planner) Add(s Slot) error {
	if s.End.Before(s.Start) {
		return fmt.Errorf("invalid slot")
	}
	p.Slots = append(p.Slots, s)
	return nil
}
func (p Planner) Available(w string, weight float64) []Slot {
	out := []Slot{}
	for _, s := range p.Slots {
		if s.Warehouse == w && s.Capacity-s.Used >= weight && time.Now().Before(s.End) {
			out = append(out, s)
		}
	}
	sort.Slice(out, func(i, j int) bool { return out[i].Start.Before(out[j].Start) })
	return out
}
func (p *Planner) Reserve(id string, weight float64) error {
	for i := range p.Slots {
		if p.Slots[i].ID == id {
			if p.Slots[i].Capacity-p.Slots[i].Used < weight {
				return fmt.Errorf("capacity")
			}
			p.Slots[i].Used += weight
			return nil
		}
	}
	return fmt.Errorf("slot not found")
}
func (p Planner) Utilization() float64 {
	var c, u float64
	for _, s := range p.Slots {
		c += s.Capacity
		u += s.Used
	}
	if c == 0 {
		return 0
	}
	return u / c
}
func (p Planner) ForDay(day time.Time) []Slot {
	out := []Slot{}
	for _, s := range p.Slots {
		if s.Start.Year() == day.Year() && s.Start.YearDay() == day.YearDay() {
			out = append(out, s)
		}
	}
	return out
}
