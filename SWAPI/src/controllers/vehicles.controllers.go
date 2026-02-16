package controllers

import (
	"fmt"
	"net/http"
	"swapi/src/helpers"
	"swapi/src/models"
	"swapi/src/services"
	"swapi/src/templates"
)

// Structure pour envoyer Données + Filtres à la page
type VehiclePageData struct {
	Data    *models.VeheiclesRep
	Filters services.VehicleFilters
}

func DisplayVehicles(w http.ResponseWriter, r *http.Request) {
	page := r.URL.Query().Get("page")
	if page == "" {
		page = "1"
	}

	// 1. Récupérer les filtres du formulaire
	filters := services.VehicleFilters{
		CostMax:   r.URL.Query().Get("cost_max"),
		CargoMin:  r.URL.Query().Get("cargo_min"),
		LengthMax: r.URL.Query().Get("length_max"),
	}

	char := "vehicles/"

	// 2. Appeler le service AVEC les filtres
	vehicles, status, err := services.SearchVehicle(char, page, filters)

	if status != http.StatusOK || err != nil {
		helpers.RedirectToError(w, r, status, "Erreur lors de la récupération des véhicules")
		fmt.Println(err)
		return
	}

	// 3. Envoyer le tout au template
	pageData := VehiclePageData{
		Data:    vehicles,
		Filters: filters,
	}

	templates.RenderTemplate(w, r, "vehicles", pageData)
}
