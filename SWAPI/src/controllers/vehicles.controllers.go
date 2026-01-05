package controllers

import (
	"fmt"
	"net/http"
	"swapi/src/helpers"
	"swapi/src/services"
	"swapi/src/templates"
)

func DisplayVehicles(w http.ResponseWriter, r *http.Request) {
	page := r.URL.Query().Get("page")
	if page == "" {
		page = "1"
	}
	char := "vehicles/"
	vehicles, status, err := services.SearchVehicle(char, page)
	if status != http.StatusOK || err != nil {
		helpers.RedirectToError(w, r, status, "Erreur lors de la récupération des albums")
		fmt.Println(err)
		return
	}
	templates.RenderTemplate(w, r, "vehicles", vehicles)
}
