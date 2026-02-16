package routers

import (
	"net/http"
	"swapi/src/controllers"
)

func SearchRouters(routers *http.ServeMux) {
	routers.HandleFunc("/search", controllers.SearchHandler)
}
