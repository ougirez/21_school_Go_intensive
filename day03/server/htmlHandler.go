package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

type httpResponse struct {
	Places []Place `json:"places"`
	Total  int     `json:"total"`
	Page   int     `json:"-"`
}

func htmlError(w http.ResponseWriter, r *http.Request) {
	response := fmt.Sprintf("Invalid 'page' value '%s'", r.URL.Query()["page"][0])
	w.WriteHeader(400)
	w.Write([]byte(response))
	return
}

func htmlGetPlaces(w http.ResponseWriter, r *http.Request) {
	var response httpResponse
	if len(r.URL.Query()["page"]) != 1 {
		return
	}
	page, err := strconv.Atoi(r.URL.Query()["page"][0])
	if err != nil || page < 0 {
		htmlError(w, r)
	}
	places, totalPlaces, _ := GetPlaces(10, page*10)
	if err != nil || page > totalPlaces/10 {
		htmlError(w, r)
	}
	response.Places = places
	response.Total = totalPlaces
	response.Page = page

	funcMap := template.FuncMap{
		"inc": func(i int) int {
			return i + 1
		},
		"dec": func(i int) int {
			return i - 1
		},
		"div": func(a, b int) int {
			return a / b
		},
	}
	t, err := template.New("page.html").Funcs(funcMap).ParseFiles("page.html")
	if err != nil {
		log.Fatal(err)
	}
	err = t.Execute(w, response)
	if err != nil {
		log.Println("Error executing template :", err)
		return
	}
}
