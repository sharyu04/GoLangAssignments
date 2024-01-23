
package main

import "fmt"

func main() {
	var principal float64
	var ROI float64
	var Time float64
	fmt.Printf("Enter the principal amount : ")
	fmt.Scanln(&principal)
	fmt.Printf("Enter the rate of interest : ")
	fmt.Scanln(&ROI)
	fmt.Printf("Enter the Time period in years : ")
	fmt.Scanln(&Time)
	fmt.Println("Simple Interest = ",SiCalculator(principal,ROI,Time))
}

func SiCalculator(principal float64, ROI float64, Time float64) string{
	SI := (principal * ROI * Time)/100
	var ans string
	ans = fmt.Sprintf("%.2f",float64(SI))
	return ans
}