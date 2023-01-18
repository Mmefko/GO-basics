package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func main() {
	//Seed the random number generator
	rand.Seed(time.Now().UnixNano())

	//Shuffle a sequence of values
	const numstring = "one two three four five six"
	words := strings.Fields(numstring)
	rand.Shuffle(len(words), func(i, j int) {
		words[i], words[j] = words[j], words[i]
	})

	fmt.Println(words) //[six five three four one two]

	//Generate a rendom string
	const upletters = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	const lowletters = "abcdefghijklmnopqrstuvwxyz"
	const digits = "0123456789"
	const specialchars = "~=+%^*()[]{}!@#$?|"
	const allchars = upletters + lowletters + digits + specialchars

	var sb strings.Builder
	length := 10

	for i := 0; i < length; i++ {
		sb.WriteRune(rune(allchars[rand.Intn(len(allchars))]))
	}

	fmt.Println("String result: ", sb.String()) //String result:  nL|7WML(1V

	//Generate a random string with rules
	buf := make([]byte, length)
	buf[0] = digits[rand.Intn(len(digits))]
	buf[1] = specialchars[rand.Intn(len(specialchars))]
	buf[2] = upletters[rand.Intn(len(upletters))]
	buf[3] = lowletters[rand.Intn(len(lowletters))]
	for i := 4; i < length; i++ {
		buf[i] = allchars[rand.Intn(len(allchars))]
	}

	rand.Shuffle(len(buf), func(i, j int) {
		buf[i], buf[j] = buf[j], buf[i]
	})

	fmt.Println(string(buf)) //p@5yUb!A|H
}
