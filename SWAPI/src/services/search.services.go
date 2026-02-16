package services

import (
	"crypto/tls"
	"encoding/json"
	"net/http"
	"swapi/src/models"
	"sync"
	"time"
)

func fetchCategory(url string, target interface{}, wg *sync.WaitGroup) {
	defer wg.Done()

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := http.Client{Timeout: 5 * time.Second, Transport: tr}

	resp, err := client.Get(url)
	if err != nil || resp.StatusCode != 200 {
		return
	}
	defer resp.Body.Close()

	json.NewDecoder(resp.Body).Decode(target)
}

func SearchGlobal(query string) *models.GlobalSearchResult {
	var result models.GlobalSearchResult
	result.Query = query

	var charResp models.CharcResponse
	var filmResp models.FilmsResponse
	var planetResp models.PlanetsResponse
	var specieResp models.SpeciesRep
	var vehicleResp models.VeheiclesRep
	var starshipResp models.StarshipsResponse

	var wg sync.WaitGroup

	wg.Add(6)

	go fetchCategory("https://swapi.dev/api/people/?search="+query, &charResp, &wg)
	go fetchCategory("https://swapi.dev/api/films/?search="+query, &filmResp, &wg)
	go fetchCategory("https://swapi.dev/api/planets/?search="+query, &planetResp, &wg)
	go fetchCategory("https://swapi.dev/api/species/?search="+query, &specieResp, &wg)
	go fetchCategory("https://swapi.dev/api/vehicles/?search="+query, &vehicleResp, &wg)
	go fetchCategory("https://swapi.dev/api/starships/?search="+query, &starshipResp, &wg)

	wg.Wait()

	result.Characters = charResp.Results
	result.Films = filmResp.Results
	result.Planets = planetResp.Results
	result.Species = specieResp.Results
	result.Vehicles = vehicleResp.Results
	result.Starships = starshipResp.Results

	result.Count = len(result.Characters) + len(result.Films) + len(result.Planets) +
		len(result.Species) + len(result.Vehicles) + len(result.Starships)

	return &result
}
