
package main

import "fmt"

func main() {
	principal := 1000
	ROI := 5
	Time := 3
	SI := (principal * ROI * Time)/100
	ans := fmt.Sprintf("%.2f",float64(SI))
	fmt.Println("Simple Interest = ", ans)
}
