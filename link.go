package tables

import (
	"html/template"
	"strings"

	"github.com/coscms/forms/widgets"
)

func NewLinks() *Links {
	return &Links{}
}

type Links []*Link

func (l Links) Render() template.HTML {
	var r strings.Builder
	for _, v := range l {
		r.WriteString(v.render())
		r.WriteString(` `)
	}
	return template.HTML(r.String())
}

func (l *Links) Add(link *Link) *Links {
	*l = append(*l, link)
	return l
}

func (l *Links) Reset() *Links {
	*l = (*l)[0:0]
	return l
}

type Link struct {
	Href       string      `json:"href" xml:"href"`
	Title      string      `json:"title" xml:"title"`
	Icon       string      `json:"icon" xml:"icon"`
	Confirm    string      `json:"confirm" xml:"confirm"`
	Class      string      `json:"class" xml:"class"`
	Toggle     string      `json:"toggle" xml:"toggle"`
	Content    interface{} `json:"content" xml:"content"`
	Attributes Attributes  `json:"attributes,omitempty" xml:"attributes,omitempty"`
	Template   string      `json:"template" xml:"template"`
	Theme      string      `json:"theme" xml:"theme"`
	widget     widgets.WidgetInterface
}

func (l *Link) SetContent(content interface{}) *Link {
	l.Content = content
	return l
}

func (l *Link) SetHref(href string) *Link {
	l.Href = href
	return l
}

func (l *Link) SetTitle(title string) *Link {
	l.Title = title
	return l
}

func (l *Link) SetIcon(icon string) *Link {
	l.Icon = icon
	return l
}

func (l *Link) SetConfirm(confirm string) *Link {
	l.Confirm = confirm
	return l
}

func (l *Link) SetClass(class string) *Link {
	l.Class = class
	return l
}

func (l *Link) AddClass(class string) *Link {
	if len(l.Class) > 0 {
		l.Class += ` `
	}
	l.Class += class
	return l
}

func (l *Link) SetToggle(toggle string) *Link {
	l.Toggle = toggle
	return l
}

func (c *Link) SetWidget(widget widgets.WidgetInterface) *Link {
	c.widget = widget
	return c
}

func (c *Link) SetAttributes(attributes Attributes) *Link {
	c.Attributes = attributes
	return c
}

func (c *Link) SetTemplate(tmpl string) *Link {
	c.Template = tmpl
	return c
}

func (c *Link) SetTheme(theme string) *Link {
	c.Theme = theme
	return c
}

func (c *Link) defaultHTMLString() string {
	tag := TagLink
	v := GetContentString(c.Content)
	class := c.Class
	if len(class) > 0 {
		class = ` class="` + class + `"`
	}
	href := c.Href
	if len(href) > 0 {
		href = ` href="` + href + `"`
	}
	title := c.Title
	if len(title) > 0 {
		title = ` title="` + title + `"`
	}
	toggle := c.Toggle
	if len(toggle) > 0 {
		toggle = ` data-toggle="` + toggle + `"`
	}
	confirm := c.Confirm
	if len(confirm) > 0 {
		confirm = ` onclick="return confirm('` + confirm + `')"`
	}
	icon := c.Icon
	if len(icon) > 0 {
		icon = `<i class="` + icon + `"></i>`
		if len(v) > 0 {
			v = icon + ` ` + v
		} else {
			v = icon
		}
	}
	return `<` + tag + class + href + title + toggle + confirm + GenAttr(c.Attributes) + `>` + v + `</` + tag + `>`
}

func (c *Link) render() string {
	if len(c.Template) == 0 {
		return c.defaultHTMLString()
	}
	if c.widget == nil {
		tag := TagLink
		c.widget = widgets.BaseWidget(c.Theme, tag, c.Template)
	}
	data := map[string]interface{}{
		`content`:    c.Content,
		`attributes`: c.Attributes,
	}
	return c.widget.Render(data)
}

func (c *Link) Render() template.HTML {
	return template.HTML(c.render())
}

func (c *Link) SetAttr(k, v string) *Link {
	c.Attributes.Set(k, v)
	return c
}

func NewLink(options ...func(c *Link)) *Link {
	l := &Link{}
	for _, option := range options {
		option(l)
	}
	return l
}

func NewLabelLink(colorName string, options ...func(c *Link)) *Link {
	return NewLink(options...).AddClass(`label label-` + colorName)
}

func NewButtonLink(colorName string, options ...func(c *Link)) *Link {
	return NewLink(options...).AddClass(`btn btn-` + colorName)
}

func LinkContent(content interface{}) func(c *Link) {
	return func(c *Link) {
		c.Content = content
	}
}

func LinkHref(href string) func(c *Link) {
	return func(c *Link) {
		c.Href = href
	}
}

func LinkTitle(title string) func(c *Link) {
	return func(c *Link) {
		c.Title = title
	}
}

func LinkIcon(icon string) func(c *Link) {
	return func(c *Link) {
		c.Icon = icon
	}
}

func LinkConfirm(confirm string) func(c *Link) {
	return func(c *Link) {
		c.Confirm = confirm
	}
}

func LinkClass(class string) func(c *Link) {
	return func(c *Link) {
		c.Class = class
	}
}

func LinkToggle(toggle string) func(c *Link) {
	return func(c *Link) {
		c.Toggle = toggle
	}
}

func LinkAttributes(attributes Attributes) func(c *Link) {
	return func(c *Link) {
		c.Attributes = attributes
	}
}

func LinkTemplate(tmpl string) func(c *Link) {
	return func(c *Link) {
		c.Template = tmpl
	}
}

func LinkTheme(theme string) func(c *Link) {
	return func(c *Link) {
		c.Theme = theme
	}
}

func LinkWidget(widget widgets.WidgetInterface) func(c *Link) {
	return func(c *Link) {
		c.widget = widget
	}
}

func LinkAttr(k, v string) func(c *Link) {
	return func(c *Link) {
		c.Attributes.Set(k, v)
	}
}

func LinkAddClass(class string) func(c *Link) {
	return func(c *Link) {
		c.AddClass(class)
	}
}
