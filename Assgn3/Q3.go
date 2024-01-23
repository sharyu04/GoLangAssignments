package main

import "fmt"

func main(){
	var idx1 int
	var idx2 int
	fmt.Printf("Enter two integers : ")
	fmt.Scanln(&idx1, &idx2)

	var arr = [8]string{"qwe", "wer", "ert", "rty", "tyu", "yui", "uio", "iop"}
	
	if idx1 > idx2 || idx1<0 || idx2>7{
		fmt.Printf("Invalid Indexes")
		return
	}

	var arr1 []string
	var arr2 []string
	var arr3 []string

	if idx1+1<=7{
		arr1 = arr[:idx1+1]
	}else{
		arr1 = arr[:]
	}

	if idx2+1>=7{
		arr2 = arr[idx1:]
	}else{
		arr2 = arr[idx1:idx2+1]
	}

	arr3 = arr[idx2:]

	fmt.Println(arr1)
	fmt.Println(arr2)
	fmt.Println(arr3)
	

}
