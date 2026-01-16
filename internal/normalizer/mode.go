package normalizer

import "cred-sanitizer/internal/model"

func NormalizeWithMode(in []model.Credential, strict, loose bool) []model.Credential {
	switch {
	case strict:
		return normalizeStrict(in)
	case loose:
		return normalizeLoose(in)
	default:
		return Normalize(in)
	}
}
