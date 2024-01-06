package main

import (
	"fmt"
	"math"
)

func Square(n float64) float64{
	s:= math.Pow(n, 2)
	return s
}

func main() {
	var n1 float64
	qs:=`Program to calculate Square of a given number
--------------------------------------------`
	fmt.Println(qs)
	fmt.Print("Enter number: ")
	fmt.Scan(&n1)

	fmt.Println(n1, "^ 2 = ", Square(n1))
}
