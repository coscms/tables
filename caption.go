package tables

import (
	"fmt"
	"html"
	"html/template"

	"github.com/coscms/forms/widgets"
)

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
