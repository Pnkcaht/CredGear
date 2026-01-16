package parser

import (
	"strings"
	"unicode"

	"cred-sanitizer/internal/model"
)

var separators = []rune{':', '|', ';', ' '}

func Parse(raw string) []model.Credential {
	lines := strings.Split(raw, "\n")
	out := make([]model.Credential, 0, len(lines))

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		tokens := splitLine(line)
		if len(tokens) < 2 {
			continue
		}

		var url, login, password string

		// identifica URL se existir
		for _, t := range tokens {
			if url == "" && looksLikeURL(t) {
				url = t
			}
		}

		// regra do teste:
		// primeiro token significativo = login
		// último token significativo = senha
		login = firstValid(tokens)
		password = lastValid(tokens)

		// ignora senha lixo óbvia
		if password == "?" {
			continue
		}

		if login == "" || password == "" {
			continue
		}

		out = append(out, model.Credential{
			URL:      url,
			Login:    login,
			Password: password,
		})
	}

	return out
}

func splitLine(line string) []string {
	out := []string{}
	current := strings.Builder{}

	for _, r := range line {
		if isSeparator(r) {
			flush(&out, &current)
			continue
		}
		current.WriteRune(r)
	}
	flush(&out, &current)

	return out
}

func flush(out *[]string, b *strings.Builder) {
	s := strings.TrimSpace(b.String())
	s = strings.Trim(s, "/") // remove / no fim (caso Santander)

	if s != "" && isMeaningful(s) {
		*out = append(*out, s)
	}
	b.Reset()
}

func isSeparator(r rune) bool {
	for _, s := range separators {
		if r == s {
			return true
		}
	}
	return false
}

func isMeaningful(s string) bool {
	for _, r := range s {
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			return true
		}
	}
	return false
}

func firstValid(tokens []string) string {
	for _, t := range tokens {
		if isMeaningful(t) {
			return t
		}
	}
	return ""
}

func lastValid(tokens []string) string {
	for i := len(tokens) - 1; i >= 0; i-- {
		if isMeaningful(tokens[i]) {
			return tokens[i]
		}
	}
	return ""
}

func ParseWithFormat(raw, format string) []model.Credential {
	switch format {
	case "csv":
		return parseCSV(raw)
	case "sql":
		return parseSQL(raw)
	default:
		return Parse(raw)
	}
}
