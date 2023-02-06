package main

import (
	"encoding/xml"
	"fmt"
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
	xmldata := `
	<personaldata age="49">
        <firstname>Mary</firstname>
        <lasttname>Maryivich</lasttname>
        <addr>2 Little Town Ave.</addr>
        <faveColors>Green</faveColors>
        <faveColors>Red</faveColors>
	</personaldata>
	`

	//Data will be decoded into a person struct
	var p person

	//Use the unmarshal function to decode raw XML data
	xml.Unmarshal([]byte(xmldata), &p)
	fmt.Printf("%#v\n", p)

}

func main() {
	encodeExample()
}
