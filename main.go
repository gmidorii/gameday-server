package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/BurntSushi/toml"
	_ "github.com/go-sql-driver/mysql"
)

const (
	port    = "8080"
	cfgFile = "./config.toml"
)

var cfg Config

// Config is setting for this web server
// `config.toml` struct
type Config struct {
	Animal AnimalCfg
}

// AnimalCfg is setting for animal db
type AnimalCfg struct {
	User     string
	Password string
	Host     string
}

func init() {
	_, err := toml.DecodeFile(cfgFile, &cfg)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("success loading config file: %s", cfgFile)
}

func main() {
	http.HandleFunc("/ping", pingHandler)
	http.HandleFunc("/animal", animalHandler)

	log.Printf("start game server port: %s", port)
	err := http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
	if err != nil {
		log.Fatal(err)
	}
}

func pingHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("OK"))
}
