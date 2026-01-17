package metrics

import (
	"fmt"
)

var spinner = []rune{'|', '/', '-', '\\'}
var spinIdx int

func Progress(cur, total int) {
	if total == 0 {
		return
	}

	spin := spinner[spinIdx%len(spinner)]
	spinIdx++

	pct := (cur * 100) / total
	barSize := 30
	filled := (pct * barSize) / 100

	bar := ""
	for i := 0; i < barSize; i++ {
		if i < filled {
			bar += "█"
		} else {
			bar += "░"
		}
	}

	fmt.Printf(
		"\r\033[36m%c\033[0m \033[32m[%s]\033[0m %3d%%",
		spin, bar, pct,
	)

	if cur == total {
		fmt.Println()
	}
}
