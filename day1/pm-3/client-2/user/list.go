package user

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/golang/protobuf/ptypes/empty"
)

func List(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	userList, err := GetService().List(context.Background(), new(empty.Empty))
	if err != nil {
		http.Error(w, err.Error(), http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	responseBody, err := json.Marshal(userList)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(responseBody)
}
