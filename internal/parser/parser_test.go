package parser

import "testing"

func TestSantanderExample(t *testing.T) {
	raw := `
guestag.santander.com.br/
guestag.santander.com.br/::
guestag.santander.com.br/:9999772777:?
`

	res := Parse(raw)

	if len(res) != 1 {
		t.Fatalf("esperado 1 credencial, veio %d", len(res))
	}

	if res[0].Login != "guestag.santander.com.br" {
		t.Fatal("login errado")
	}

	if res[0].Password != "9999772777" {
		t.Fatal("senha errada")
	}
}
