package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"
)

// Animals is multiple Animal struct
type Animals struct {
	Animals []Animal
}

// Animal is `t_animal` table mapping struct
type Animal struct {
	ID   int
	Name string
}

func animalHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET", "":
		animalGET(w, r)
	case "POST":
		animalPOST(w, r)
	}
}

func animalGET(w http.ResponseWriter, r *http.Request) {
	fID := r.URL.Query().Get("id")

	animals, err := selectAnimals(fID)
	if err != nil {
		failed(err, w)
		return
	}
	tmpl, err := template.ParseFiles("./template/animal.html")
	if err != nil {
		failed(err, w)
		return
	}
	tmpl.Execute(w, animals)
}

func animalPOST(w http.ResponseWriter, r *http.Request) {
	fName := r.FormValue("name")
	if fName == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("required `name` parameter"))
		return
	}

	if err := insertAnimals(fName); err != nil {
		failed(err, w)
		return
	}

	animals, err := selectAnimals("")
	if err != nil {
		failed(err, w)
		return
	}
	tmpl, err := template.ParseFiles("./template/animal.html")
	if err != nil {
		failed(err, w)
		return
	}
	tmpl.Execute(w, animals)
}

func insertAnimals(fName string) error {
	db, err := sql.Open("mysql",
		fmt.Sprintf("%s:%s@tcp(%s:3306)/animal", cfg.Animal.User, cfg.Animal.Password, cfg.Animal.Host))
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	insertAnimal, err := db.Prepare("INSERT INTO t_animal (name) values(?)")
	if err != nil {
		return err
	}

	tx, err := db.Begin()
	if err != nil {
		return err
	}

	_, err = tx.Stmt(insertAnimal).Exec(fName)
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()

	return nil
}

func selectAnimals(fID string) (Animals, error) {
	db, err := sql.Open("mysql",
		fmt.Sprintf("%s:%s@tcp(%s:3306)/animal", cfg.Animal.User, cfg.Animal.Password, cfg.Animal.Host))
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	sql := "SELECT id, name FROM t_animal"
	if fID != "" {
		sql = fmt.Sprintf("%s WHERE id=%s", sql, fID)
	}
	rows, err := db.Query(sql)
	if err != nil {
		return Animals{}, err
	}
	defer rows.Close()

	var animals Animals
	for rows.Next() {
		var animal Animal
		rows.Scan(&animal.ID, &animal.Name)

		animals.Animals = append(animals.Animals, animal)
	}
	return animals, nil
}

func failed(err error, w http.ResponseWriter) {
	log.Print(err)
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte("Failed"))
}
