package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

type JsonResponse struct {
	Places   []Place `json:"places"`
	Name     string  `json:"name"`
	Total    int     `json:"total"`
	NextPage int     `json:"next_page,omitempty"`
	PrevPage int     `json:"prev_page,omitempty"`
	LastPage int     `json:"last_page"`
}

type JsonError struct {
	Error string `json:"error"`
}

func PageError(w http.ResponseWriter, r *http.Request) {
	resp := JsonError{
		Error: fmt.Sprintf("Invalid 'page' value '%s'", r.URL.Query()["page"][0]),
	}
	w.WriteHeader(400)
	jsonResponse, err := json.Marshal(resp)
	if err != nil {
		log.Fatal(err)
	}
	w.Write(jsonResponse)
	return
}

func apiGetPlaces(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if len(r.URL.Query()["page"]) != 1 {
		return
	}
	page, err := strconv.Atoi(r.URL.Query()["page"][0])
	if err != nil || page < 0 {
		PageError(w, r)
		return
	}
	places, totalPlaces, _ := GetPlaces(10, page*10)
	if page < 0 || page > totalPlaces/10 {
		PageError(w, r)
		return
	}
	var resp JsonResponse
	resp.Total = totalPlaces
	resp.Name = "Places"
	resp.Places = places
	if page > 0 {
		resp.PrevPage = page - 1
	}
	resp.LastPage = totalPlaces / 10
	if page != resp.LastPage {
		resp.NextPage = page + 1
	}
	jsonResponse, err := json.Marshal(resp)
	if err != nil {
		log.Fatal(err)
	}
	w.Write(jsonResponse)
	return
}
