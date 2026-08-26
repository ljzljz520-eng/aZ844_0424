package query

import (
	"coldchain/domain"
	"coldchain/store"
	"context"
	"path/filepath"
	"testing"
)

func TestSearch(t *testing.T) {
	s, e := store.Open(filepath.Join(t.TempDir(), "db"))
	if e != nil {
		t.Fatal(e)
	}
	defer s.Close()
	s.SaveProfile(domain.Profile{ID: "u", Name: "U", Role: "worker", Active: true, PermissionScope: []string{"W1"}})
	s.SaveRecord(domain.Record{ID: "a", BatchID: "a", WarehouseID: "W1", Zone: "C"})
	q := Service{Store: s}
	rs, e := q.Search(context.Background(), "u", "")
	if e != nil || len(rs) != 1 {
		t.Fatalf("%v %d", e, len(rs))
	}
}
