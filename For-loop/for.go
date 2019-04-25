package main

import "fmt"

func main(){
	i := 1
	for i <= 10 {
		fmt.Println(i)
		i = i+2
	}

	for j := 5 ; j <= 12 ; j++{
		fmt.Println(j)
	}

	for {
		fmt.Println("loop")
		break
	}

	for n := 20 ; n <= 25 ; n++{
		if (n)%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
}