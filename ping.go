package main

import (
	"net/http"
	"log"
	"html/template"
)

func pingHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("%s /ping\n", r.Method)
	tmpl, err := template.ParseFiles("./template/ping.html")
	if err != nil {
		failed(err, w)
		return
	}
	w.WriteHeader(http.StatusOK)
	tmpl.Execute(w, Result{Message: "Pong!"})
}

