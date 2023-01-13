package main

import (
	"fmt"
)

type circle struct {
	radius int
	border int
}

func main() {
	x := 20
	f := 123.45

	//Basic formating
	fmt.Printf("%d\n", x)
	fmt.Printf("%x\n", x)

	//Booleans can be printed as "true" or "false"
	fmt.Printf("%t\n", x > 10)

	//Floating point
	fmt.Printf("%f\n", f)
	fmt.Printf("%e\n", f)

	//Using explicit argument indexes
	fmt.Printf("%d %d\n", 50, 40)       //50 40
	fmt.Printf("%[2]d %[1]d\n", 50, 40) //40 50

	//Argument indexes can be used to print values repeatedly
	fmt.Printf("%d %#[1]o %[1]x\n", 50) //50 062 32

	//Print a value in dafault format
	c := circle{
		radius: 20,
		border: 5,
	}
	fmt.Printf("%v\n", c)  //{20 5}
	fmt.Printf("%+v\n", c) //{radius:20 border:5}

	//Sprintf is the same as Printf, but returns a string
	str := fmt.Sprintf("%+v\n", c)
	fmt.Println(str) //{radius:20 border:5}
}
