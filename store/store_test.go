package store

import (
	"coldchain/domain"
	"path/filepath"
	"testing"
)

func TestStoreRoundTrip(t *testing.T) {
	p := filepath.Join(t.TempDir(), "db")
	s, e := Open(p)
	if e != nil {
		t.Fatal(e)
	}
	defer s.Close()
	r := domain.Record{ID: "r1", BatchID: "b", WarehouseID: "W1", Zone: "FROZEN"}
	if e = s.SaveRecord(r); e != nil {
		t.Fatal(e)
	}
	if _, e = s.GetRecord("r1"); e != nil {
		t.Fatal(e)
	}
}
