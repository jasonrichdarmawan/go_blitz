package garage

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/jasonrichdarmawan/go_blitz/day1/pm-3/common/model"
)

func addValidator(requestBody *model.GarageAndUserId) error {
	if requestBody.UserId == "" {
		return errors.New("user_id can't be empty")
	}
	if requestBody.Garage.Id == "" {
		return errors.New("garage.id can't be empty")
	}
	if requestBody.Garage.Name == "" {
		return errors.New("name can't be empty")
	}
	if requestBody.Garage.Coordinate.Latitude == 0 {
		return errors.New("garage.latitude can't be empty")
	}
	if requestBody.Garage.Coordinate.Longitude == 0 {
		return errors.New("garage longitude can't be empty")
	}
	return nil
}

func Add(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	decoder := json.NewDecoder(r.Body)

	var requestBody model.GarageAndUserId
	err := decoder.Decode(&requestBody)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err = addValidator(&requestBody); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, err = GetService().Add(context.Background(), &requestBody)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
