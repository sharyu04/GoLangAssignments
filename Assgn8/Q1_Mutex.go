package main

import (
	"fmt"
	"sync"
	"time"
)

func isEven(n int) bool {
	return n%2 == 0
}

func main() {
	n := 0

	var mu sync.Mutex

	go func() {
		mu.Lock()
		nIsEven := isEven(n)
		time.Sleep(5 * time.Millisecond)
		if nIsEven {
			fmt.Println(n, " is even")
			return
		}
		fmt.Println(n, "is odd")
		mu.Unlock()
	}()

	go func() {
		mu.Lock()
		n++
		mu.Unlock()
	}()

	time.Sleep(time.Second)
}
