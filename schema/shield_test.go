package schema

import (
	"testing"
)

var shieldMarkdownTests = []struct {
	imageName string
	imageURL  string
	linkURL   string
	markdown  string
}{
	{"", `https://github.com/grokify/spectrum/workflows/go%20build/badge.svg`, "https://github.com/grokify/spectrum/actions",
		`[![](https://github.com/grokify/spectrum/workflows/go%20build/badge.svg)](https://github.com/grokify/spectrum/actions)`},
	//`[![][https://github.com/grokify/spectrum/workflows/go%20build/badge.svg]]["https://github.com/grokify/spectrum/actions"]`},
}

// TestShieldMarkdown ensures correct markdown is being produced
func TestShieldMarkdown(t *testing.T) {
	for _, tt := range shieldMarkdownTests {
		shield := Shield{
			ImageName: tt.imageName,
			ImageURL:  tt.imageURL,
			LinkURL:   tt.linkURL}
		md, err := shield.Markdown()
		if err != nil {
			t.Errorf("shield.Markdown() Error [%s]", err.Error())
		}
		if md != tt.markdown {
			t.Errorf("shield.Markdown(...) Mismatch: want [%s], got [%s]", tt.markdown, md)
		}
	}
}
