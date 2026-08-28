package operations

import (
	"coldchain/domain"
	"fmt"
)

func IsTerminal(r domain.Record) bool {
	return r.Status == "archived" || r.Status == "cancelled" || r.Status == "expired"
}
func IsInTransit(r domain.Record) bool {
	return r.Status == "processing" || r.Status == "ready" || r.Status == "dispatched"
}
func CanEdit(r domain.Record) bool     { return !IsTerminal(r) }
func CanDispatch(r domain.Record) bool { return r.Status == "ready" }
func CanArchive(r domain.Record) bool  { return r.Status == "ready" || r.Status == "dispatched" }
func NextStatuses(s string) []string {
	switch s {
	case "received":
		return []string{"processing", "cancelled"}
	case "processing":
		return []string{"ready", "cancelled"}
	case "ready":
		return []string{"dispatched", "archived"}
	case "dispatched":
		return []string{"archived"}
	default:
		return []string{}
	}
}
func RequireTransition(r domain.Record, to string) error {
	if !domain.CanTransition(r.Status, to) {
		return fmt.Errorf("transition denied")
	}
	return nil
}
func StatusRank(s string) int {
	for i, v := range []string{"received", "processing", "ready", "dispatched", "archived"} {
		if s == v {
			return i
		}
	}
	return -1
}
func CompareStatus(a, b domain.Record) int {
	if StatusRank(a.Status) < StatusRank(b.Status) {
		return -1
	}
	if StatusRank(a.Status) > StatusRank(b.Status) {
		return 1
	}
	return 0
}
