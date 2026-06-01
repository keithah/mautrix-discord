package main

import (
	"testing"
	"time"

	"github.com/bwmarrin/discordgo"
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

func TestIsLinkPreviewOnlyDiscordUpdate(t *testing.T) {
	now := time.Now()
	tests := []struct {
		name string
		msg  *discordgo.Message
		want bool
	}{
		{
			name: "link preview update",
			msg: &discordgo.Message{
				Content: "https://example.com",
				Embeds: []*discordgo.MessageEmbed{{
					Type: discordgo.EmbedTypeLink,
					URL:  "https://example.com",
				}},
			},
			want: true,
		},
		{
			name: "actual edit",
			msg: &discordgo.Message{
				EditedTimestamp: &now,
				Content:         "https://example.com",
				Embeds: []*discordgo.MessageEmbed{{
					Type: discordgo.EmbedTypeLink,
					URL:  "https://example.com",
				}},
			},
			want: false,
		},
		{
			name: "rich embed update",
			msg: &discordgo.Message{
				Content: "bot embed",
				Embeds: []*discordgo.MessageEmbed{{
					Type:        discordgo.EmbedTypeRich,
					Description: "important embed",
				}},
			},
			want: false,
		},
		{
			name: "gif update",
			msg: &discordgo.Message{
				Content: "https://tenor.com/view/example",
				Embeds: []*discordgo.MessageEmbed{{
					Type: discordgo.EmbedTypeGifv,
					URL:  "https://tenor.com/view/example",
					Video: &discordgo.MessageEmbedVideo{
						URL: "https://media.tenor.com/example.mp4",
					},
				}},
			},
			want: false,
		},
		{
			name: "attachment update",
			msg: &discordgo.Message{
				Content:     "https://example.com",
				Attachments: []*discordgo.MessageAttachment{{ID: "attachment"}},
				Embeds: []*discordgo.MessageEmbed{{
					Type: discordgo.EmbedTypeLink,
					URL:  "https://example.com",
				}},
			},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isLinkPreviewOnlyDiscordUpdate(tt.msg); got != tt.want {
				t.Fatalf("isLinkPreviewOnlyDiscordUpdate() = %v, want %v", got, tt.want)
			}
		})
	}
}
