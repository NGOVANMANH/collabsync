package main

import (
	"log"
	"net/http"

	"weather-proxy/internal/weather/rest"
)

func main() {
	mux := http.NewServeMux()
	rest.RegisterHelloRoutes(mux)

	log.Println("Starting server on :8081")
	if err := http.ListenAndServe(":8081", mux); err != nil {
		log.Fatal(err)
	}
}
