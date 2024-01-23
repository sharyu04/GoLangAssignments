package main

import (
	"fmt"
	"math"
)

func main() {
	var radius float64
	fmt.Printf("Enter radius : ")
	fmt.Scanln(&radius)
	fmt.Println("Area = ",area(radius))
}

func area(radius float64)string{
	var area float64
	area = math.Pi * math.Pow(radius, 2)
	ans := fmt.Sprintf("%.2f",area)
	return ans
}