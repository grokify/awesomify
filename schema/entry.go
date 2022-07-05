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
	Shields      []Shield
}

func (e *Entry) Markdown() (string, error) {
	nameParts := []string{}
	if len(strings.TrimSpace(e.Name)) == 0 {
		return "", ErrEntryNameMissing
	}
	nameParts = append(nameParts, markdown.Linkify(e.URL, e.Name))
	for _, shield := range e.Shields {
		shieldMD, err := shield.Markdown()
		if err != nil {
			return "", err
		}
		nameParts = append(nameParts, shieldMD)
	}
	nameString := strings.Join(nameParts, " ")
	e.Description = strings.TrimSpace(e.Description)
	if len(e.Description) == 0 {
		return nameString, nil
	}
	return nameString + " — " + e.Description, nil
}
