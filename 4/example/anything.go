package example

import (
	"fmt"
	"time"
)

// Delay 5 seconds, print end and pass true
func Delay(ch chan bool) {
	time.Sleep(time.Second * 5)
	fmt.Println("end")
	ch <- true
}
