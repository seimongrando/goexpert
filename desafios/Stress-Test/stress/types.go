package stress

import "time"

type Report struct {
	TotalTime     time.Duration
	TotalRequests int
	Status200     int
	StatusOthers  map[int]int
}
