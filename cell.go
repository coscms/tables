package tables

import (
	"fmt"
	"html"
	"html/template"

	"github.com/coscms/forms/widgets"
)

type Cells []*Cell

func (c Cells) render() string {
	var r string
	for _, v := range c {
		r += v.render()
	}
	return r
}

func (c Cells) Render() template.HTML {
	return template.HTML(c.render())
}

type Cell struct {
	IsHead     bool        `json:"isHead" xml:"isHead"`
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
	if c.IsHead {
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
		if c.IsHead {
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
