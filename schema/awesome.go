package schema

type Awesome struct {
	Categories Categories
	Entries    Entries
}

func (a *Awesome) Markdown() string {
	return ""
}
