package query

import (
	"coldchain/domain"
	"sort"
)

type Report struct {
	Total    int
	ByStatus map[string]int
	ByZone   map[string]int
	Weight   float64
}

func BuildReport(rs []domain.Record) Report {
	r := Report{ByStatus: map[string]int{}, ByZone: map[string]int{}}
	for _, v := range rs {
		r.Total++
		r.ByStatus[v.Status]++
		r.ByZone[v.Zone]++
		r.Weight += v.Weight
	}
	return r
}
func TopZones(rs []domain.Record, n int) []string {
	m := map[string]int{}
	for _, r := range rs {
		m[r.Zone]++
	}
	type pair struct {
		k string
		v int
	}
	a := []pair{}
	for k, v := range m {
		a = append(a, pair{k, v})
	}
	sort.Slice(a, func(i, j int) bool { return a[i].v > a[j].v })
	if n > len(a) {
		n = len(a)
	}
	out := []string{}
	for _, p := range a[:n] {
		out = append(out, p.k)
	}
	return out
}
