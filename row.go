package tables

import (
	"html/template"
)

type Rows []*Row

func (c Rows) render() string {
	var r string
	for _, v := range c {
		r += v.render()
	}
	return r
}

func (c Rows) Render() template.HTML {
	return template.HTML(c.render())
}

type Row struct {
	Attributes Attributes `json:"attributes,omitempty"`
	Cells      Cells      `json:"cells,omitempty"`
}

func (c *Row) defaultHTMLString() string {
	return `<` + TagRow + GenAttr(c.Attributes) + `>` + c.Cells.render() + `</` + TagRow + `>`
}

func (c *Row) render() string {
	return c.defaultHTMLString()
}

func (c *Row) Render() template.HTML {
	return template.HTML(c.render())
}
