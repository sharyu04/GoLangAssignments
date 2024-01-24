package main

import "fmt"

func main(){
	var size int
	fmt.Printf("\nEnter the size of the slice: ")
	fmt.Scanln(&size)
	slc := make([]int,size)
	fmt.Printf("\nEnter slice elements: ")
	for i := 0; i<size; i++{
		fmt.Scanf("%d",&slc[i])
	}
	
	accessSlice(slc)

	fmt.Printf("\nTesting Panic and Recover")

}

func accessSlice(slc []int){

	defer func(){
		if r:=recover(); r!=nil{
			fmt.Printf("\nInternal error: %v",r)
		}
	}()

	var idx int
	fmt.Printf("\nEnter index value: ")
	fmt.Scanln(&idx)
	

	fmt.Printf("\nitem %d, value %d",idx,slc[idx])
}