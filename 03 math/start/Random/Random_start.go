package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	//Initialize the random seed to an unknown value
	rand.Seed(time.Now().UnixNano())

	//Generate random integer numbers
	fmt.Println(rand.Int())    //4860731531654632520
	fmt.Println(rand.Intn(10)) //2

	//Generate random floating Point numbers
	fmt.Println(rand.Float32()) //0.12687156
	fmt.Println(rand.Float64()) //0.4985925691466604

	//use the Perm function to create permutations
	s := []string{"apple", "pear", "grape", "orange", "kiwi", "melon"}
	indexes := rand.Perm(len(s))
	fmt.Println(indexes) // 0 5 4 3 1 2]

	for i := 0; i < len(s); i++ {
		fmt.Printf("%v - ", s[indexes[i]]) //apple - melon - kiwi - orange - pear - grape
	}
	fmt.Println()

	//generate a random integer between a and b
	const a, b = 10, 50
	for i := 0; i < 10; i++ {
		n := a + rand.Intn(b-a+1)
		fmt.Print(" ", n) // 15 36 46 46 41 24 16 47 47
	}
	fmt.Println()

	//Generate a random uppercase character
	for i := 0; i < 10; i++ {
		c := string('A' + rune(rand.Intn('Z'-'A'+1)))
		fmt.Print(" ", c) //  U S P Q Z W K U V A
	}
	fmt.Println()
}
