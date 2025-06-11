package rest

import (
	"net/http"

	"weather-proxy/internal/weather/service"
)

func RegisterHelloRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/hello", helloHandler)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	message := service.GetHelloMessage(name)

	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(message))
}
