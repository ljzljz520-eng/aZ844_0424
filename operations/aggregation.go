package operations

import (
	"coldchain/domain"
	"sort"
)

type WarehouseTotal struct {
	Warehouse string
	Records   int
	Weight    float64
}

func Aggregate(rs []domain.Record) []WarehouseTotal {
	m := map[string]*WarehouseTotal{}
	for _, r := range rs {
		v := m[r.WarehouseID]
		if v == nil {
			v = &WarehouseTotal{Warehouse: r.WarehouseID}
			m[r.WarehouseID] = v
		}
		v.Records++
		v.Weight += r.Weight
	}
	out := []WarehouseTotal{}
	for _, v := range m {
		out = append(out, *v)
	}
	sort.Slice(out, func(i, j int) bool { return out[i].Warehouse < out[j].Warehouse })
	return out
}
func TotalWeight(rs []domain.Record) float64 {
	var n float64
	for _, r := range rs {
		n += r.Weight
	}
	return n
}
func TotalRecords(rs []domain.Record) int { return len(rs) }
func Zones(rs []domain.Record) []string {
	m := map[string]bool{}
	for _, r := range rs {
		m[r.Zone] = true
	}
	out := []string{}
	for z := range m {
		out = append(out, z)
	}
	sort.Strings(out)
	return out
}
func Statuses(rs []domain.Record) []string {
	m := map[string]bool{}
	for _, r := range rs {
		m[r.Status] = true
	}
	out := []string{}
	for s := range m {
		out = append(out, s)
	}
	sort.Strings(out)
	return out
}
