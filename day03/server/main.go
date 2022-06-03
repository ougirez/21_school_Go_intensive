package main

import (
	"fmt"
	_ "github.com/lib/pq"
	"net/http"
)

type Place struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Address  string `json:"address"`
	Phone    string `json:"phone"`
	Location struct {
		Longitude float32 `json:"lon"`
		Latitude  float32 `json:"lat"`
	} `json:"location"`
}

func main() {
	http.HandleFunc("/", htmlGetPlaces)
	http.HandleFunc("/api/places/", apiGetPlaces)
	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		fmt.Println(err)
	}
}
