package tables

import "html/template"

type Foot struct {
	Attributes Attributes `json:"attributes,omitempty"`
	Rows       Rows       `json:"rows,omitempty"`
}

func (c *Foot) defaultHTMLString() string {
	return `<` + TagFoot + GenAttr(c.Attributes) + `>` + c.Rows.render() + `</` + TagFoot + `>`
}

func (c *Foot) render() string {
	return c.defaultHTMLString()
}

func (c *Foot) Render() template.HTML {
	return template.HTML(c.render())
}
