package main

import "fmt"

func main(){
	var index int
	fmt.Print("Enter the index : ")
	fmt.Scanln(&index)
	day := checkTheDay(index)
	if day==""{
		fmt.Printf("Not a day")
	}else{
		fmt.Printf("%s",day)
	}
}

func checkTheDay(index int) string{
	var days map[int]string = map[int]string{1:"Monday", 2:"Tuesday", 3:"Wednesday", 4:"Thursday", 5:"Friday", 6:"Saturday", 7:"Sunday"}
	return days[index]
}