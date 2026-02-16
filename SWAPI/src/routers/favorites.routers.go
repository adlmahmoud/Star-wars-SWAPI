package routers

import (
	"net/http"
	"swapi/src/controllers"
)

func FavoritesRouters(routers *http.ServeMux) {
	routers.HandleFunc("/favorites", controllers.DisplayFavorites)

	routers.HandleFunc("/api/favorites", controllers.ApiGetFavorites)
	routers.HandleFunc("/api/favorites/toggle", controllers.ApiToggleFavorite)
}
