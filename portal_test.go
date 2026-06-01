package main

import (
	"testing"

	"go.mau.fi/util/variationselector"
)

func TestIsValidDiscordUnicodeReactionKey(t *testing.T) {
	tests := []struct {
		name string
		key  string
		want bool
	}{
		{name: "simple emoji", key: "👍", want: true},
		{name: "skin tone", key: "👍🏽", want: true},
		{name: "fully qualified text presentation emoji", key: variationselector.FullyQualify("✂"), want: true},
		{name: "flag", key: "🇫🇮", want: true},
		{name: "zwj sequence", key: "👨‍👩‍👧‍👦", want: true},
		{name: "keycap", key: "2️⃣", want: true},
		{name: "heisenbridge edit marker", key: "✂ 2 lines", want: false},
		{name: "plain text", key: "lol", want: false},
		{name: "number text", key: "123", want: false},
		{name: "leading space", key: " 👍", want: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isValidDiscordUnicodeReactionKey(tt.key); got != tt.want {
				t.Fatalf("isValidDiscordUnicodeReactionKey(%q) = %v, want %v", tt.key, got, tt.want)
			}
		})
	}
}
