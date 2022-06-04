package main

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type RequestBody struct {
	Money      int    `json:"money"`
	CandyType  string `json:"candyType"`
	CandyCount int    `json:"candyCount"`
}

func getClient() (client *http.Client) {
	//data, err := ioutil.ReadAll("../ca/minica.pem")
	c, err := tls.LoadX509KeyPair("./cert.pem", "./key.pem")
	if err != nil {
		log.Fatal(err)
	}
	certs := []tls.Certificate{c}
	if len(certs) == 0 {
		client = &http.Client{}
		return
	}
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{Certificates: certs, InsecureSkipVerify: true},
	}
	client = &http.Client{Transport: tr}
	return client
}

func main() {
	client := getClient()
	var rBody RequestBody
	rBody.Money = 50
	rBody.CandyType = "AA"
	rBody.CandyCount = 2

	jsonValue, _ := json.Marshal(rBody)
	resp, err := client.Post("https://127.0.0.1:3333/buy_candy", "application/json", bytes.NewBuffer(jsonValue))
	if err != nil {
		log.Fatal(err)
	}
	data, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(data))
	resp.Body.Close()
}
