package routers

import (
	"net/http"
	"swapi/src/controllers"
)

func VehicilesRouters(reouters *http.ServeMux) {
	reouters.HandleFunc("/vehicles", controllers.DisplayVehicles)
}
