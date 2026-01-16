package deduplicator

import (
	"cred-sanitizer/internal/model"
)

func Deduplicate(in []model.Credential) []model.Credential {
	seen := make(map[string]model.Credential)
	out := make([]model.Credential, 0)

	for _, c := range in {
		key := c.Login + "\x00" + c.Password

		existing, ok := seen[key]
		if !ok {
			seen[key] = c
			continue
		}

		// mantém URL mais curta
		if len(existing.URL) == 0 || len(c.URL) < len(existing.URL) {
			seen[key] = c
		}
	}

	for _, v := range seen {
		out = append(out, v)
	}

	return out
}
