package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "The quick brown fox jumps over the lazy dog"

	//Length of string
	fmt.Println(len(s)) //43

	//Iterate iver each character
	for _, val := range s {
		fmt.Print(string(val), ",") //T,h,e, ,q,u,i,c,k, ,b,r,o,w,n, ,f,o,x, ,j,u,m,p,s, ,o,v,e,r, ,t,h,e, ,l,a,z,y, ,d,o,g,
	}
	fmt.Println()

	//Using poerations < > == !=
	fmt.Println("dog" < "cat")   //false
	fmt.Println("dog" < "horse") //true
	fmt.Println("dog" == "Dog")  //false

	//Comparing two strings
	result := strings.Compare("dog", "cat")
	fmt.Println(result) //1

	result = strings.Compare("dog", "dog")
	fmt.Println(result) //0

	//EqualFold tests using Unicode case-folding
	fmt.Println(strings.EqualFold("Hello", "hello")) //true

	// ToUpper, ToLower, Title
	s1 := strings.ToUpper(s)
	s2 := strings.ToLower(s)
	s3 := strings.Title(s)
	fmt.Println(s1) //THE QUICK BROWN FOX JUMPS OVER THE LAZY DOG
	fmt.Println(s2) //the quick brown fox jumps over the lazy dog
	fmt.Println(s3) //The Quick Brown Fox Jumps Over The Lazy Dog
}
