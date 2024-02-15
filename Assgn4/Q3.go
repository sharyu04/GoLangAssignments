package main

import "fmt"

type Rectangle struct{
	length int
	breadth int
}

type Square struct{
	side int
}

func (rec Rectangle) Area() int{
	return rec.length * rec.breadth
}

func (sq Square) Area() int{
	return sq.side * sq.side
}

func (rec Rectangle) Perimeter()int{
	return 2 * (rec.length + rec.breadth)
}

func (sq Square) Perimeter()int {
	return 4 * sq.side
}

type Quadrilateral interface{
	Area() int
	Perimeter() int
}

func PrintCustom(i Quadrilateral){
	fmt.Printf("\nArea: %d",i.Area())
	fmt.Printf("\nPerimeter: %d",i.Perimeter())
}

func main(){
	rec := Rectangle{10,20}
	sq := Square{10}

	var inp int
	fmt.Printf("Enter 1 for rectangle and 2 for square : ")
	fmt.Scanf("%d",&inp)

	switch inp{
	case 1:
		PrintCustom(rec)
	case 2:
		PrintCustom(sq)
	default:
		fmt.Printf("Invalid input")
	}
}