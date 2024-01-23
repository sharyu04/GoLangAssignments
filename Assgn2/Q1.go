package main

import "fmt"

func main() {
	var str string
	fmt.Print("Enter the Roman Numerical value : ")
	fmt.Scanln(&str)
	// str = "MCMXCIV"
	
	fmt.Println("Number in Integer : ",convertToInt(str))
}

func convertToInt(str string)int{
	var sum int
	sum = 0
	for i:=0; i<len(str); i++{
		if str[i]=='I'{
			if i+1 < len(str) && str[i+1]=='V'{
				sum = sum + 4
				i++
			}else if i+1 < len(str) && str[i+1]=='X'{
				sum = sum + 9
				i++
			}else{
				sum = sum + 1
			}
		}else if str[i]=='V'{
			sum = sum + 5
		}else if str[i]=='X'{
			if i+1 < len(str) && str[i+1]=='L'{
				sum = sum + 40
				i++
			}else if i+1 < len(str) && str[i+1]=='C'{
				sum = sum + 90
				i++
			}else{
				sum = sum + 10
			}
		}else if str[i]=='L'{
			sum = sum + 50
		}else if str[i]=='C'{
			if i+1 < len(str) && str[i+1]=='D'{
				sum = sum + 400
				i++
			}else if str[i+1]=='M'{
				sum = sum + 900
				i++
			}else{
				sum = sum + 100
			}
		}else if str[i]=='D'{
			sum = sum + 500
		}else if str[i]=='M'{
			sum = sum + 1000
		}
	}
	return sum
}