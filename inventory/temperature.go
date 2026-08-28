package inventory

import (
	"coldchain/domain"
	"fmt"
	"time"
)

type Reading struct {
	RecordID string
	Celsius  float64
	At       time.Time
}
type Monitor struct {
	Min, Max float64
	Readings map[string][]Reading
}

func NewMonitor(min, max float64) *Monitor {
	return &Monitor{Min: min, Max: max, Readings: map[string][]Reading{}}
}
func (m *Monitor) Add(r Reading) error {
	if r.RecordID == "" {
		return fmt.Errorf("record")
	}
	if r.At.IsZero() {
		r.At = time.Now().UTC()
	}
	m.Readings[r.RecordID] = append(m.Readings[r.RecordID], r)
	return nil
}
func (m Monitor) Safe(c float64) bool { return c >= m.Min && c <= m.Max }
func (m Monitor) Latest(id string) (Reading, bool) {
	a := m.Readings[id]
	if len(a) == 0 {
		return Reading{}, false
	}
	return a[len(a)-1], true
}
func (m Monitor) Violations(id string) []Reading {
	out := []Reading{}
	for _, r := range m.Readings[id] {
		if !m.Safe(r.Celsius) {
			out = append(out, r)
		}
	}
	return out
}
func (m Monitor) Status(id string) string {
	r, ok := m.Latest(id)
	if !ok {
		return "unknown"
	}
	if m.Safe(r.Celsius) {
		return "safe"
	}
	return "alert"
}
func WeightByZone(rs []domain.Record) map[string]float64 {
	m := map[string]float64{}
	for _, r := range rs {
		m[r.Zone] += r.Weight
	}
	return m
}
