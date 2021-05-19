// Channel : queue의 일종, Thread safe, Fixed size

package main

import (
	"fmt"
)

func main() {
	// var ch1 chan int
	// ch1 = make(chan int, 1)

	// ch1 := make(chan int, 1) // make(chan type, size),  fixedSize!, size가 명시되어 있지 않으면 0개짜리 채녈임!
	ch1 := make(chan int, 0) // fatal error: all goroutines are asleep - deadlock!
	// 0개 짜리 채널이란? : 다른 쓰레드에서 넣은것을 빼주지 않으면 끝나지 않음!

	go pop(ch1)
	ch1 <- 10

	fmt.Println("End!")
}

func pop(c chan int) {
	fmt.Println("PopFunc")
	v := <-c
	fmt.Println(v)
}
