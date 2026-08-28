package operations

import (
	"coldchain/domain"
	"fmt"
	"strings"
)

func CheckID(r domain.Record) error {
	if strings.TrimSpace(r.ID) == "" {
		return fmt.Errorf("id required")
	}
	return nil
}
func CheckBatch(r domain.Record) error {
	if strings.TrimSpace(r.BatchID) == "" {
		return fmt.Errorf("batch required")
	}
	return nil
}
func CheckWarehouse(r domain.Record) error {
	if strings.TrimSpace(r.WarehouseID) == "" {
		return fmt.Errorf("warehouse required")
	}
	return nil
}
func CheckZone(r domain.Record) error {
	if r.Zone == "" {
		return fmt.Errorf("zone required")
	}
	return nil
}
func CheckWeight(r domain.Record) error {
	if r.Weight < 0 {
		return fmt.Errorf("negative weight")
	}
	return nil
}
func ValidateIntake(r domain.Record) []error {
	return []error{CheckID(r), CheckBatch(r), CheckWarehouse(r), CheckZone(r), CheckWeight(r)}
}
func HasErrors(es []error) bool {
	for _, e := range es {
		if e != nil {
			return true
		}
	}
	return false
}
func NormalizeStatus(r domain.Record) domain.Record {
	if r.Status == "" {
		r.Status = "received"
	}
	return r
}
func IsColdZone(z string) bool          { return z == "FROZEN" || z == "CHILLED" }
func RequiresSeal(r domain.Record) bool { return IsColdZone(r.Zone) && r.Weight > 50 }
func SealCode(r domain.Record) string {
	if !RequiresSeal(r) {
		return ""
	}
	return "SEAL-" + r.ID
}
func IntakeLabel(r domain.Record) string {
	return fmt.Sprintf("%s-%s-%s", r.WarehouseID, r.BatchID, r.Zone)
}
