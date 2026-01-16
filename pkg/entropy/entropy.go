package entropy

import "math"

func Shannon(s string) float64 {
	freq := map[rune]float64{}
	for _, r := range s {
		freq[r]++
	}

	var entropy float64
	length := float64(len(s))

	for _, count := range freq {
		p := count / length
		entropy -= p * math.Log2(p)
	}

	return entropy
}
