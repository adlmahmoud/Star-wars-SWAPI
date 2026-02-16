package controllers

import (
	"encoding/json"
	"net/http"
	"swapi/src/models"
	"swapi/src/services"
	"swapi/src/templates"
)

func DisplayFavorites(w http.ResponseWriter, r *http.Request) {
	favs, err := services.GetFavorites()
	if err != nil {
		favs = []models.Favorite{}
	}
	templates.RenderTemplate(w, r, "favorites", favs)
}

func ApiGetFavorites(w http.ResponseWriter, r *http.Request) {
	favs, _ := services.GetFavorites()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(favs)
}

func ApiToggleFavorite(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var fav models.Favorite

	if err := json.NewDecoder(r.Body).Decode(&fav); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	added, err := services.ToggleFavorite(fav)
	if err != nil {
		http.Error(w, "Erreur serveur", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]bool{"added": added})
}
