package tables

import "html/template"

type Head struct {
	Attributes Attributes `json:"attributes,omitempty" xml:"attributes,omitempty"`
	Rows       Rows       `json:"rows,omitempty" xml:"rows,omitempty"`
}

func (c *Head) defaultHTMLString() string {
	return `<` + TagHead + GenAttr(c.Attributes) + `>` + c.Rows.render() + `</` + TagHead + `>`
}

func (c *Head) AddRow(rows ...*Row) *Head {
	c.Rows.Add(rows...)
	return c
}

func (c *Head) AddHeadRow(rows ...*Row) *Head {
	for _, row := range rows {
		for _, cell := range row.Cells {
			cell.IsHead = true
		}
	}
	c.AddRow(rows...)
	return c
}

func (c *Head) SetAttr(k, v string) *Head {
	c.Attributes.Set(k, v)
	return c
}

func (c *Head) render() string {
	return c.defaultHTMLString()
}

func (c *Head) Render() template.HTML {
	return template.HTML(c.render())
}
