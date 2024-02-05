package main

import (
	"fmt"
)

func main() {

	chAlice := make(chan string)
	chBob := make(chan string)
	chSignal := make(chan struct{})

	go func() {

		inp := "helloBob$helloalice#howareyou?#Iamgood.howareyou?$^"

		str := ""
		for _, letter := range inp {
			if letter == '$' {
				chAlice <- str
				str = ""
			} else if letter == '#' {
				chBob <- str
				str = ""
			} else if letter == '^' {
				chSignal <- struct{}{}
			} else {
				str = str + string(letter)
			}
		}
	}()

	ans := ""

	for {
		select {
		case test := <-chAlice:
			ans = ans + "alice : " + test + ","
		case test := <-chBob:
			ans = ans + "bob : " + test + ","
		case <-chSignal:
			fmt.Println(ans)
			return
		}

	}

}
