package concurrency

import (
	"fmt"
	"testing"
	"time"
)

func TestForSelect(t *testing.T) {
	tick := time.Tick(100 * time.Millisecond)
	timer := time.After(1 * time.Second)
	boom := time.After(1200 * time.Millisecond)
	for {
		select {
		case <-tick:
			fmt.Println("tick.")
		case <-timer:
			fmt.Println(". 1 second passed !!! ")
		case <-boom:
			fmt.Println("BOOM!")
			return // exit
		default:
			fmt.Println(".")
			time.Sleep(50 * time.Millisecond)
		}
	}
}
