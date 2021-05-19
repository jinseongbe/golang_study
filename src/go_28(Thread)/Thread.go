package main

import (
	"fmt"
	"time"
)

func main() {

	go fun1()
	go fun2()
	go fun3()

	for i := 0; i < 20; i++ {
		fmt.Println("main", i)
	}
	fmt.Scanln()

	fmt.Println()
}

func fun1() {
	for i := 0; i < 10; i++ {
		time.Sleep(2 * time.Millisecond)
		fmt.Println("fun1:", i)
	}
}

func fun2() {
	for i := 0; i < 10; i++ {
		time.Sleep(3 * time.Millisecond)
		fmt.Println("fun2:", i)
	}
}

func fun3() {
	for i := 0; i < 10; i++ {
		time.Sleep(1 * time.Millisecond)
		fmt.Println("fun3:", i)
	}
}
