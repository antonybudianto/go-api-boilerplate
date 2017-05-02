package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Starting the app...")

	router := GetRouter()

	http.ListenAndServe("localhost:1337", router)
}
