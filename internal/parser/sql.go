package parser

import (
	"regexp"

	"cred-sanitizer/internal/model"
)

var sqlValues = regexp.MustCompile(`\(([^)]+)\)`)

func parseSQL(raw string) []model.Credential {
	matches := sqlValues.FindAllStringSubmatch(raw, -1)
	var out []model.Credential

	for _, m := range matches {
		parts := splitLine(m[1])
		if len(parts) >= 3 {
			out = append(out, model.Credential{
				URL:      parts[0],
				Login:    parts[1],
				Password: parts[2],
			})
		}
	}
	return out
}
