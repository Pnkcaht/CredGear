package normalizer

import "fmt"

var metrics = map[string]int{}

func mark(reason string) {
	metrics[reason]++
}

func PrintMetrics() {
	fmt.Println("== MÉTRICAS ==")
	for k, v := range metrics {
		fmt.Printf("%s: %d\n", k, v)
	}
}
