package normalizer

import "strings"

func normalizePassword(s string) string {
	s = strings.TrimSpace(s)

	if s == "" || s == "?" {
		return ""
	}

	if len(s) < 6 {
		return ""
	}

	allDigits := true
	for _, r := range s {
		if r < '0' || r > '9' {
			allDigits = false
			break
		}
	}
	if allDigits && len(s) < 8 {
		return ""
	}

	hasAlnum := false
	for _, r := range s {
		if (r >= 'a' && r <= 'z') ||
			(r >= 'A' && r <= 'Z') ||
			(r >= '0' && r <= '9') {
			hasAlnum = true
			break
		}
	}
	if !hasAlnum {
		return ""
	}

	return s
}

func isWeakPassword(s string) bool {
	hasLetter := false
	hasDigit := false

	for _, r := range s {
		if (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') {
			hasLetter = true
		}
		if r >= '0' && r <= '9' {
			hasDigit = true
		}
	}

	if !hasLetter && hasDigit {
		return true
	}

	if len(s) < 6 {
		return true
	}

	return false
}
