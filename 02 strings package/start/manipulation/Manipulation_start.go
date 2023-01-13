package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	s := "The quick brown fox jumps over the lazy dog"
	s2 := []string{"one", "two", "three", "four"}
	s3 := "This is a string. With some punctuation, for a demo! Yep."

	//The Split function splits a string into substrings
	sub1 := strings.Split(s, " ")
	fmt.Printf("%q\n", sub1[2]) //brown
	sub2 := strings.Split(s, "the")
	fmt.Printf("%q\n", sub2) //["The quick brown fox jumps over " " lazy dog"]

	// Join concatenates substring, with the separator between each
	fmt.Println(strings.Join(s2, "-")) //one-two-three-four\

	//The fields function splits a string around white space
	result := strings.Fields(s)
	fmt.Printf("%q\n", result) //["The" "quick" "brown" "fox" "jumps" "over" "the" "lazy" "dog"]

	//FieldFunc is a customizable version of field that uses a callback
	f := func(c rune) bool {
		return unicode.IsPunct(c)
	}

	result2 := strings.FieldsFunc(s3, f)
	fmt.Printf("%q\n", result2) //["This is a string" " With some punctuation" " for a demo" " Yep"]

	//Replacer can be used for multiple replacement operations
	rep := strings.NewReplacer(".", "|", ",", "|", "!", "|")
	result3 := rep.Replace(s3)
	fmt.Println(result3) //This is a string| With some punctuation| for a demo| Yep|
}
