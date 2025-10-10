package tables

import "html/template"

func New() *Table {
	return &Table{
		Attributes: Attributes{},
		Head:       &Head{},
		Body:       &Body{},
		Foot:       &Foot{},
	}
}

type Table struct {
	Attributes Attributes `json:"attributes,omitempty" xml:"attributes,omitempty"`
	Caption    *Caption   `json:"caption,omitempty" xml:"caption,omitempty"`
	Cols       Cols       `json:"cols,omitempty" xml:"cols,omitempty"`
	Head       *Head      `json:"head,omitempty" xml:"head,omitempty"`
	Body       *Body      `json:"body,omitempty" xml:"body,omitempty"`
	Foot       *Foot      `json:"foot,omitempty" xml:"foot,omitempty"`
}

func (c *Table) defaultHTMLString() string {
	var r string
	if c.Caption != nil {
		r += c.Caption.render()
	}
	if len(c.Cols) > 0 {
		r += c.Cols.render()
	}
	if c.Head != nil {
		r += c.Head.render()
	}
	if c.Body != nil {
		r += c.Body.render()
	}
	if c.Foot != nil {
		r += c.Foot.render()
	}
	return `<` + TagTable + GenAttr(c.Attributes) + `>` + r + `</` + TagTable + `>`
}

func (c *Table) render() string {
	return c.defaultHTMLString()
}

func (c *Table) Render() template.HTML {
	return template.HTML(c.render())
}

func (c *Table) SetAttr(k, v string) *Table {
	c.Attributes.Set(k, v)
	return c
}

func (c *Table) initCaption() {
	if c.Caption == nil {
		c.Caption = &Caption{}
	}
}

func (c *Table) SetCaptionAttr(k, v string) *Table {
	c.initCaption()
	c.Caption.SetAttr(k, v)
	return c
}

func (c *Table) SetCaptionStyle(style string) *Table {
	c.initCaption()
	c.Caption.Style = style
	return c
}

func (c *Table) SetCaptionTemplate(tmpl string) *Table {
	c.initCaption()
	c.Caption.Template = tmpl
	return c
}

func (c *Table) SetCaptionContent(content interface{}) *Table {
	c.initCaption()
	c.Caption.Content = content
	return c
}

func (c *Table) SetHeadAttr(k, v string) *Table {
	c.Head.SetAttr(k, v)
	return c
}

func (c *Table) AddHeadRow(rows ...*Row) *Table {
	c.Head.AddRow(rows...)
	return c
}

func (c *Table) SetBodyAttr(k, v string) *Table {
	c.Body.SetAttr(k, v)
	return c
}

func (c *Table) AddBodyRow(rows ...*Row) *Table {
	c.Body.AddRow(rows...)
	return c
}

func (c *Table) SetFootAttr(k, v string) *Table {
	c.Foot.SetAttr(k, v)
	return c
}

func (c *Table) AddFootRow(rows ...*Row) *Table {
	c.Foot.AddRow(rows...)
	return c
}

func (c *Table) AddCol(cols ...*Col) *Table {
	c.Cols.Add(cols...)
	return c
}

func (c *Table) ToMaps() []map[string]interface{} {
	r := make([]map[string]interface{}, len(c.Body.Rows))
	for i, row := range c.Body.Rows {
		m := make(map[string]interface{}, len(c.Cols))
		for j, col := range c.Cols {
			if len(col.Key) == 0 {
				continue
			}
			if len(row.Cells) > j {
				m[col.Key] = row.Cells[j].Content
			} else {
				break
			}
		}
		r[i] = m
	}
	return r
}
