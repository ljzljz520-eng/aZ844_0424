package service

import (
	"coldchain/domain"
	"coldchain/store"
	"path/filepath"
	"testing"
)

func TestRegister(t *testing.T) {
	s, e := store.Open(filepath.Join(t.TempDir(), "db"))
	if e != nil {
		t.Fatal(e)
	}
	defer s.Close()
	i := Intake{Store: s}
	if e = i.Register(domain.Record{ID: "r", BatchID: "b", WarehouseID: "W1", Zone: "C"}, "u"); e != nil {
		t.Fatal(e)
	}
}
