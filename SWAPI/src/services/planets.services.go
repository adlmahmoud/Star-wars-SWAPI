package services

/*
To do:
recuperer l'api:
URL= "https://swapi.dev/api/planets/"
*/
import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"net/http"
	"swapi/src/models"
	"time"
)

func SearchPlanets(char string, page string) (*models.PlanetsResponse, int, error) {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	client := http.Client{
		Timeout:   5 * time.Second,
		Transport: tr,
	}
	url := "https://swapi.dev/api/" + char + "?page=" + page
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

	var planetRep models.PlanetsResponse

	decoderErr := json.NewDecoder(response.Body).Decode(&planetRep)
	if decoderErr != nil {
		fmt.Printf("Erreur decodage JSON - %s\n", decoderErr.Error())
	}
	return &planetRep, response.StatusCode, nil
}
