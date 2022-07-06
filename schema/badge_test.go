package schema

import (
	"testing"
)

var badgeMarkdownTests = []struct {
	imageName string
	imageURL  string
	linkURL   string
	markdown  string
}{
	{"", `https://github.com/grokify/spectrum/workflows/go%20build/badge.svg`, "https://github.com/grokify/spectrum/actions",
		`[![](https://github.com/grokify/spectrum/workflows/go%20build/badge.svg)](https://github.com/grokify/spectrum/actions)`},
	//`[![][https://github.com/grokify/spectrum/workflows/go%20build/badge.svg]]["https://github.com/grokify/spectrum/actions"]`},
}

// TestBadgeMarkdown ensures correct markdown is being produced
func TestBadgeMarkdown(t *testing.T) {
	for _, tt := range badgeMarkdownTests {
		badge := Badge{
			ImageName: tt.imageName,
			ImageURL:  tt.imageURL,
			LinkURL:   tt.linkURL}
		md, err := badge.Markdown()
		if err != nil {
			t.Errorf("badge.Markdown() Error [%s]", err.Error())
		}
		if md != tt.markdown {
			t.Errorf("badge.Markdown(...) Mismatch: want [%s], got [%s]", tt.markdown, md)
		}
	}
}
