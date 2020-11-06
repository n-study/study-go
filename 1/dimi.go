package main

import "fmt"

func main() {
	fmt.Println("hello world")

	var a int
	fmt.Println(a)
	a = 10
	fmt.Println(a)
	a = 50
	fmt.Println(a)
	a = 40
	fmt.Printf("%d", a)
	fmt.Println(a)

	a := 14.3
	fmt.Println(a)
	a = 5
	fmt.Println(a)

	s := "hello"
	fmt.Println(s)
	s = "go"
	// fmt.Println(s)

	s := "hello"
	fmt.Println(s)
	fmt.Println(*(&s))

	n := 5.3
	p := &n
	fmt.Println(p, *p)
}
