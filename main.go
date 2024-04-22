package main

import (
	"calculateFactorial/routes"
	"log"
	"net/http"
)

func main() {
	router := routes.NewRouter(routes.AllRoutes())
	log.Fatal(http.ListenAndServe(":8989", router))
}
