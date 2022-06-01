package main

import (
	"encoding/json"
	"encoding/xml"
	"io"
	"log"
	"os"
	"strings"
)

type DBReader interface {
	parseFile() Recipes
}

type JsonReader struct {
	fileName string
}

type XMLReader struct {
	fileName string
}

func (r JsonReader) parseFile() Recipes {
	var recipes Recipes

	jsonFile, err := os.Open(r.fileName)
	if err != nil {
		log.Fatal(err)
	}
	byteValue, _ := io.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &recipes)
	jsonFile.Close()
	return recipes
}

func (r XMLReader) parseFile() Recipes {
	var recipes Recipes

	xmlFile, err := os.Open(r.fileName)
	if err != nil {
		log.Fatal(err)
	}
	byteValue, _ := io.ReadAll(xmlFile)
	xml.Unmarshal(byteValue, &recipes)
	xmlFile.Close()
	return recipes
}

func getReader(fileName string) DBReader{
	if strings.HasSuffix(fileName, ".json") {
		reader := new(JsonReader)
		reader.fileName = fileName
		return reader
	} else if strings.HasSuffix(fileName, ".xml") {
		reader := new(XMLReader)
		reader.fileName = fileName
		return reader
	} else {
		log.Fatal("Invalid extension of file")
	}
	return nil
}
