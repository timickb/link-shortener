package shortener

import (
	"strings"
	"testing"
)

func TestGenerateShortening(t *testing.T) {
	const (
		lower    = "abcdefghijklmnopqrstuvwxyz"
		upper    = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
		numbers  = "0123456789"
		special  = "_"
		alphabet = lower + upper + numbers + special
	)

	url := "https://github.com"

	short, err := generateShortening(url)
	if err != nil {
		t.Fatal(err)
	}

	if len(short) != 10 {
		t.Fatal("expected len 10")
	}

	for _, c := range short {
		if !strings.Contains(alphabet, string(c)) {
			t.Fatalf("invalid symbol %s", string(c))
		}
	}
}
