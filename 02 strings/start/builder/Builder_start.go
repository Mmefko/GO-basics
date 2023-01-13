package main

import (
	"fmt"
	"strings"
)

func main() {
	//Declare any empty string
	var sb strings.Builder

	//Write some content
	sb.WriteString("some string 1 \n")
	sb.WriteString("some string 2 \n")
	sb.WriteString("some string 3 \n")

	//Output the concatenated string
	fmt.Println(sb.String())

	//Examine the builder's capacity
	fmt.Println("Capacity:", sb.Cap()) //64

	//Grow the capacity - use this if you know in advance how much data
	//you are going to be writting into the buffer to minimize copies
	sb.Grow(1024)
	fmt.Println("Capacity:", sb.Cap()) //1152

	for i := 0; i <= 10; i++ {
		fmt.Fprintf(&sb, "String %d -- ", i)
	}
	fmt.Println(sb.String()) //String 0 -- String 1 -- String 2 -- String 3 -- String 4 -- String 5 -- String 6 -- String 7 -- String 8 -- String 9 -- String 10 --

	//Length of the final string will be
	fmt.Println("builder size is", sb.Len()) //builder size is 178

	//Reset function will reset the builder to original state
	sb.Reset()
	fmt.Println("Capacity:", sb.Cap()) //Capacity: 0
}
