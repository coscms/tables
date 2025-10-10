package tables

import "html/template"

type Body struct {
	Attributes Attributes `json:"attributes,omitempty" xml:"attributes,omitempty"`
	Rows       Rows       `json:"rows,omitempty" xml:"rows,omitempty"`
}

func (c *Body) defaultHTMLString() string {
	return `<` + TagBody + GenAttr(c.Attributes) + `>` + c.Rows.render() + `</` + TagBody + `>`
}

func (c *Body) render() string {
	return c.defaultHTMLString()
}

func (c *Body) AddRow(rows ...*Row) *Body {
	c.Rows.Add(rows...)
	return c
}

func (c *Body) SetAttr(k, v string) *Body {
	c.Attributes.Set(k, v)
	return c
}

func (c *Body) Render() template.HTML {
	return template.HTML(c.render())
}
