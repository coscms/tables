package tables

import (
	"html/template"
	"strings"
)

type Rows []*Row

func (c Rows) render() string {
	var r strings.Builder
	for _, v := range c {
		r.WriteString(v.render())
	}
	return r.String()
}

func (c *Rows) Add(rows ...*Row) *Rows {
	*c = append(*c, rows...)
	return c
}

func (c Rows) Render() template.HTML {
	return template.HTML(c.render())
}

func NewRow(cells ...*Cell) *Row {
	return &Row{
		Cells:      cells,
		Attributes: Attributes{},
	}
}

type Row struct {
	Attributes Attributes `json:"attributes,omitempty"`
	Cells      Cells      `json:"cells,omitempty"`
}

func (c *Row) defaultHTMLString() string {
	return `<` + TagRow + GenAttr(c.Attributes) + `>` + c.Cells.render() + `</` + TagRow + `>`
}

func (c *Row) AddCell(cells ...*Cell) *Row {
	c.Cells.Add(cells...)
	return c
}

func (c *Row) SetAttr(k, v string) *Row {
	c.Attributes.Set(k, v)
	return c
}

func (c *Row) render() string {
	return c.defaultHTMLString()
}

func (c *Row) Render() template.HTML {
	return template.HTML(c.render())
}
