package lifecycle

import (
	"coldchain/domain"
	"fmt"
)

type Issue struct{ RecordID, Message string }

func Reconcile(rs []domain.Record) []Issue {
	out := []Issue{}
	for _, r := range rs {
		if r.ID == "" {
			out = append(out, Issue{Message: "missing id"})
		}
		if r.Status == "archived" && r.UpdatedAt.IsZero() {
			out = append(out, Issue{RecordID: r.ID, Message: "archived without timestamp"})
		}
		if r.Weight < 0 {
			out = append(out, Issue{RecordID: r.ID, Message: "negative weight"})
		}
	}
	return out
}
func Explain(i Issue) string {
	if i.RecordID == "" {
		return i.Message
	}
	return fmt.Sprintf("%s: %s", i.RecordID, i.Message)
}
func Healthy(rs []domain.Record) bool { return len(Reconcile(rs)) == 0 }
