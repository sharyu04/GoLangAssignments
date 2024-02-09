package main

import (
	"fmt"
	"time"
)

func isEven(n int) bool {
	return n%2 == 0
}

func main() {
	n := 0

	ch := make(chan struct{})

	go func() {
		nIsEven := isEven(n)
		time.Sleep(5 * time.Millisecond)
		if nIsEven {
			fmt.Println(n, " is even")
			return
		}
		fmt.Println(n, "is odd")
		ch <- struct{}{}
	}()

	go func() {
		select {
		case <-ch:
			break
		}
		n++
	}()

	time.Sleep(time.Second)
}
