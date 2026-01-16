package normalizer

import "cred-sanitizer/internal/model"

func normalizeStrict(in []model.Credential) []model.Credential {
	out := Normalize(in)

	filtered := out[:0]
	for _, c := range out {
		if len(c.Password) >= 8 {
			filtered = append(filtered, c)
		} else {
			mark("strict_password_too_short")
		}
	}
	return filtered
}
