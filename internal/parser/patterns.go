package parser

import "regexp"

var (
	urlRegex   = regexp.MustCompile(`(?i)^(https?://)?([a-z0-9.-]+\.[a-z]{2,})(/.*)?$`)
	emailRegex = regexp.MustCompile(`^[^@\s]+@[^@\s]+\.[^@\s]+$`)
)

func looksLikeURL(s string) bool {
	return urlRegex.MatchString(s)
}

func looksLikeEmail(s string) bool {
	return emailRegex.MatchString(s)
}
