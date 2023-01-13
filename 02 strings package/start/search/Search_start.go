package main

import (
	"fmt"
	"strings"
)

func main() {
	fname := "filename.txt"
	fname2 := "temp_picture.jpeg"
	vowels := "aeiouAEIOU"
	s := "The quick brown fox jumps over the lazy dog"

	//Contains to see if a substring is in a string
	fmt.Println(strings.Contains(s, "jump")) //true

	fmt.Println(strings.ContainsAny(s, "abcd")) //true

	//Use index to find the offset of the first instance of a substring
	fmt.Println(strings.Index(s, "fox")) //16
	fmt.Println(strings.Index(s, "cat")) //-1

	fmt.Println(strings.IndexAny("golang", vowels)) //1 ---> g[o]lang
	fmt.Println(strings.IndexAny("gzrbl", vowels))  //-1

	//HasPreffix / Suffix
	fmt.Println(strings.HasSuffix(fname, "txt"))   //true
	fmt.Println(strings.HasPrefix(fname2, "temp")) //true

	//Count returns the number of non-overlapping instances of a substring
	fmt.Println(strings.Count(s, "the")) //1

}
