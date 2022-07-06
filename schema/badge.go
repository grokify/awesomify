package schema

import (
	"errors"
	"fmt"
	"strings"

	"github.com/grokify/mogo/text/markdown"
)

const (
	BadgeNameAPIDocs     = "API Docs"
	BadgeNameGoReference = "Go Reference"
	BadgeNameLicense     = "License"
	BadgeNameBuildStatus = "Build Status"
)

type Badges []Badge

type Badge struct {
	ImageName string
	ImageURL  string
	LinkURL   string
}

func (b *Badge) Markdown() (string, error) {
	imageURL := strings.TrimSpace(b.ImageURL)
	linkURL := strings.TrimSpace(b.LinkURL)
	imgMD, err := ImageMarkdown(imageURL, b.ImageName)
	if err != nil {
		return "", err
	}
	if len(linkURL) == 0 {
		return imgMD, nil
	}
	return markdown.Linkify(linkURL, imgMD), nil
}

var ErrImageURLRequired = errors.New("imageURL required but not present")

func ImageMarkdown(imageURL, imageText string) (string, error) {
	imageURL = strings.TrimSpace(imageURL)
	if len(imageURL) == 0 {
		return "", ErrImageURLRequired
	}
	return fmt.Sprintf("![%s](%s)", imageText, imageURL), nil
}
