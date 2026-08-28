package lifecycle

import (
	"coldchain/domain"
	"time"
)

type Retention struct{ ArchiveAfter time.Duration }

func (r Retention) Eligible(v domain.Record, now time.Time) bool {
	return v.Status == "archived" && !v.UpdatedAt.Add(r.ArchiveAfter).After(now)
}
func Expired(rs []domain.Record, r Retention, now time.Time) []domain.Record {
	out := []domain.Record{}
	for _, v := range rs {
		if r.Eligible(v, now) {
			out = append(out, v)
		}
	}
	return out
}
func MarkExpired(rs []domain.Record, r Retention, now time.Time) []domain.Record {
	out := append([]domain.Record{}, rs...)
	for i := range out {
		if r.Eligible(out[i], now) {
			out[i].Status = "expired"
		}
	}
	return out
}
