package main

import (
	"database/sql"
	"log"
)

func connectToDB() *sql.DB {
	conninfo := "dbname=Day03 user=ougirez password=12345 host=localhost sslmode=disable"
	db, err := sql.Open("postgres", conninfo)
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func totalCount() int {
	db := connectToDB()
	defer db.Close()
	var total int
	err := db.QueryRow("SELECT count(*) FROM places").Scan(&total)
	if err != nil {
		log.Fatal(err)
	}
	return total
}

type Store interface {
	GetPlaces(limit int, offset int) ([]Place, int, error)
}

func GetPlaces(limit int, offset int) ([]Place, int, error) {
	db := connectToDB()
	defer db.Close()
	var places []Place
	rows, err := db.Query("SELECT * FROM places LIMIT $1 OFFSET $2", limit, offset)
	defer rows.Close()
	if err != nil {
		log.Fatal(err)
	}
	if err != nil {
		log.Fatal(err)
	}
	totalPlaces := totalCount()
	for rows.Next() {
		var place Place
		err = rows.Scan(&place.Id, &place.Name, &place.Phone, &place.Address, &place.Location.Longitude, &place.Location.Latitude)
		if err != nil {
			log.Fatal(err)
		}
		places = append(places, place)
	}
	return places, totalPlaces, nil
}
