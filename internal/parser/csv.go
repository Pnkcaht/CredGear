package parser

import (
	"encoding/csv"
	"strings"

	"cred-sanitizer/internal/model"
)

func parseCSV(raw string) []model.Credential {
	r := csv.NewReader(strings.NewReader(raw))
	records, _ := r.ReadAll()

	var out []model.Credential
	for _, row := range records {
		if len(row) < 2 {
			continue
		}
		out = append(out, model.Credential{
			URL:      row[0],
			Login:    row[1],
			Password: row[2],
		})
	}
	return out
}
