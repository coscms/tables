package tables

import "html/template"

type Cols []*Col

func (c Cols) render() string {
	var r string
	for _, v := range c {
		r += v.render()
	}
	return r
}

func (c Cols) Render() template.HTML {
	return template.HTML(c.render())
}

type Col struct {
	Attributes Attributes `json:"attributes,omitempty" xml:"attributes,omitempty"`
}

func (c *Col) defaultHTMLString() string {
	return `<` + TagCol + GenAttr(c.Attributes) + ` />`
}

func (c *Col) render() string {
	return c.defaultHTMLString()
}

func (c *Col) Render() template.HTML {
	return template.HTML(c.render())
}
