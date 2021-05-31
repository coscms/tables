package tables

import "html/template"

func New() *Table {
	return &Table{
		Head: &Head{},
		Body: &Body{},
		Foot: &Foot{},
	}
}

type Table struct {
	Attributes Attributes `json:"attributes,omitempty" xml:"attributes,omitempty"`
	Caption    *Caption   `json:"caption,omitempty" xml:"caption,omitempty"`
	Cols       Cols       `json:"cols,omitempty" xml:"cols,omitempty"`
	Head       *Head      `json:"head,omitempty" xml:"head,omitempty"`
	Body       *Body      `json:"body,omitempty" xml:"body,omitempty"`
	Foot       *Foot      `json:"foot,omitempty" xml:"foot,omitempty"`
}

func (c *Table) defaultHTMLString() string {
	var r string
	if c.Caption != nil {
		r += c.Caption.render()
	}
	if len(c.Cols) > 0 {
		r += c.Cols.render()
	}
	if c.Head != nil {
		r += c.Head.render()
	}
	if c.Body != nil {
		r += c.Body.render()
	}
	if c.Foot != nil {
		r += c.Foot.render()
	}
	return `<` + TagTable + GenAttr(c.Attributes) + `>` + r + `</` + TagTable + `>`
}

func (c *Table) render() string {
	return c.defaultHTMLString()
}

func (c *Table) Render() template.HTML {
	return template.HTML(c.render())
}
