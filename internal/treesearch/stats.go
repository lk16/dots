package treesearch

import (
	"time"
)

type Stats struct {
	Nodes     uint64
	StartTime time.Time
	Duration  time.Duration
}

func (s *Stats) StartClock() {
	s.StartTime = time.Now()
}

func (s *Stats) StopClock() {
	s.Duration += time.Now().Sub(s.StartTime)
}

func (s *Stats) NodesPerSecond() float64 {

	duration := s.Duration.Seconds()

	if duration == 0.0 {
		return 0.0
	}

	return float64(s.Nodes) / duration
}

func (s *Stats) Reset() {
	s.Duration = 0
	s.Nodes = 0
}

func (s *Stats) Add(other Stats) {
	s.Nodes += other.Nodes
	s.Duration += other.Duration
}
