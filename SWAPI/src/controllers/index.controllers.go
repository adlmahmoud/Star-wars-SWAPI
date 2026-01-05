package controllers

import (
	"net/http"
	"swapi/src/templates"
)

func DisplayIndex(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	templates.RenderTemplate(w, r, "index", nil)
}
