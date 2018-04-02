/*
	Kenneth Bailey
	4/1/18

	JSON and Go

	Notes:
		-

*/

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

type person struct {
	FirstName string `json:"fname"`
	MidName   string `json:"mname"`
	LastName  string `json:"lname"`
	Age       int    `json:"age"`
	FavColor  string `json:"favorite_color"`
}

func (p person) speak() {
	fmt.Printf("My name is %s and my favorite color is %s\n", p.FirstName, p.FavColor)
}

func main() {

	// Get byte array of json file
	file, err := ioutil.ReadFile("./example.json")
	if err != nil {
		log.Fatal(err)
	}

	// marshal json to our person struct
	var ken person
	json.Unmarshal(file, &ken)

	// json person speaks
	ken.speak()
}
