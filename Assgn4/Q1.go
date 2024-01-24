package main

import "fmt"

type hello struct{
	name string
}

func main(){
	var inp int
	fmt.Scanln(&inp)

	hel := hello{"Hello"}

	switch inp{
	case 1:
		AcceptAnything(1)
	case 2:
		AcceptAnything("string")
	case 3:
		AcceptAnything(true)
	case 4:
		AcceptAnything(hel)
	default:
		fmt.Printf("Invalid option")
	}

}

func AcceptAnything(i interface{}){
	switch value := i.(type){
	case int:
		fmt.Printf("This value is of type Integer, %d", value)
	case string:
		fmt.Printf("This value is of type String, %s", value)
	case bool:
		fmt.Printf("This value is of type Boolean, %t", value)
	case hello:
		fmt.Printf("This value is of type Hello, %v", value)
	}
}