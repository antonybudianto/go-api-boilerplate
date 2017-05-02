package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Starting the app...")

	config, err := LoadConfig("./config.json")
	if err != nil {
		panic(err)
	}

	addr := config.Host + ":" + config.Port

	router := GetRouter()

	http.ListenAndServe(addr, router)
}
