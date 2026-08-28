package app

import (
	"fmt"
	"os"
	"time"
)

type Config struct {
	DataPath string
	HTTPAddr string
	Shutdown time.Duration
}

func LoadConfig() Config {
	p := os.Getenv("COLDCHAIN_DB")
	if p == "" {
		p = "coldchain.db"
	}
	a := os.Getenv("COLDCHAIN_ADDR")
	if a == "" {
		a = ":8080"
	}
	return Config{DataPath: p, HTTPAddr: a, Shutdown: 5 * time.Second}
}
func (c Config) Validate() error {
	if c.DataPath == "" || c.HTTPAddr == "" {
		return fmt.Errorf("invalid config")
	}
	return nil
}
