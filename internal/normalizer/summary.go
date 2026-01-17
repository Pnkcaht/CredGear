package normalizer

import "time"

func ResetMetrics() {
	totalIn = 0
	totalOut = 0
	totalDrop = 0
	startTime = time.Time{}
	dropReasons = map[string]int{}
}
