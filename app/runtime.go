package app

import (
	"coldchain/api"
	"coldchain/auth"
	"coldchain/query"
	"coldchain/service"
	"coldchain/store"
	"context"
	"net/http"
)

type Runtime struct {
	Config Config
	Store  *store.Store
	Server *http.Server
}

func NewRuntime(c Config) (*Runtime, error) {
	s, e := store.Open(c.DataPath)
	if e != nil {
		return nil, e
	}
	p := auth.Policy{Store: s}
	q := query.Service{Store: s, Policy: p}
	in := service.Intake{Store: s}
	srv := &http.Server{Addr: c.HTTPAddr, Handler: (api.Server{Intake: in, Query: q}).Routes()}
	return &Runtime{Config: c, Store: s, Server: srv}, nil
}
func (r *Runtime) Run(ctx context.Context) error {
	go func() { <-ctx.Done(); r.Server.Shutdown(context.Background()); r.Store.Close() }()
	e := r.Server.ListenAndServe()
	if e == http.ErrServerClosed {
		return nil
	}
	return e
}
