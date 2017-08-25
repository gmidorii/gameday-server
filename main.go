package main

import (
	"database/sql"
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

type Config struct {
	DbCfg DB
}

type DB struct {
	User     string
	Password string
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

func animalHandler(w http.ResponseWriter, r *http.Request) {
	var cfg Config
	toml.DecodeFile(cfgFile, &cfg)
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@localhost/animal", cfg.DbCfg.User, cfg.DbCfg.Password))
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
}
