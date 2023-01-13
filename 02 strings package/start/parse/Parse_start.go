package main

import (
	"fmt"
	"log"
	"strconv"
)

func main() {
	sampleint := 100
	samplestr := "250"

	//Convert an int to string
	newstr := string(sampleint)
	fmt.Printf("result: %[1]v - Type: %[1]T\n", newstr) //result: d - Type: string
	newstr2 := strconv.Itoa(sampleint)
	fmt.Printf("result: %[1]v - Type: %[1]T\n", newstr2) //result: 100 - Type: string

	//Convert a string to int
	sampleint, err := strconv.Atoi(samplestr)
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Printf("result: %[1]v - Type: %[1]T\n", sampleint) //result: 250 - Type: int

	//Other parse functions
	b, _ := strconv.ParseBool("true")
	fmt.Printf("result: %[1]v - Type: %[1]T\n", b) //result: true - Type: bool
	f, _ := strconv.ParseFloat("3.1415", 64)
	fmt.Printf("result: %[1]v - Type: %[1]T\n", f) //result: 3.1415 - Type: float64
	i, _ := strconv.ParseInt("-42", 10, 64)
	fmt.Printf("result: %[1]v - Type: %[1]T\n", i) //result: -42 - Type: int64
	ui, _ := strconv.ParseUint("42", 10, 64)
	fmt.Printf("result: %[1]v - Type: %[1]T\n", ui) //result: 42 - Type: uint64

	//Other format functions
	s := strconv.FormatBool(true)
	fmt.Printf("result: %[1]v - Type: %[1]T\n", s) //result: true - Type: string
	s = strconv.FormatFloat(3.1415, 'E', -1, 64)
	fmt.Printf("result: %[1]v - Type: %[1]T\n", s) //result: 3.1415E+00 - Type: string
	s = strconv.FormatInt(-42, 10)
	fmt.Printf("result: %[1]v - Type: %[1]T\n", s) //result: -42 - Type: string
	s = strconv.FormatUint(42, 10)
	fmt.Printf("result: %[1]v - Type: %[1]T\n", s) //result: 42 - Type: string

}
