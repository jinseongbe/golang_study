package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan int)

	go push(ch1)

	timerChane := time.After(20 * time.Second)
	tickTimerChan := time.Tick(2 * time.Second)

	for {
		select {
		case v := <-ch1:
			fmt.Println(v)
		case <-timerChane:
			fmt.Println("timeout")
			return
		case <-tickTimerChan:
			fmt.Println("Tick")

			// default:
			// 	fmt.Println("Idle")
			// 	time.Sleep(1 * time.Second)
		}
	}
}

func push(c chan int) {
	i := 0
	for {
		time.Sleep(1 * time.Second)
		c <- i
		i++
	}
}
