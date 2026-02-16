package services

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"swapi/src/models"
	"sync"
)

var favMutex sync.Mutex

const favFile = "./../../favorites.json"

func GetFavorites() ([]models.Favorite, error) {
	favMutex.Lock()
	defer favMutex.Unlock()

	if _, err := os.Stat(favFile); os.IsNotExist(err) {
		return []models.Favorite{}, nil
	}

	data, err := ioutil.ReadFile(favFile)
	if err != nil {
		return nil, err
	}

	var favs []models.Favorite
	json.Unmarshal(data, &favs)
	return favs, nil
}

func ToggleFavorite(newFav models.Favorite) (bool, error) {
	favs, err := GetFavorites()
	if err != nil {
		return false, err
	}

	favMutex.Lock()
	defer favMutex.Unlock()

	for i, f := range favs {
		if f.ID == newFav.ID && f.Type == newFav.Type {
			favs = append(favs[:i], favs[i+1:]...)
			saveFavorites(favs)
			return false, nil
		}
	}

	favs = append(favs, newFav)
	saveFavorites(favs)
	return true, nil
}

func saveFavorites(favs []models.Favorite) error {
	data, _ := json.MarshalIndent(favs, "", "  ")
	return ioutil.WriteFile(favFile, data, 0644)
}
