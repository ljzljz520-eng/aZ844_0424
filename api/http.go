package api

import (
	"coldchain/domain"
	"coldchain/query"
	"coldchain/service"
	"context"
	"encoding/json"
	"net/http"
)

type Server struct {
	Intake service.Intake
	Query  query.Service
}

func (s Server) Routes() *http.ServeMux {
	m := http.NewServeMux()
	m.HandleFunc("/health", s.health)
	m.HandleFunc("/records", s.records)
	return m
}
func (s Server) health(w http.ResponseWriter, _ *http.Request) { w.WriteHeader(http.StatusNoContent) }
func (s Server) records(w http.ResponseWriter, r *http.Request) {
	actor := r.Header.Get("X-Actor")
	q := r.URL.Query().Get("q")
	items, e := s.Query.Search(context.Background(), actor, q)
	if e != nil {
		http.Error(w, e.Error(), http.StatusForbidden)
		return
	}
	json.NewEncoder(w).Encode(items)
}
func DecodeRecord(r *http.Request) (domain.Record, error) {
	var v domain.Record
	e := json.NewDecoder(r.Body).Decode(&v)
	return v, e
}
