package normalizer

import (
	"strings"

	"cred-sanitizer/internal/model"
)

func Normalize(in []model.Credential) []model.Credential {
	initMetrics()
	resetMetrics()

	out := make([]model.Credential, 0, len(in))
	total := len(in)

	for i, c := range in {
		progress(i+1, total)
		incIn()

		c.URL = normalizeURL(c.URL)
		c.Login = normalizeLogin(c.Login)
		c.Password = normalizePassword(c.Password)

		if c.Login == "" {
			incDrop()
			mark("invalid_login")
			continue
		}

		if isGarbageLogin(c.Login) {
			incDrop()
			mark("garbage_login")
			continue
		}

		if c.Password == "" {
			incDrop()
			mark("empty_password")
			continue
		}

		if c.URL != "" {
			host := strings.TrimPrefix(c.URL, "https://")
			host = strings.TrimPrefix(host, "http://")

			if strings.EqualFold(c.Login, host) &&
				!strings.Contains(c.Login, "@") &&
				isWeakPassword(c.Password) {
				incDrop()
				mark("login_equals_host")
				continue
			}
		}

		incOut()
		out = append(out, c)
	}

	return out
}

func isGarbageLogin(s string) bool {
	switch strings.ToLower(s) {
	case "http", "https", "www":
		return true
	}
	return false
}
