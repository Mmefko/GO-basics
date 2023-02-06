package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	Name       string   `json:"name"`
	Address    string   `json:"addr"`
	Age        int      `json:"age"`
	FaveColors []string `json:"favecolors"`
}

func decodeExample() {
	//Create some JSON data to decode
	data := []byte(`
		{
			"name" : "Mary Maryivich",
			"addr" : "2 Little Town Ave.",
			"age" : 49,
			"favecolors": ["Green", "Red"]
		}
	`)

	//JSON wll be decoded into struct
	var p person

	//Test to see if the JSON is valid, and then decode
	valid := json.Valid(data)
	if valid {
		json.Unmarshal(data, &p)
		fmt.Printf("%#v\n", p)
	}
	//data can also be decoded into a map structure
	var m map[string]interface{}

	//Unmarshal into a map
	json.Unmarshal(data, &m)
	fmt.Printf("%#v\n", m)
	for key, val := range m {
		fmt.Printf("%v : %v\n", key, val)
	}
}

func main() {
	decodeExample()
}
