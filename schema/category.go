package schema

import (
	"errors"
	"strings"
)

type Categories []Category

type Category struct {
	Path             string
	Name             string
	ParentNames      []string
	SubCategoryNames []string
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
	c.Path = strings.Join(catPath, "\t")
	return nil
}
