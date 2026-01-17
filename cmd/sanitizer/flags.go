package main

import (
	"flag"
	"fmt"
	"io"
	"os"

	"cred-sanitizer/internal/integration"
)

func runFlags() {
	var (
		strict  = flag.Bool("strict", false, "ativa regras rígidas")
		loose   = flag.Bool("loose", false, "regras permissivas")
		metrics = flag.Bool("metrics", false, "exibe métricas")
		format  = flag.String("format", "auto", "auto|txt|csv|sql")
		outFile = flag.String("out", "", "arquivo de saída (json)")
	)

	flag.Parse()

	if flag.NArg() == 0 {
		fmt.Fprintln(os.Stderr, "uso: sanitizer [flags] <arquivo>")
		os.Exit(1)
	}

	inputFile := flag.Arg(0)
	data, err := os.ReadFile(inputFile)
	if err != nil {
		panic(err)
	}

	cfg := integration.Config{
		Strict:  *strict,
		Loose:   *loose,
		Metrics: *metrics,
		Format:  *format,
	}

	result := integration.Run(string(data), cfg)

	var out io.Writer = os.Stdout
	if *outFile != "" {
		f, err := os.Create(*outFile)
		if err != nil {
			panic(err)
		}
		defer f.Close()
		out = f
	}

	out.Write(result)
}
