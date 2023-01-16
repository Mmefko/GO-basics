package main

import (
	"fmt"
	"math"
)

func main() {
	x := 10.0

	//Absolute value
	fmt.Println(math.Abs(x), math.Abs(-x)) //10 10

	//Pow(x^y) and exp(e^x)
	fmt.Println(math.Pow(x, 2.0))    //100
	fmt.Println(math.E, math.Exp(x)) //2.718281828459045 22026.465794806718

	//Trigonometry and othr functions
	fmt.Println(math.Cos(math.Pi))     //-1
	fmt.Println(math.Sin(math.Pi / 2)) //1
	fmt.Println(math.Tan(math.Pi))     //-1.2246467991473515e-16
	fmt.Println(math.Log(10))          //2.302585092994046

	//Square and cube roots
	fmt.Println(math.Sqrt(25))  //5
	fmt.Println(math.Cbrt(125)) //5

	//Calculate the Hypothenuse of a right triangle
	fmt.Println(math.Hypot(30, 40)) //50
}
