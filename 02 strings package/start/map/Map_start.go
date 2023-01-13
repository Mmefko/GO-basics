package main

import (
	"fmt"
	"strings"
)

func main() {
	shift := 2
	s := "The quick brown fox jumps over the lazy dog"

	//Create the mapping function
	transform := func(r rune) rune {
		switch {
		case r >= 'A' && r <= 'Z':
			value := int('A') + (int(r) - int('A') + shift)
			if value > 91 {
				value -= 26
			} else if value < 65 {
				value += 26
			}
			return rune(value)
		case r >= 'a' && r <= 'z':
			value := int('a') + (int(r) - int('a') + shift)
			if value > 122 {
				value -= 26
			} else if value < 97 {
				value += 26
			}
			return rune(value)
		}
		return r
	}

	//Encode the message
	encode := strings.Map(transform, s)
	fmt.Println(encode) //Vjg swkem dtqyp hqz lworu qxgt vjg ncba fqi

	//Decode the message
	shift = -2
	decode := strings.Map(transform, encode)
	fmt.Println(decode) //The quick brown fox jumps over the lazy dog
}
