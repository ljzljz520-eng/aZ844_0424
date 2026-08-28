package domain

import "testing"

func TestRecordNormalize(t *testing.T) {
	r := Record{ID: " x ", BatchID: "b", WarehouseID: " wh1 ", Zone: "cold"}.Normalize()
	if r.WarehouseID != "WH1" {
		t.Fatal(r)
	}
}
func TestTransition(t *testing.T) {
	r := Record{ID: "1", Status: "received"}
	if e := Transition(&r, "processing"); e != nil {
		t.Fatal(e)
	}
}
