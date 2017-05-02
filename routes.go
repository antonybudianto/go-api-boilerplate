package main

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// GetRouter return the router of our API
func GetRouter() *httprouter.Router {
	router := httprouter.New()

	router.GET("/", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		fmt.Fprint(w, "{ \"message\": \"Welcome\" }")
	})

	return router
}
