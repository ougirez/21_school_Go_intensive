package main

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type RequestBody struct {
	Money      int    `json:"money"`
	CandyType  string `json:"candyType"`
	CandyCount int    `json:"candyCount"`
}

type SuccessResponseBody struct {
	Thanks string `json:"thanks"`
	Change int    `json:"change"`
}

type FailResponseBody struct {
	Error string `json:"error"`
}

var CandyPrices = map[string]int{
	"CE": 10,
	"AA": 15,
	"NT": 17,
	"DE": 21,
	"YR": 23,
}

func buyCandy(w http.ResponseWriter, r *http.Request) {
	var requestBody RequestBody

	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		log.Fatal(err)
	}
	if _, ok := CandyPrices[requestBody.CandyType]; ok == false {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(FailResponseBody{
			Error: "Invalid candy type",
		})
		return
	}
	if requestBody.CandyCount < 0 {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(FailResponseBody{
			Error: "Negative candy count is forbidden",
		})
		return
	}
	if requestBody.Money >= CandyPrices[requestBody.CandyType]*requestBody.CandyCount {
		w.WriteHeader(200)
		json.NewEncoder(w).Encode(SuccessResponseBody{
			Thanks: "Thank you!",
			Change: requestBody.Money - CandyPrices[requestBody.CandyType]*requestBody.CandyCount,
		})
		return
	} else {
		w.WriteHeader(402)
		diff := CandyPrices[requestBody.CandyType]*requestBody.CandyCount - requestBody.Money
		json.NewEncoder(w).Encode(FailResponseBody{
			Error: fmt.Sprintf("You need %d more money!", diff),
		})
		return
	}
}

func getServer() *http.Server {
	server := &http.Server{
		Addr: ":3333",
		TLSConfig: &tls.Config{
			ClientAuth: tls.RequestClientCert,
		},
	}
	return server
}

func main() {
	http.HandleFunc("/buy_candy", buyCandy)

	server := getServer()
	err := server.ListenAndServeTLS("minica.pem", "minica-key.pem")
	if err != nil {
		log.Fatal(err)
	}
}
