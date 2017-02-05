package handlers

import (
	"html/template"
	"net/http"
)

var tpl *template.Template
var err error

func loadTemplates() {
	tpl, err = template.New("").Funcs(template.FuncMap{
		"currency": CurrencyExpander,
	}).ParseGlob("views/*.html")

	if err != nil {
		panic(err)
	}
}

func render(w http.ResponseWriter, name string, data interface{}) {
	if tpl == nil {
		loadTemplates()
	}
	tpl.ExecuteTemplate(w, name, data)
}
