package store

import (
	"coldchain/domain"
	"sort"
	"strings"
)

func filterRecords(in []domain.Record, zone, status string) []domain.Record {
	out := make([]domain.Record, 0, len(in))
	for _, r := range in {
		if zone != "" && r.Zone != zone {
			continue
		}
		if status != "" && r.Status != status {
			continue
		}
		out = append(out, r)
	}
	sort.Slice(out, func(i, j int) bool { return out[i].UpdatedAt.Before(out[j].UpdatedAt) })
	return out
}
func matchesBatch(r domain.Record, q string) bool {
	return q == "" || strings.Contains(strings.ToLower(r.BatchID), strings.ToLower(q))
}
func paginate(in []domain.Record, offset, limit int) []domain.Record {
	if offset < 0 {
		offset = 0
	}
	if limit <= 0 {
		limit = 50
	}
	if offset >= len(in) {
		return []domain.Record{}
	}
	end := offset + limit
	if end > len(in) {
		end = len(in)
	}
	return in[offset:end]
}
func normalizeQuery(q string) string { return strings.TrimSpace(strings.ToLower(q)) }
