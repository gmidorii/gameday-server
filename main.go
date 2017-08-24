package main

import (
	"fmt"
	"log"
	"net/http"
)

const port = "8080"

func main() {
	http.HandleFunc("/ping", pingHandler)

	log.Printf("start game server port: %s", port)
	err := http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
	if err != nil {
		log.Fatal(err)
	}
}

func pingHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("OK"))
}
