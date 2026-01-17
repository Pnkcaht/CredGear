package integration

import "time"

type Metrics struct {
	LinesRead   int
	ParsedCreds int
	Discarded   int
	Start       time.Time
	End         time.Time
}

func NewMetrics() *Metrics {
	return &Metrics{
		Start: time.Now(),
	}
}

func (m *Metrics) Finish() {
	m.End = time.Now()
}

func (m *Metrics) Duration() time.Duration {
	return m.End.Sub(m.Start)
}
