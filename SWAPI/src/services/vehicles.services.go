package services

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"swapi/src/models"
	"time"
)

type VehicleFilters struct {
	CostMax   string
	CargoMin  string
	LengthMax string
}

func SearchVehicle(char string, page string, filters VehicleFilters) (*models.VeheiclesRep, int, error) {
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

	var vehicleRep models.VeheiclesRep

	decoderErr := json.NewDecoder(response.Body).Decode(&vehicleRep)
	if decoderErr != nil {
		fmt.Printf("Erreur decodage JSON - %s\n", decoderErr.Error())
	}
	if filters.CostMax == "" && filters.CargoMin == "" && filters.LengthMax == "" {
		return &vehicleRep, response.StatusCode, nil
	}

	fCostMax := 999999999
	if filters.CostMax != "" {
		val, _ := strconv.Atoi(filters.CostMax)
		if val > 0 {
			fCostMax = val
		}
	}

	fCargoMin := 0
	if filters.CargoMin != "" {
		val, _ := strconv.Atoi(filters.CargoMin)
		fCargoMin = val
	}

	fLengthMax := 999999.0
	if filters.LengthMax != "" {
		val, _ := strconv.ParseFloat(filters.LengthMax, 64)
		if val > 0 {
			fLengthMax = val
		}
	}

	var filteredResults []models.Veheicles

	for _, v := range vehicleRep.Results {
		keep := true
		vCost := 0
		if v.Cost_in_credits != "unknown" && v.Cost_in_credits != "" {
			vCost, _ = strconv.Atoi(v.Cost_in_credits)
		}
		if vCost > fCostMax {
			keep = false
		}
		vCargo := 0
		if v.Cargo_capacity != "unknown" && v.Cargo_capacity != "" {
			vCargo, _ = strconv.Atoi(v.Cargo_capacity)
		}
		if vCargo < fCargoMin {
			keep = false
		}
		vLength := 0.0
		if v.Length != "unknown" && v.Length != "" {
			cleanLen := strings.Replace(v.Length, ",", ".", -1)
			vLength, _ = strconv.ParseFloat(cleanLen, 64)
		}
		if vLength > fLengthMax {
			keep = false
		}

		if keep {
			filteredResults = append(filteredResults, v)
		}
	}

	vehicleRep.Results = filteredResults

	return &vehicleRep, response.StatusCode, nil
}
