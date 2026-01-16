package integration

import (
	"cred-sanitizer/internal/deduplicator"
	"cred-sanitizer/internal/normalizer"
	"cred-sanitizer/internal/output"
	"cred-sanitizer/internal/parser"
)

func Run(raw string, cfg Config) []byte {
	parsed := parser.ParseWithFormat(raw, cfg.Format)
	normalized := normalizer.NormalizeWithMode(parsed, cfg.Strict, cfg.Loose)
	dedup := deduplicator.Deduplicate(normalized)

	if cfg.Metrics {
		normalizer.PrintMetrics()
	}

	data, err := output.JSON(dedup)
	if err != nil {
		panic(err)
	}

	return data
}
