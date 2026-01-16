package normalizer

import (
	"strings"

	"cred-sanitizer/internal/model"
)

func Normalize(in []model.Credential) []model.Credential {
	out := make([]model.Credential, 0, len(in))

	for _, c := range in {
		c.URL = normalizeURL(c.URL)
		c.Login = normalizeLogin(c.Login)
		c.Password = normalizePassword(c.Password)

		if c.Login == "" {
			mark("login_invalido")
			continue
		}

		if c.URL != "" {
			host := strings.TrimPrefix(c.URL, "https://")
			host = strings.TrimPrefix(host, "http://")

			if strings.EqualFold(c.Login, host) &&
				!strings.Contains(c.Login, "@") &&
				isWeakPassword(c.Password) {
				continue
			}
		}

		out = append(out, c)
	}

	return out
}
