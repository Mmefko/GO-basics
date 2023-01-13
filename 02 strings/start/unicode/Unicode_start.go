package main

import (
	"fmt"
	"unicode"
)

func main() {
	const s = "The 'quick' brown fox, jumped over the *LAZY* dog!"
	punctCount := 0
	lowerCount, upperCount := 0, 0
	spaceCount := 0
	hexDigitCount := 0

	// iterate over the characters
	for _, ch := range s {
		if unicode.IsPunct(ch) {
			punctCount++
		}
		if unicode.IsLower(ch) {
			lowerCount++
		}
		if unicode.IsUpper(ch) {
			upperCount++
		}
		if unicode.IsSpace(ch) {
			spaceCount++
		}
		if unicode.Is(unicode.Hex_Digit, ch) {
			hexDigitCount++
		}
	}

	//Print the result

	fmt.Printf("punctCount = %d\n", punctCount)       //punctCount = 6
	fmt.Printf("lowerCount = %d\n", lowerCount)       //lowerCount = 31
	fmt.Printf("upperCount = %d\n", upperCount)       //upperCount = 5
	fmt.Printf("spaceCount  = %d\n", spaceCount)      //spaceCount  = 8
	fmt.Printf("hexDigitCount = %d\n", hexDigitCount) //hexDigitCount = 10
}
