package main

import "fmt"
import "sort"

func initial(){
	var a [5]int
	fmt.Println(a)

	a[4] = 55
	fmt.Println(a)
}

func secondry(){
	a := [] int {1,6,3,4}
	sort.Ints(a)
	fmt.Println(a)
}

func tertiary(){
	a := [] string {"chiklu", "tinku", "ranga", "nath"}
	sort.Strings(a)
	fmt.Println("nominates names:", a)
	
}

func main(){
	tertiary()
}