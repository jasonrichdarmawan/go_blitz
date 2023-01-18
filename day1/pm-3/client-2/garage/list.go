package garage

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/jasonrichdarmawan/go_blitz/day1/pm-3/common/model"
)

func List(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	decoder := json.NewDecoder(r.Body)

	var requestBody model.GarageUserId
	err := decoder.Decode(&requestBody)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if requestBody.UserId == "" {
		http.Error(w, "user_id can't be empty", http.StatusBadRequest)
		return
	}

	garageList, err := GetService().List(context.Background(), &requestBody)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	responseBody, err := json.Marshal(garageList)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(responseBody)
}
