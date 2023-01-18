package main

import (
	"log"
	"net/http"

	"github.com/jasonrichdarmawan/go_blitz/day1/pm-3/client-2/garage"
	"github.com/jasonrichdarmawan/go_blitz/day1/pm-3/client-2/user"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/user/register", user.Register)
	mux.HandleFunc("/user/list", user.List)

	mux.HandleFunc("/garage/add", garage.Add)
	mux.HandleFunc("/garage/list", garage.List)

	err := http.ListenAndServe("localhost:8080", mux)
	if err != nil {
		log.Panicln(err)
	}
}
