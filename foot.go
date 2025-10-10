package tables

import "html/template"

type Foot struct {
	Attributes Attributes `json:"attributes,omitempty" xml:"attributes,omitempty"`
	Rows       Rows       `json:"rows,omitempty" xml:"rows,omitempty"`
}

func (c *Foot) defaultHTMLString() string {
	return `<` + TagFoot + GenAttr(c.Attributes) + `>` + c.Rows.render() + `</` + TagFoot + `>`
}

func (c *Foot) AddRow(rows ...*Row) *Foot {
	c.Rows.Add(rows...)
	return c
}

func (c *Foot) SetAttr(k, v string) *Foot {
	c.Attributes.Set(k, v)
	return c
}

func (c *Foot) render() string {
	return c.defaultHTMLString()
}

func (c *Foot) Render() template.HTML {
	return template.HTML(c.render())
}
