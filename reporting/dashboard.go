package reporting

import (
	"coldchain/domain"
	"sort"
	"time"
)

type Dashboard struct {
	Generated  time.Time
	Total      int
	Warehouses []string
	Statuses   map[string]int
}

func BuildDashboard(rs []domain.Record) Dashboard {
	m := map[string]bool{}
	st := CountByStatus(rs)
	for _, r := range rs {
		m[r.WarehouseID] = true
	}
	ws := []string{}
	for w := range m {
		ws = append(ws, w)
	}
	sort.Strings(ws)
	return Dashboard{Generated: time.Now().UTC(), Total: len(rs), Warehouses: ws, Statuses: st}
}
func Recent(rs []domain.Record, since time.Time) []domain.Record {
	out := []domain.Record{}
	for _, r := range rs {
		if r.UpdatedAt.After(since) {
			out = append(out, r)
		}
	}
	sort.Slice(out, func(i, j int) bool { return out[i].UpdatedAt.After(out[j].UpdatedAt) })
	return out
}
func AverageWeight(rs []domain.Record) float64 {
	if len(rs) == 0 {
		return 0
	}
	var n float64
	for _, r := range rs {
		n += r.Weight
	}
	return n / float64(len(rs))
}
