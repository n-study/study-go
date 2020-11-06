package main

import "fmt"

func main() {
	/* 	var n int
	   	fmt.Scan(&n)
	   	fmt.Println(n)

	   	// ==, !=, >, <, >=, <=
	   	if n > 0 {
	   		fmt.Println("양수입니다.")
	   	} else if n < 0 {
	   		fmt.Println("음수입니다.")
	   	} else {
	   		fmt.Println("0입니다.")
			 } */

	/*
		var s string
		fmt.Scan(&s)
		fmt.Println(s)

		// strcmp
		if s == "go" {
			fmt.Println("same")
		} */

	/* 	var s string
	   	fmt.Scan(&s)

	   	switch s {
	   	case "go":
	   		fmt.Println("Go!!!!")
	   		// break
	   		// fallthrough
	   	case "js":
	   		fmt.Println("Javascript!!!!")
	   	} */

	/* 	var n int
	   	fmt.Scan(&n)
	   	switch n {
	   	case 1, 3, 5, 7, 8, 10, 12:
	   		fmt.Println(31)
	   	case 2:
	   		fmt.Println(29)
	   	default:
	   		fmt.Println(30)
			 } */
Outer:
	for i := 2; i <= 9; i++ {
		for j := 2; j <= 9; j++ {
			if j == 5 {
				continue Outer
			}
			fmt.Printf("%d * %d = %d\t", i, j, i*j)
		}
		fmt.Println()
	}

	/* 	var n int

	   	for {
	   		fmt.Scan(&n)

	   		fmt.Println("입력값:", n)

	   		if n == 0 {
	   			break
	   		}
	   	} */
}
