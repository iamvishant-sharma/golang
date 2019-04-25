package main

import "fmt"

func main(){
	a := make([]string, 3)
	fmt.Println("emp", a)

	a[0] = "a"
	a[1] = "b"
	a[2] = "c"

	fmt.Println("name", a)
	fmt.Println("name", a[2])

	fmt.Println("length:", len(a))

	a = append(a, "x")
	a = append(a, "y", "z")
	fmt.Println("append", a)

	b := make([]string, len(a))
	copy(b, a)
	fmt.Println("copy:", b)

	c := a[2:5]
	fmt.Println("middle order:", c)

	d := a[3:]
	fmt.Println("after value:", d)

	e := b[:3]
	fmt.Println("before value:", e)
}