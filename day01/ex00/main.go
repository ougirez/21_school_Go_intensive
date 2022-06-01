package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

type Cakes struct {
	XMLName xml.Name `xml:"recipes" json:"-"`
	Cakes   []struct {
		Name        string `json:"name" xml:"name"`
		Time        string `json:"time" xml:"stovetime"`
		Ingredients []struct {
			Name  string `json:"ingredient_name" xml:"itemname"`
			Count string `json:"ingredient_count" xml:"itemcount"`
			Unit  string `json:"ingredient_unit,omitempty" xml:"itemunit"`
		} `json:"ingredients" xml:"ingredients>item"`
	} `json:"cake" xml:"cake"`
}

type DBReader interface {
	parseFile() Cakes
}

type JsonReader struct {
	fileName string
}

type XMLReader struct {
	fileName string
}

func (r JsonReader) parseFile() Cakes {
	var cakes Cakes

	jsonFile, err := os.Open(r.fileName)
	if err != nil {
		log.Fatal(err)
	}
	byteValue, _ := io.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &cakes)
	xmlCakes, err := xml.MarshalIndent(cakes, "", "    ")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(xmlCakes))

	jsonFile.Close()
	return cakes
}

func (r XMLReader) parseFile() Cakes {
	var cakes Cakes

	xmlFile, err := os.Open(r.fileName)
	if err != nil {
		log.Fatal(err)
	}
	byteValue, _ := io.ReadAll(xmlFile)
	xml.Unmarshal(byteValue, &cakes)

	jsonCakes, err := json.MarshalIndent(cakes, "", "    ")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(jsonCakes))
	xmlFile.Close()
	return cakes
}

func main() {
	if len(os.Args) != 3 {
		log.Fatal("Wrong number of arguments")
	}
	if os.Args[1] != "-f" {
		log.Fatal("You should use -f flag to specify the file")
	}

	if strings.HasSuffix(os.Args[2], ".json") {
		jsonReader := new(JsonReader)
		jsonReader.fileName = os.Args[2]
		jsonReader.parseFile()
	} else if strings.HasSuffix(os.Args[2], ".xml") {
		xmlReader := new(XMLReader)
		xmlReader.fileName = os.Args[2]
		xmlReader.parseFile()
	}

}
