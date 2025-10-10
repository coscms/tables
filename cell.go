package tables

import (
	"fmt"
	"html"
	"html/template"

	"github.com/coscms/forms/widgets"
)

func NewCell(content interface{}, options ...func(c *Cell)) *Cell {
	c := &Cell{
		Content: content,
	}
	for _, opt := range options {
		opt(c)
	}
	return c
}

type Cells []*Cell

func (c Cells) render() string {
	var r string
	for _, v := range c {
		r += v.render()
	}
	return r
}

func (c *Cells) Add(cells ...*Cell) *Cells {
	*c = append(*c, cells...)
	return c
}

func (c Cells) Render() template.HTML {
	return template.HTML(c.render())
}

func CellIsHead(isHead bool) func(c *Cell) {
	return func(c *Cell) {
		c.IsHead = &isHead
	}
}

func CellStyle(style string) func(c *Cell) {
	return func(c *Cell) {
		c.Style = style
	}
}

func CellTemplate(tmpl string) func(c *Cell) {
	return func(c *Cell) {
		c.Template = tmpl
	}
}

func CellAttributes(attributes Attributes) func(c *Cell) {
	return func(c *Cell) {
		c.Attributes = attributes
	}
}

func CellContent(content interface{}) func(c *Cell) {
	return func(c *Cell) {
		c.Content = content
	}
}

type Cell struct {
	IsHead     *bool       `json:"isHead" xml:"isHead"`
	Style      string      `json:"style" xml:"style"`
	Template   string      `json:"template" xml:"template"`
	Attributes Attributes  `json:"attributes,omitempty" xml:"attributes,omitempty"`
	Content    interface{} `json:"content" xml:"content"`
	widget     widgets.WidgetInterface
}

func (c *Cell) String() string {
	return fmt.Sprint(c.Content)
}

func (c *Cell) defaultHTMLString() string {
	tag := TagCell
	if c.IsHead != nil && *c.IsHead {
		tag = TagHeadCell
	}
	return `<` + tag + GenAttr(c.Attributes) + `>` + html.EscapeString(fmt.Sprint(c.Content)) + `</` + tag + `>`
}

func (c *Cell) render() string {
	if len(c.Template) == 0 {
		return c.defaultHTMLString()
	}
	if c.widget == nil {
		tag := TagCell
		if c.IsHead != nil && *c.IsHead {
			tag = TagHeadCell
		}
		c.widget = widgets.BaseWidget(c.Style, tag, c.Template)
	}
	data := map[string]interface{}{
		`content`:    c.Content,
		`attributes`: c.Attributes,
	}
	return c.widget.Render(data)
}

func (c *Cell) Render() template.HTML {
	return template.HTML(c.render())
}

func (c *Cell) SetAttr(k, v string) *Cell {
	c.Attributes.Set(k, v)
	return c
}

func (c *Cell) SetStyle(style string) *Cell {
	c.Style = style
	return c
}

func (c *Cell) SetTemplate(tmpl string) *Cell {
	c.Template = tmpl
	return c
}

func (c *Cell) SetContent(content interface{}) *Cell {
	c.Content = content
	return c
}

func (c *Cell) SetIsHead(isHead bool) *Cell {
	c.IsHead = &isHead
	return c
}
