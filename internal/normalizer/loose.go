package normalizer

import "cred-sanitizer/internal/model"

func normalizeLoose(in []model.Credential) []model.Credential {
	// loose = só normaliza, quase não descarta
	return Normalize(in)
}
