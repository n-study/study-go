package main

import (
	"fmt"
	"studygo/example"
)

/*
func delay(ch chan bool) {
	time.Sleep(time.Second * 5)
	fmt.Println("end")
	ch <- true
}
*/

func main() {
	ch := make(chan bool)

	// go delay(ch)
	go example.Delay(ch)
	example.

		// go (func() {
		// 	time.Sleep(time.Second * 5)
		// 	fmt.Println("enkd")
		// 	ch <- true
		// })()

		fmt.Println("another")
	fmt.Println(<-ch)
}
