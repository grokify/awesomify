package schema

import (
	"errors"
	"strings"
)

const CategoryDelimiter = "\t"

type Categories []Category

func (cats Categories) Map() map[string]Category {
	catsMap := map[string]Category{}
	for _, cat := range cats {
		catsMap[cat.FullQualifedPath()] = cat
	}
	return catsMap
}

func (cats Categories) Strings() []string {
	strs := []string{}
	for _, cat := range cats {
		strs = append(strs, cat.FullQualifedPath())
	}
	return strs
}

type Category struct {
	Path             []string
	Name             string
	Description      string
	ParentNames      []string
	SubCategoryNames []string
}

func (c *Category) condensePath() {
	condensed := []string{}
	for _, part := range c.Path {
		part = strings.TrimSpace(part)
		if len(part) > 0 {
			condensed = append(condensed, part)
		}
	}
	c.Path = condensed
}

func (c *Category) FullQualifedPath() string {
	c.condensePath()
	return strings.Join(c.Path, CategoryDelimiter)
}

func (c *Category) Inflate(force bool) error {
	if len(c.Path) > 0 && !force {
		return nil
	}
	catPath := []string{}
	catPath = append(catPath, c.ParentNames...)
	if len(c.Name) == 0 {
		return errors.New("category name is not set")
	}
	catPath = append(catPath, c.Name)
	c.Path = catPath
	//c.Path = strings.Join(catPath, "\t")
	return nil
}
