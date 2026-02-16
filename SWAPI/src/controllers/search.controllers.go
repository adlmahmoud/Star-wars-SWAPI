package controllers

import (
	"net/http"
	"swapi/src/services"
	"swapi/src/templates"
)

func SearchHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("query")

	if query == "" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	results := services.SearchGlobal(query)
	templates.RenderTemplate(w, r, "search", results)
}
