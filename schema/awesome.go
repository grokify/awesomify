package schema

import (
	"errors"
	"strings"

	"github.com/grokify/mogo/type/stringsutil"
)

type Awesome struct {
	Categories Categories
	Entries    Entries
}

func (a *Awesome) Validate() (catNoEntry []string, entryCatMissing []string, entryNoCatCount uint, err error) {
	catsMap := a.Categories.Map()
	for _, entry := range a.Entries {
		catPath := entry.Category.FullQualifedPath()
		if len(catPath) == 0 {
			entryNoCatCount++
		} else {
			if _, ok := catsMap[catPath]; !ok {
				entryCatMissing = append(entryCatMissing, catPath)
			}
		}
	}
	catsMapEntries := a.Entries.CategoriesMap()
	for catPath := range catsMap {
		if _, ok := catsMapEntries[catPath]; !ok {
			catNoEntry = append(catNoEntry, catPath)
		}
	}
	entryCatMissing = stringsutil.Dedupe(entryCatMissing)
	if len(catNoEntry) > 0 || len(entryCatMissing) > 0 || entryNoCatCount > 0 {
		err = errors.New("category validation failure")
	}
	return
}

func (a *Awesome) Markdown() (string, error) {
	md := ""
	for _, cat := range a.Categories {
		mdCat, err := a.MarkdownForCategory(cat.FullQualifedPath())
		if err != nil {
			return "", err
		}
		md += mdCat
	}
	return md, nil
}

func (a *Awesome) MarkdownTOC() string {
	return ""
}

func (a *Awesome) MarkdownForCategory(catPath string) (string, error) {
	catPathParts := strings.Split(catPath, "\t")
	header := ""
	for i, l := 0, len(catPathParts)+1; i < l; i++ {
		header += "#"
	}
	header += " " + catPathParts[len(catPathParts)-1]
	lines := []string{header}
	entriesByCat := a.EntriesByCategory(Categories{{Path: catPathParts}})
	for _, entry := range entriesByCat {
		md, err := entry.Markdown()
		if err != nil {
			return "", err
		}
		lines = append(lines, "- "+md)
	}
	return strings.Join(lines, "\n") + "\n", nil
	//	header := len(catPathParts) + 1
}

func (a *Awesome) EntriesByCategory(catFilter Categories) Entries {
	if len(catFilter) == 0 {
		return a.Entries
	}
	matches := Entries{}
	m := map[string]int{}
	for _, cat := range catFilter {
		m[cat.FullQualifedPath()] = 1
	}
	for _, entry := range a.Entries {
		if _, ok := m[entry.Category.FullQualifedPath()]; ok {
			matches = append(matches, entry)
		}
	}
	return matches
}
