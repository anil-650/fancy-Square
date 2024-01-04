package main

import (
	"fmt"
	"math"
)

func square(n int) int{
	s:= math.Round(math.Pow(float64(n),2))
	return int(s)
}

func main() {
	var n1 int
	qs:=`Program to calculate Square of a given number
--------------------------------------------`
	fmt.Println(qs)
	fmt.Print("Enter number: ")
	fmt.Scan(&n1)

	fmt.Println(n1, "^ 2 = ", square(n1))
}
