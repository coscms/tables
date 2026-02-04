package tables

import (
	"fmt"
	"html"
	"html/template"

	"github.com/coscms/forms/common"
	"github.com/coscms/forms/widgets"
)

// BaseWidget creates a Widget based on theme and inpuType parameters, both defined in the common package.
func BaseWidget(theme, inputType, tmplName string) *widgets.Widget {
	cachedKey := theme + ", " + inputType + ", " + tmplName
	tmpl, err := common.GetOrSetCachedTemplate(cachedKey, func() (*template.Template, error) {
		var (
			fpath = common.TmplDir(theme) + "/" + theme + "/"
			urls  = []string{common.LookupPath(fpath + "generic.html")}
			tpath = widgetTmpl(inputType, tmplName)
		)
		urls = append(urls, common.LookupPath(fpath+tpath+".html"))
		return common.ParseFiles(urls...)
	})
	if err != nil {
		panic(err)
	}
	tmpl.Funcs(common.TplFuncs())
	return widgets.New(tmpl)
}

func widgetTmpl(inputType, tmpl string) (tpath string) {
	return inputType + `/` + tmpl
}

func GenAttr(a Attributes) string {
	var attrs string
	if len(a) > 0 {
		attrs = a.String()
		if len(attrs) > 0 {
			attrs = ` ` + attrs
		}
	}
	return attrs
}

func GetContentString(c interface{}) string {
	switch vv := c.(type) {
	case template.HTML:
		return string(vv)
	case Renderer:
		return string(vv.Render())
	default:
		return html.EscapeString(fmt.Sprint(vv))
	}
}

func MakeIconClass(typ string, icon string) string {
	if len(icon) > 0 {
		return typ + ` ` + typ + `-` + icon
	}
	return typ
}
