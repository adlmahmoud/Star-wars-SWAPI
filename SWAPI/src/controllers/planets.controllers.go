package controllers

import (
	"fmt"
	"net/http"
	"swapi/src/helpers"
	"swapi/src/services"
	"swapi/src/templates"
)

func DisplayPlanets(w http.ResponseWriter, r *http.Request) {
	char := "planets/"
	planets, status, err := services.SearchPlanets(char)
	if status != http.StatusOK || err != nil {
		helpers.RedirectToError(w, r, status, "Erreur lors de la récupération des albums")
		fmt.Println(err)
		return
	}
	templates.RenderTemplate(w, r, "planets", planets)
}
