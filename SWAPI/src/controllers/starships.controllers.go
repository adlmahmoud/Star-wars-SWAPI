package controllers

import (
	"fmt"
	"net/http"
	"swapi/src/helpers"
	"swapi/src/services"
	"swapi/src/templates"
)

func DisplayStarships(w http.ResponseWriter, r *http.Request) {
	char := "starships/"
	starships, status, err := services.SearchStarship(char)
	if status != http.StatusOK || err != nil {
		helpers.RedirectToError(w, r, status, "Erreur lors de la récupération des albums")
		fmt.Println(err)
		return
	}
	templates.RenderTemplate(w, r, "starships", starships)
}
