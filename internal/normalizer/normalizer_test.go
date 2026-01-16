package normalizer

import "testing"

func TestNormalizeURL(t *testing.T) {
	cases := map[string]string{
		"guestag.santander.com.br/":                  "https://guestag.santander.com.br",
		"https://guestag.santander.com.br/":          "https://guestag.santander.com.br",
		"https://guestag.santander.com.br/entrar":    "https://guestag.santander.com.br",
		"http://emprestimos.gruporecovery.com/login": "http://emprestimos.gruporecovery.com",
		"": "",
	}

	for in, expected := range cases {
		out := normalizeURL(in)
		if out != expected {
			t.Fatalf("normalizeURL(%q) = %q, esperado %q", in, out, expected)
		}
	}
}
