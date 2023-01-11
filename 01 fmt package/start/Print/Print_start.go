package main

import (
	"fmt"
)

func main() {
	fmt.Print("Welcome to GO! ")
	fmt.Println("This string ends with a new line!")

	const answer = 42
	fmt.Println("The answer to life is", answer)

	const a, b, c = 1, 2, 3
	fmt.Println("And:", a, "and:", b, "and:", c)

	items := []int{1, 2, 3, 4}
	length, err := fmt.Println(items)
	fmt.Println(length, err)
}
