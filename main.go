package main

import (
	"log"
	"net/http"
	"os"
	"time"
)

var startedAt = time.Now()

func main() {
	http.HandleFunc("/", RootRequestHandler)
	http.HandleFunc("/most-adorable-countries-on-the-planet", MostAdorableCountriesOnThePlanetRequestHandler)
	http.HandleFunc("/secrets", SecretsRequestHandler)
	http.HandleFunc("/healthz", HealthzRequestHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func HealthzRequestHandler(w http.ResponseWriter, r *http.Request) {
	uptime := time.Since(startedAt)

	statusCode := 204

	//livenessProbeError := uptime.Seconds() > 25
	readinessProbeError := uptime.Seconds() < 25

	if readinessProbeError {
		statusCode = 500
	}

	w.WriteHeader(statusCode)
}

func SecretsRequestHandler(w http.ResponseWriter, r *http.Request) {
	listSecret := [...]string{os.Getenv("SECRET_1"), os.Getenv("SECRET_2"), os.Getenv("SECRET_3")}

	for _, secret := range listSecret {
		w.Write([]byte(secret))
		w.Write([]byte("\n"))
	}
}

func MostAdorableCountriesOnThePlanetRequestHandler(w http.ResponseWriter, r *http.Request) {
	data, err := os.ReadFile("./countries/countries.txt")

	if err != nil {
		log.Fatal(err)
	}

	w.Write(data)
}

func RootRequestHandler(w http.ResponseWriter, r *http.Request) {
	name := os.Getenv("NAME")
	wish := os.Getenv("WISH")

	message := "[V4] Olá!! My name is " + name + " and I wish: " + wish

	w.Write([]byte(message))
}
