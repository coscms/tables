package tables

import "html/template"

type Head struct {
	Attributes Attributes `json:"attributes,omitempty"`
	Rows       Rows       `json:"rows,omitempty"`
}

func (c *Head) defaultHTMLString() string {
	return `<` + TagHead + GenAttr(c.Attributes) + `>` + c.Rows.render() + `</` + TagHead + `>`
}

func (c *Head) render() string {
	return c.defaultHTMLString()
}

func (c *Head) Render() template.HTML {
	return template.HTML(c.render())
}
