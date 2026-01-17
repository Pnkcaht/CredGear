package normalizer

import (
	"fmt"
	"time"
)

const (
	red   = "\033[31m"
	reset = "\033[0m"
)

var (
	totalIn   int
	totalOut  int
	totalDrop int
	startTime time.Time

	dropReasons = map[string]int{}
)

func initMetrics() {
	if startTime.IsZero() {
		startTime = time.Now()
	}
}

func incIn()   { totalIn++ }
func incOut()  { totalOut++ }
func incDrop() { totalDrop++ }

func mark(reason string) {
	dropReasons[reason]++
}

func progress(cur, total int) {
	if total == 0 {
		return
	}
	pct := (cur * 100) / total
	fmt.Printf("\r[%3d%%] Sanitizing...", pct)
	if cur == total {
		fmt.Println()
	}
}

func PrintMetrics() {
	elapsed := time.Since(startTime)

	fmt.Println()
	fmt.Println(red + "── Metrics ─────────────────────" + reset)
	fmt.Println(" Input credentials :", totalIn)
	fmt.Println(" Valid credentials :", totalOut)
	fmt.Println(" Discarded         :", totalDrop)
	fmt.Println(" Time elapsed      :", elapsed)

	if len(dropReasons) > 0 {
		fmt.Println("\n Drop reasons:")
		for k, v := range dropReasons {
			fmt.Printf("  - %-25s %d\n", k, v)
		}
	}

	fmt.Println("────────────────────────────────")
}

func resetMetrics() {
	totalIn = 0
	totalOut = 0
	totalDrop = 0
	dropReasons = map[string]int{}
	startTime = time.Now()
}
