package tables

import (
	"html/template"
	"strings"
)

type Cols []*Col

func (c Cols) render() string {
	var r strings.Builder
	for _, v := range c {
		r.WriteString(v.render())
	}
	return r.String()
}

func (c *Cols) Add(cols ...*Col) *Cols {
	*c = append(*c, cols...)
	return c
}

func (c Cols) Render() template.HTML {
	return template.HTML(c.render())
}

func NewCol(attributes Attributes, key ...string) *Col {
	c := &Col{
		Attributes: attributes,
	}
	if len(key) > 0 {
		c.Key = key[0]
	}
	return c
}

type Col struct {
	Attributes Attributes `json:"attributes,omitempty" xml:"attributes,omitempty"`
	Key        string     `json:"key,omitempty" xml:"key,omitempty"`
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

func (c *Col) SetAttr(k, v string) *Col {
	c.Attributes.Set(k, v)
	return c
}

func (c *Col) SetKey(key string) *Col {
	c.Key = key
	return c
}
