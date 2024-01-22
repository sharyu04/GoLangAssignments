package main

import (
	"fmt"
	"math"
)

func main() {
	var radius float64
	var area float64
	radius = 2
	area = math.Pi * math.Pow(radius, 2)
	ans := fmt.Sprintf("%.2f",area)
	fmt.Println("Area = ",ans)
}