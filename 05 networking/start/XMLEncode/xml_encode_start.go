package main

import (
	"encoding/xml"
	"fmt"
	"log"
)

type person struct {
	XMLName    xml.Name `xml:"personaldata"`
	FirstName  string   `xml:"firstname"`
	LastName   string   `xml:"lasttname"`
	Address    string   `xml:"addr"`
	Age        int      `xml:"age,attr"`
	FaveColors []string `xml:"faveColors"`
}

func encodeExample() {
	//Create some prople data
	people := []person{
		{FirstName: "John", LastName: "Johnovich", Address: "1 Big City St.", Age: 28, FaveColors: nil},
		{FirstName: "Mary", LastName: "Maryivich", Address: "2 Little Town Ave.", Age: 49,
			FaveColors: []string{"Green", "Red"}},
	}

	//The MarshalIndent function indents the XML text
	result, err := xml.MarshalIndent(people, "", "\t")
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Printf("%s%s\n", xml.Header, result)
}

func main() {
	encodeExample()
}
