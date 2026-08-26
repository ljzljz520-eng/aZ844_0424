package domain

import (
	"fmt"
	"time"
)

var allowedTransitions = map[string]map[string]bool{"received": {"processing": true, "cancelled": true}, "processing": {"ready": true, "cancelled": true}, "ready": {"dispatched": true, "archived": true}, "dispatched": {"archived": true}, "archived": {}}

func ValidStatus(s string) bool { _, ok := allowedTransitions[s]; return ok }
func CanTransition(from, to string) bool {
	if !ValidStatus(from) || !ValidStatus(to) {
		return false
	}
	return allowedTransitions[from][to]
}
func Transition(r *Record, to string) error {
	if r == nil {
		return fmt.Errorf("nil record")
	}
	if !CanTransition(r.Status, to) {
		return fmt.Errorf("invalid transition %s -> %s", r.Status, to)
	}
	r.Status = to
	r.UpdatedAt = nowUTC()
	return nil
}
func nowUTC() (t time.Time) { return time.Now().UTC() }
