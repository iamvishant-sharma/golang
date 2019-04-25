package main

// import "strings"
import "os"
import "fmt"

func lignite(a string, b string){
	// var a = "vishant"
	// b := "sharma"
	fmt.Println("our hero is " + (a) + " " + (b))
}

func metaphor(){
	c := "chkilu"
	fmt.Println("My name is " + (c))
}

func main(){
	lignite(os.Args[1], os.Args[2])
}
