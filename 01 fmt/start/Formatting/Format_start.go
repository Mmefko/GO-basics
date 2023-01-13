package main

import (
	"fmt"
)

func main() {
	f := 123.4567

	//Control the decimal precision
	fmt.Printf("%.2f\n", f) //123.46

	//Print with width 10 and default precision
	fmt.Printf("%10f\n", f) //123.4567000

	//Print with padding and precision
	fmt.Printf("%10.2f\n", f) //    123.46

	//Always use a + sign
	fmt.Printf("%+10.2f\n", f) //   +123.46

	//Pad with 0s instead of spaces
	fmt.Printf("%010.2f\n", f) //0000123.46

}
