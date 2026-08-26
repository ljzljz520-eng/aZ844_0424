package auth

import (
	"coldchain/domain"
	"testing"
)

type fakeProfiles struct{ p domain.Profile }

func (f fakeProfiles) GetProfile(string) (domain.Profile, error) { return f.p, nil }
func TestAuthorize(t *testing.T) {
	p := Policy{Store: fakeProfiles{domain.Profile{ID: "u", Role: "worker", Active: true, PermissionScope: []string{"W1"}}}}
	if _, e := p.Authorize(Request{Actor: "u", Warehouse: "W2"}); e == nil {
		t.Fatal("expected forbidden")
	}
}
