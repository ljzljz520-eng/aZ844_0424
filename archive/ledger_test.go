package archive

import (
	"coldchain/domain"
	"coldchain/store"
	"path/filepath"
	"testing"
)

func TestLedger(t *testing.T) {
	s, e := store.Open(filepath.Join(t.TempDir(), "db"))
	if e != nil {
		t.Fatal(e)
	}
	defer s.Close()
	l := Ledger{Store: s}
	if e = l.Record(domain.Audit{ID: "a", Actor: "u", Action: "x"}); e != nil {
		t.Fatal(e)
	}
}
