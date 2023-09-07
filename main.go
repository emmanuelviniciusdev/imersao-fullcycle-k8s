package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		name := os.Getenv("NAME")
		wish := os.Getenv("WISH")

		message := "[V4] Olá!! My name is " + name + " and I wish: " + wish

		w.Write([]byte(message))
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
