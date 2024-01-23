package main

import (
	"fmt"
	"strings"
	"bufio"
	"os"
)

func main(){

	reader := bufio.NewReader(os.Stdin)
    fmt.Printf("Enter a string input: ")
    str, _ := reader.ReadString('\n')

	ans, freq := highestFreq(str)

	fmt.Println(ans, freq)

}

func highestFreq(str string) ([]string, int){
	slc := strings.Split(str, " ")
	mp := map[string]int{}
	for _,word := range slc{
		mp[word]++
	}

	var maxFreq int
	var ans []string
	// var ans = []string{}
	for word,freq := range mp{
		if freq > maxFreq{
			maxFreq = freq
			ans = []string{word}
		}else if freq == maxFreq{
			ans = append(ans,word)
		}
	}

	return ans, maxFreq

}