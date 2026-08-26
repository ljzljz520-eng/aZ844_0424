package query

import (
	"coldchain/domain"
	"coldchain/store"
	"context"
	"path/filepath"
	"testing"
)

func TestUnauthorizedWarehouseHidden(t *testing.T) {
	s, e := store.Open(filepath.Join(t.TempDir(), "db"))
	if e != nil {
		t.Fatal(e)
	}
	defer s.Close()
	s.SaveProfile(domain.Profile{ID: "worker", Name: "Worker", Role: "dispatcher", Active: true, PermissionScope: []string{"W1"}})
	s.SaveRecord(domain.Record{ID: "one", BatchID: "b1", WarehouseID: "W1", Zone: "C"})
	s.SaveRecord(domain.Record{ID: "two", BatchID: "b2", WarehouseID: "W2", Zone: "F"})
	rs, e := (Service{Store: s}).Search(context.Background(), "worker", "")
	if e != nil {
		t.Fatal(e)
	}
	if len(rs) != 1 {
		t.Fatalf("unauthorized records visible: %d", len(rs))
	}
}
