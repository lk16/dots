package treesearch

import (
	"fmt"
	"time"
)

type Stats struct {
	Nodes     uint64
	StartTime time.Time
	Duration  time.Duration
}

func NewStats() Stats {
	return Stats{}
}

func (s *Stats) StartClock() {
	s.StartTime = time.Now()
}

func (s *Stats) StopClock() {
	s.Duration += time.Since(s.StartTime)
}

func (s *Stats) NodesPerSecond() uint64 {

	duration := s.Duration.Seconds()

	if duration == 0.0 {
		return 0
	}

	return uint64(float64(s.Nodes) / duration)
}

func (s *Stats) Reset() {
	s.Duration = 0
	s.Nodes = 0
}

func (s *Stats) Add(other Stats) {
	s.Nodes += other.Nodes
	s.Duration += other.Duration
}

func (s Stats) String() string {
	return fmt.Sprintf("%5s nodes in %.3f seconds = %5s nodes/second",
		FormatBigNumber(s.Nodes), s.Duration.Seconds(), FormatBigNumber(s.NodesPerSecond()))
}

func FormatBigNumber(number uint64) string {

	if number < 1000 {
		return fmt.Sprintf("%d", number)
	}

	n := float64(number)
	index := 0

	for n >= 1000.0 {
		n /= 1000.0
		index++
	}

	postfixes := []string{"", "K", "M", "G", "T", "P", "E"}

	if n < 10.0 {
		return fmt.Sprintf("%.2f%s", n, postfixes[index])
	}

	if n < 100.0 {
		return fmt.Sprintf("%.1f%s", n, postfixes[index])
	}

	return fmt.Sprintf("%.0f%s", n, postfixes[index])
}
