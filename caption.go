package tables

import (
	"fmt"
	"html"
	"html/template"

	"github.com/coscms/forms/widgets"
)

func NewCaption(content interface{}, options ...func(c *Caption)) *Caption {
	c := &Caption{
		Content: content,
	}
	for _, opt := range options {
		opt(c)
	}
	return c
}

func CaptionStyle(style string) func(c *Caption) {
	return func(c *Caption) {
		c.Style = style
	}
}

func CaptionTemplate(tmpl string) func(c *Caption) {
	return func(c *Caption) {
		c.Template = tmpl
	}
}

func CaptionAttributes(attributes Attributes) func(c *Caption) {
	return func(c *Caption) {
		c.Attributes = attributes
	}
}

func CaptionContent(content interface{}) func(c *Caption) {
	return func(c *Caption) {
		c.Content = content
	}
}

type Caption struct {
	Attributes Attributes  `json:"attributes,omitempty" xml:"attributes,omitempty"`
	Style      string      `json:"style" xml:"style"`
	Template   string      `json:"template" xml:"template"`
	Content    interface{} `json:"content" xml:"content"`
	widget     widgets.WidgetInterface
}

func (c *Caption) String() string {
	return fmt.Sprint(c.Content)
}

func (c *Caption) defaultHTMLString() string {
	return `<` + TagCaption + GenAttr(c.Attributes) + `>` + html.EscapeString(fmt.Sprint(c.Content)) + `</` + TagCaption + `>`
}

func (c *Caption) render() string {
	if len(c.Template) == 0 {
		return c.defaultHTMLString()
	}
	if c.widget == nil {
		c.widget = widgets.BaseWidget(c.Style, TagCaption, c.Template)
	}
	return c.widget.Render(c.Content)
}

func (c *Caption) Render() template.HTML {
	return template.HTML(c.render())
}

func (c *Caption) SetAttr(k, v string) *Caption {
	c.Attributes.Set(k, v)
	return c
}

func (c *Caption) SetStyle(style string) *Caption {
	c.Style = style
	return c
}

func (c *Caption) SetTemplate(tmpl string) *Caption {
	c.Template = tmpl
	return c
}

func (c *Caption) SetContent(content interface{}) *Caption {
	c.Content = content
	return c
}
