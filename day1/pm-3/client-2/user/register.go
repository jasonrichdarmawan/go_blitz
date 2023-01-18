package user

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/jasonrichdarmawan/go_blitz/day1/pm-3/common/model"
)

func registerValidator(requestBody *model.User) error {
	if requestBody.Id == "" {
		return errors.New("id can't be empty")
	}
	if requestBody.Name == "" {
		return errors.New("name can't be empty")
	}
	if requestBody.Password == "" {
		return errors.New("password can't be empty")
	}
	if requestBody.Gender == 0 {
		return errors.New("gender can't be empty")
	}

	return nil
}

func Register(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	decoder := json.NewDecoder(r.Body)

	var requestBody model.User
	err := decoder.Decode(&requestBody)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err = registerValidator(&requestBody); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, err = GetService().Register(context.Background(), &requestBody)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
