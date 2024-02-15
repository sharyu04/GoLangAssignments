package main

import "fmt"

type Rectangle struct{
	Length int
	Breadth int
}

func (rec *Rectangle) Area() int{
	return rec.Length * rec.Breadth
} 

func (rec *Rectangle) Perimeter() int{
	return 2 * (rec.Length + rec.Breadth)
}

func main(){
	var len int
	var bre int

	fmt.Printf("Enter length of the rectangle : ")
	fmt.Scanln(&len)
	fmt.Printf("Enter breadth of the rectangle : ")
	fmt.Scanln(&bre)

	rec := Rectangle{len,bre}

	fmt.Printf("\nArea : %d", rec.Area())
	fmt.Printf("\nPerimeter : %d", rec.Perimeter())
}