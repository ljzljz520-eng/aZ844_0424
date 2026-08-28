package store

import (
	"coldchain/domain"
	"path/filepath"
	"testing"
)

func TestPersistenceSurvivesReopen(t *testing.T) {
	p := filepath.Join(t.TempDir(), "db")
	s, e := Open(p)
	if e != nil {
		t.Fatal(e)
	}
	if e = s.SaveRecord(domain.Record{ID: "persist", BatchID: "b", WarehouseID: "W1", Zone: "COLD"}); e != nil {
		t.Fatal(e)
	}
	s.Close()
	s, e = Open(p)
	if e != nil {
		t.Fatal(e)
	}
	defer s.Close()
	if _, e = s.GetRecord("persist"); e != nil {
		t.Fatal(e)
	}
}
