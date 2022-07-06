package schema

import (
	"errors"
	"strings"

	"github.com/grokify/mogo/text/markdown"
)

var ErrEntryNameMissing = errors.New("entry name is required, but not present, for Markdown conversion")

type Entries []Entry

type Entry struct {
	URL          string
	Name         string
	Description  string
	AccessDate   string   // RFC-3339 `full-date`
	CategoryPath []string // tab delimited
	Badges       []Badge
	GitHubStars  uint
}

func (e *Entry) Category() Category {
	cat := Category{
		Path: e.CategoryPath,
	}
	if len(cat.Path) > 0 {
		cat.Name = cat.Path[len(cat.Path)-1]
		cat.ParentNames = cat.Path[:len(cat.Path)-1]
	}
}

func (e *Entry) Markdown() (string, error) {
	nameParts := []string{}
	if len(strings.TrimSpace(e.Name)) == 0 {
		return "", ErrEntryNameMissing
	}
	nameParts = append(nameParts, markdown.Linkify(e.URL, e.Name))
	for _, badge := range e.Badges {
		badgeMD, err := badge.Markdown()
		if err != nil {
			return "", err
		}
		nameParts = append(nameParts, badgeMD)
	}
	nameString := strings.Join(nameParts, " ")
	e.Description = strings.TrimSpace(e.Description)
	if len(e.Description) == 0 {
		return nameString, nil
	}
	return nameString + " — " + e.Description, nil
}
