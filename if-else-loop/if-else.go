package main

import "fmt"

func main() {
	n := 5
	if n%2 == 0 {
		fmt.Println(n,"is even number")
	} else {
		fmt.Println(n,"is odd number")
	}

	if n < 0 {
		fmt.Println(n,"is negative number")
	} else if n < 10 {
		fmt.Println(n,"is 1 digit number")
	} else {
		fmt.Println(n,"is multiple digit number")
	}
}