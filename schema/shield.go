package schema

import (
	"errors"
	"fmt"
	"strings"

	"github.com/grokify/mogo/text/markdown"
)

const (
	SchemaNameAPIDocs     = "API Docs"
	ShieldNameLicense     = "License"
	ShieldNameBuildStatus = "Build Status"
)

type Shields []Shield

type Shield struct {
	ImageName string
	ImageURL  string
	LinkURL   string
}

func (s *Shield) Markdown() (string, error) {
	imageURL := strings.TrimSpace(s.ImageURL)
	linkURL := strings.TrimSpace(s.LinkURL)
	imgMD, err := ImageMarkdown(imageURL, s.ImageName)
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
