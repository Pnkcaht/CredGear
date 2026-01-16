package integration

import (
	"testing"

	"cred-sanitizer/internal/deduplicator"
	"cred-sanitizer/internal/normalizer"
	"cred-sanitizer/internal/parser"
)

func TestRealWorldDump(t *testing.T) {
	raw := `
guestag.santander.com.br/
guestag.santander.com.br/::
guestag.santander.com.br/:9999772777:?
https://guestag.santander.com.br/::
guestag.santander.com.br/:13667079664:3108
guestag.santander.com.br/:+5515996875414:Tereza123
guestag.santander.com.br/wlc-captcha
emprestimos.gruporecovery.com/entrar:guicerini:Alemao123
emprestimos.gruporecovery.com/:sofialima031973@gmail.com:106465
https://emprestimos.gruporecovery.com/entrar
url: https://emprestimos.gruporecovery.com/entrar
`

	parsed := parser.Parse(raw)
	normalized := normalizer.Normalize(parsed)
	final := deduplicator.Deduplicate(normalized)

	if len(final) != 4 {
		t.Fatalf("esperado 4 credenciais válidas, veio %d", len(final))
	}

	for _, c := range final {
		if c.URL == "" {
			t.Fatal("URL vazia não permitida")
		}
		if c.Login == "" {
			t.Fatal("login vazio não permitido")
		}
		if c.Password == "" {
			t.Fatal("senha vazia não permitida")
		}

		// URL nunca pode ter path
		if containsSlashAfterDomain(c.URL) {
			t.Fatalf("URL com path indevido: %s", c.URL)
		}

		// senha lixo
		if c.Password == "?" {
			t.Fatal("senha inválida '?' passou")
		}
	}
}

func containsSlashAfterDomain(u string) bool {
	for i := len("https://"); i < len(u); i++ {
		if u[i] == '/' {
			return true
		}
	}
	return false
}
