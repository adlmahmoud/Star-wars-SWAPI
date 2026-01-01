package services

/*
To do:
recuperer l'api:
URL= "https://swapi.dev/api/planets/"
*/
import (
	"encoding/json"
	"fmt"
	"net/http"
	"swapi/src/models"
	"time"
)

func SearchPlanets(char string) (*models.Planets, int, error) {
	client := http.Client{
		Timeout: 5 * time.Second,
	}
	url := "https://swapi.dev/api/" + char
	request, requestErr := http.NewRequest(http.MethodGet, url, nil)
	if requestErr != nil {
		fmt.Printf("Erreur initialisiation requete - %s\n", requestErr.Error())
	}

	response, responseErr := client.Do(request)
	if responseErr != nil {
		fmt.Printf("Erreur requete HTTP - %s\n", responseErr.Error())
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		fmt.Printf("Erreur init requete - %d, status %s\n", response.StatusCode, response.Status)
	}

	var planet models.Planets

	decoderErr := json.NewDecoder(response.Body).Decode(&planet)
	if decoderErr != nil {
		fmt.Printf("Erreur decodage JSON - %s\n", decoderErr.Error())
	}
	return &planet, response.StatusCode, nil
}
