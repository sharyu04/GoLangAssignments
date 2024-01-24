package main

import (
	"bufio"
    "os"
	"fmt"
	"sync"
	"runtime"
)

func main(){

	reader := bufio.NewReader(os.Stdin)
    fmt.Printf("Enter a string input: ")
    inputStr, _ := reader.ReadString('\n')

	wg := new(sync.WaitGroup)

	wg.Add(1)
	go revStr(inputStr, wg)

	wg.Wait()
	
}

func revStr(str string, wg *sync.WaitGroup){
	defer wg.Done()
	revStr := ""
	for _,lt := range str{
		revStr = string(lt) + revStr
	}
	fmt.Println(revStr, runtime.NumGoroutine())
}