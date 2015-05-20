package main

import (
	"github.com/greatcloak/panda"
	"log"
	"net/http"
)

func main() {
	log.Println("Starting web server....")

	r := panda.NewChainRouter()

	r.PathPrefix("/").Handler(http.FileServer(http.Dir("")))

	log.Println("Listening on port 8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatalln("Http server error:", err)
	}
}
