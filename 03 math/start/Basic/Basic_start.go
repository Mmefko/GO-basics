package main

import (
	"fmt"
	"math"
)

func main() {
	//Print PI
	fmt.Println(math.Pi) //3.141592653589793

	// Ceiling and Floor Functions
	fmt.Println(math.Ceil(math.Pi))  //4
	fmt.Println(math.Floor(math.Pi)) //3

	//Truncate will return int value if X
	fmt.Printf("%.2f\n", math.Trunc(math.Pi)) //3.00

	//Min and Max
	fmt.Println(math.Max(math.Pi, math.Ln2)) //3.141592653589793
	fmt.Println(math.Min(math.Pi, math.Ln2)) //0.6931471805599453

	//Mod operator is like % but for float
	fmt.Println(17 % 5)              //2
	fmt.Println(math.Mod(17.0, 5.0)) //2

	//Round and RoundToEven
	fmt.Printf("%.1f\n", math.Round(10.5))  //11.0
	fmt.Printf("%.1f\n", math.Round(-10.5)) //-11.0

	fmt.Printf("%.1f\n", math.RoundToEven(10.5)) //10.0
	fmt.Printf("%.1f\n", math.RoundToEven(11.5)) //12.0
}
