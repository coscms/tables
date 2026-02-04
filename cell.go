package tables

import (
	"fmt"
	"html/template"
	"strings"

	"github.com/coscms/forms/widgets"
)

func NewCell(content interface{}, options ...func(c *Cell)) *Cell {
	c := &Cell{
		Content:    content,
		Attributes: Attributes{},
	}
	for _, opt := range options {
		opt(c)
	}
	return c
}

type Cells []*Cell

func (c Cells) render() string {
	var r strings.Builder
	for _, v := range c {
		r.WriteString(v.render())
	}
	return r.String()
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

func CellTheme(theme string) func(c *Cell) {
	return func(c *Cell) {
		c.Theme = theme
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

func CellAttr(k, v string) func(c *Cell) {
	return func(c *Cell) {
		c.Attributes.Set(k, v)
	}
}

func CellContent(content interface{}) func(c *Cell) {
	return func(c *Cell) {
		c.Content = content
	}
}

type Cell struct {
	IsHead     *bool       `json:"isHead" xml:"isHead"`
	Theme      string      `json:"theme" xml:"theme"`
	Template   string      `json:"template" xml:"template"`
	Attributes Attributes  `json:"attributes,omitempty" xml:"attributes,omitempty"`
	Content    interface{} `json:"content" xml:"content"`
	widget     widgets.WidgetInterface
}

func (c *Cell) String() string {
	return fmt.Sprint(c.Content)
}

func (c *Cell) ContentHTML() template.HTML {
	return GetContentHTML(c.Content)
}

func (c *Cell) defaultHTMLString() string {
	tag := TagCell
	if c.IsHead != nil && *c.IsHead {
		tag = TagHeadCell
	}
	v := string(c.ContentHTML())
	return `<` + tag + GenAttr(c.Attributes) + `>` + v + `</` + tag + `>`
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
		c.widget = widgets.BaseWidget(c.Theme, tag, c.Template)
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

func (c *Cell) SetTheme(theme string) *Cell {
	c.Theme = theme
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

func (c *Cell) SetWidget(widget widgets.WidgetInterface) *Cell {
	c.widget = widget
	return c
}

func (c *Cell) SetAttributes(attributes Attributes) *Cell {
	c.Attributes = attributes
	return c
}
