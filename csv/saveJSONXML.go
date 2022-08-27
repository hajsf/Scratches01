//go:build ignore
// +build ignore

package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"strconv"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {

	// create and populate a map from dummy JSON data

	dataStr := `{"Name":"Dummy","Age":0}`

	personMap := make(map[string]interface{})

	err := json.Unmarshal([]byte(dataStr), &personMap)

	if err != nil {
		panic(err)
	}

	var onePerson Person

	// convert map to Person struct
	onePerson.Name = fmt.Sprintf("%s", personMap["Name"])
	onePerson.Age, _ = strconv.Atoi(fmt.Sprintf("%v", personMap["Age"]))

	jsonData, err := json.Marshal(onePerson)

	if err != nil {
		panic(err)
	}

	// sanity check
	fmt.Println(string(jsonData))

	// write to JSON file

	jsonFile, err := os.Create("./Person.json")

	if err != nil {
		panic(err)
	}
	defer jsonFile.Close()

	jsonFile.Write(jsonData)
	jsonFile.Close()
	fmt.Println("JSON data written to ", jsonFile.Name())

	// write to XML file

	xmlFile, err := os.Create("./Person.xml")
	if err != nil {
		panic(err)
	}
	defer xmlFile.Close()

	xmlWriter := io.Writer(xmlFile)

	enc := xml.NewEncoder(xmlWriter)
	enc.Indent("  ", "    ")
	if err := enc.Encode(onePerson); err != nil {
		fmt.Printf("error: %v\n", err)
	}

	xmlFile.Close()
	fmt.Println("XML data written to ", xmlFile.Name())

}
