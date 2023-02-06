package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type person struct {
	Name       string   `json:"name"`
	Address    string   `json:"addr"`
	Age        int      `json:"-"`
	FaveColors []string `json:"favecolors,omitempty"`
}

func encodeExample() {
	//Create some prople data
	people := []person{
		{"John Johnovich", "1 Big City St.", 28, nil},
		{"Mary Maryivich", "2 Little Town Ave.", 49, []string{"Green", "Red"}},
	}

	//Marshal is used to convert a data structure to JSON format
	//MarshalIndent is used to format the JSON string with indentation
	result, err := json.Marshal(people)
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Printf("%s\n", result)

	result, err = json.MarshalIndent(people, "", "\t")
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Printf("%s\n", result)
}

func main() {
	encodeExample()
}
