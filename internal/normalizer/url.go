package normalizer

import "strings"

func normalizeURL(u string) string {
	u = strings.TrimSpace(u)
	if u == "" {
		return ""
	}

	if !strings.HasPrefix(u, "http://") && !strings.HasPrefix(u, "https://") {
		u = "https://" + u
	}

	u = strings.ToLower(u)

	if idx := strings.Index(u[8:], "/"); idx != -1 {
		u = u[:8+idx]
	}

	return strings.TrimSuffix(u, "/")
}
