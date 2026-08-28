package service

import (
	"coldchain/domain"
	"time"
)

type Metrics struct {
	Received, Processing, Ready, Dispatched, Archived, Cancelled int
	Weight                                                       float64
	At                                                           time.Time
}

func Collect(rs []domain.Record) Metrics {
	m := Metrics{At: time.Now().UTC()}
	for _, r := range rs {
		switch r.Status {
		case "received":
			m.Received++
		case "processing":
			m.Processing++
		case "ready":
			m.Ready++
		case "dispatched":
			m.Dispatched++
		case "archived":
			m.Archived++
		case "cancelled":
			m.Cancelled++
		}
		m.Weight += r.Weight
	}
	return m
}
func (m Metrics) Active() int { return m.Received + m.Processing + m.Ready + m.Dispatched }
