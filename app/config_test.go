package app

import "testing"

func TestLoadConfig(t *testing.T) {
	c := LoadConfig()
	if c.DataPath == "" || c.HTTPAddr == "" {
		t.Fatal(c)
	}
}
