package reporting

import (
	"coldchain/domain"
	"encoding/csv"
	"io"
	"strconv"
)

func WriteRecords(w io.Writer, rs []domain.Record) error {
	c := csv.NewWriter(w)
	if e := c.Write([]string{"id", "batch", "warehouse", "zone", "status", "weight"}); e != nil {
		return e
	}
	for _, r := range rs {
		if e := c.Write([]string{r.ID, r.BatchID, r.WarehouseID, r.Zone, r.Status, strconv.FormatFloat(r.Weight, 'f', 2, 64)}); e != nil {
			return e
		}
	}
	c.Flush()
	return c.Error()
}
func CountByStatus(rs []domain.Record) map[string]int {
	m := map[string]int{}
	for _, r := range rs {
		m[r.Status]++
	}
	return m
}
func CountByWarehouse(rs []domain.Record) map[string]int {
	m := map[string]int{}
	for _, r := range rs {
		m[r.WarehouseID]++
	}
	return m
}
