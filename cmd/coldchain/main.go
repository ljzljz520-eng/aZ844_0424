package main

import (
	"coldchain/app"
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	c := app.LoadConfig()
	r, e := app.NewRuntime(c)
	if e != nil {
		log.Fatal(e)
	}
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()
	if e = r.Run(ctx); e != nil {
		log.Fatal(e)
	}
}
