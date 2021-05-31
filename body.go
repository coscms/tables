package tables

import "html/template"

type Body struct {
	Attributes Attributes `json:"attributes,omitempty"`
	Rows       Rows       `json:"rows,omitempty"`
}

func (c *Body) defaultHTMLString() string {
	return `<` + TagBody + GenAttr(c.Attributes) + `>` + c.Rows.render() + `</` + TagBody + `>`
}

func (c *Body) render() string {
	return c.defaultHTMLString()
}

func (c *Body) Render() template.HTML {
	return template.HTML(c.render())
}
