package normalizer

import "strings"

func normalizeLogin(s string) string {
	s = strings.TrimSpace(s)
	s = strings.Trim(s, "/")

	if s == "" {
		return ""
	}

	ls := strings.ToLower(s)

	if strings.Contains(ls, "://") || strings.Contains(ls, "/") {
		return ""
	}

	if strings.HasPrefix(ls, "www.") {
		return ""
	}

	if strings.Contains(ls, "@") {
		parts := strings.Split(ls, "@")
		if len(parts) != 2 || parts[0] == "" || parts[1] == "" {
			return ""
		}
		if !strings.Contains(parts[1], ".") {
			return ""
		}
	}

	return s
}
