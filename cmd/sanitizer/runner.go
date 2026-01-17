package main

import (
	"fmt"
	"io"
	"os"

	"cred-sanitizer/internal/integration"
)

func runSanitizer(input string, cfg integration.Config) {
	result := integration.Run(input, cfg)
	fmt.Print(string(result))
}

func runFromFile(path string, cfg integration.Config) {
	data, err := os.ReadFile(path)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	runSanitizer(string(data), cfg)
}

func runFromStdin(cfg integration.Config) {
	data, err := io.ReadAll(os.Stdin)
	if err != nil {
		fmt.Println("Error reading stdin:", err)
		return
	}
	runSanitizer(string(data), cfg)
}
