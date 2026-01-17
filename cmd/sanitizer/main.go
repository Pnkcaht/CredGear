package main

import "os"

func main() {
	// sem argumentos → modo interativo
	if len(os.Args) == 1 {
		runInteractive()
		return
	}

	// com argumentos → modo CLI tradicional
	runFlags()
}
