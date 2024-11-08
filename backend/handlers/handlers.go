package handlers

import "html/template"

var tmpl *template.Template

// SetTemplate sets the template variable for use in handlers
func SetTemplate(t *template.Template) {
	tmpl = t
}
