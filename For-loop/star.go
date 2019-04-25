package main

import "fmt"

func main(){
	n := 1
	for i := 1 ; i <= 5 ; i++ {
		fmt.Println("")
		for j := 1 ; j <= n ; j++ {
			fmt.Printf("*")
		} 
		n++
	}
}