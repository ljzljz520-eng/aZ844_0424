package inventory

import (
	"coldchain/domain"
	"fmt"
	"strings"
)

func ValidateZone(z string) bool {
	switch strings.ToUpper(z) {
	case "FROZEN", "CHILLED", "AMBIENT":
		return true
	default:
		return false
	}
}
func ValidateRecord(r domain.Record) error {
	if e := r.Validate(); e != nil {
		return e
	}
	if !ValidateZone(r.Zone) {
		return fmt.Errorf("invalid zone")
	}
	return nil
}
func NeedsInspection(r domain.Record) bool { return r.Status == "received" || r.Status == "processing" }
func Priority(r domain.Record) int {
	if r.Zone == "FROZEN" {
		return 3
	}
	if r.Zone == "CHILLED" {
		return 2
	}
	return 1
}
