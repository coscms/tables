package tables

import (
	"html/template"

	"github.com/coscms/forms/common"
	"github.com/coscms/forms/widgets"
)

// BaseWidget creates a Widget based on style and inpuType parameters, both defined in the common package.
func BaseWidget(style, inputType, tmplName string) *widgets.Widget {
	cachedKey := style + ", " + inputType + ", " + tmplName
	tmpl, err := common.GetOrSetCachedTemplate(cachedKey, func() (*template.Template, error) {
		var (
			fpath = common.TmplDir(style) + "/" + style + "/"
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
	if a != nil {
		attrs = a.String()
		if len(attrs) > 0 {
			attrs = ` ` + attrs
		}
	}
	return attrs
}
