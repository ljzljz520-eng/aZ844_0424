package validation

import (
	"coldchain/domain"
	"fmt"
	"strings"
)

type Result struct {
	OK     bool
	Errors []string
}

func CheckRecord(r domain.Record) Result {
	errs := []string{}
	if e := r.Validate(); e != nil {
		errs = append(errs, e.Error())
	}
	if strings.TrimSpace(r.Zone) == "" {
		errs = append(errs, "zone missing")
	}
	return Result{OK: len(errs) == 0, Errors: errs}
}
func CheckBatch(rs []domain.Record) Result {
	errs := []string{}
	seen := map[string]bool{}
	for _, r := range rs {
		if seen[r.ID] {
			errs = append(errs, fmt.Sprintf("duplicate %s", r.ID))
		}
		seen[r.ID] = true
	}
	return Result{OK: len(errs) == 0, Errors: errs}
}
func RequireStatus(r domain.Record, status string) error {
	if r.Status != status {
		return fmt.Errorf("expected %s got %s", status, r.Status)
	}
	return nil
}
func IsWarehouseCode(s string) bool { return len(s) >= 2 && len(s) <= 8 && strings.ToUpper(s) == s }
