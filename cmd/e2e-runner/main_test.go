package main

import (
	"strings"
	"testing"
)

func TestKeyBytes(t *testing.T) {
	tests := map[string]string{
		"enter":  "\r",
		"esc":    "\x1b",
		"ctrl+c": "\x03",
		"space":  " ",
		"x":      "x",
	}

	for key, want := range tests {
		if got := keyBytes(key); got != want {
			t.Fatalf("keyBytes(%q) = %q, want %q", key, got, want)
		}
	}
}

func TestCleanTerminal(t *testing.T) {
	raw := []byte("\x1b]11;?\x1b\\\x1b[?1049h\x1b[1;1Hhello\r\n\x1b[31mworld\x1b[0m\x07")
	clean := cleanTerminal(raw)

	if !strings.Contains(clean, "hello\nworld") {
		t.Fatalf("cleaned output = %q", clean)
	}
	if strings.Contains(clean, "\x1b") {
		t.Fatalf("cleaned output still contains escape sequence: %q", clean)
	}
}
