package coldchain

import (
	"coldchain/domain"
	"coldchain/store"
	"context"
	"path/filepath"
	"testing"
)

func TestWorkflowOne(t *testing.T) {
	s, e := store.Open(filepath.Join(t.TempDir(), "db"))
	if e != nil {
		t.Fatal(e)
	}
	defer s.Close()
	w := NewWorkflow(s)
	e = w.Intake.Register(domain.Record{ID: "w1", BatchID: "b", WarehouseID: "W1", Zone: "C", Weight: 1}, "u")
	if e != nil {
		t.Fatal(e)
	}
}
func TestWorkflowTwo(t *testing.T) {
	if context.Background() == nil {
		t.Fatal("context")
	}
}
func TestWorkflowThree(t *testing.T) {
	s, e := store.Open(filepath.Join(t.TempDir(), "db"))
	if e != nil {
		t.Fatal(e)
	}
	s.Close()
}
func TestBusinessChain24(t *testing.T) {
	s, e := store.Open(filepath.Join(t.TempDir(), "db"))
	if e != nil {
		t.Fatal(e)
	}
	defer s.Close()
	w := NewWorkflow(s)
	if e = w.ReceiveProcessArchive(context.Background(), domain.Record{ID: "chain", BatchID: "b", WarehouseID: "W1", Zone: "C", Weight: 2}, "u"); e != nil {
		t.Fatal(e)
	}
}
