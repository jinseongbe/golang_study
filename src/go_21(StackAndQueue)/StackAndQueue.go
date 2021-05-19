// Stack = FILO(first input last out)
// Queue = FIFO(first input first out)

package main

import (
	"dataStruct"
	"fmt"
)

func main() {

	// stack by slice
	fmt.Println("stack by slice")
	stack := []int{}

	for i := 1; i <= 5; i++ {
		stack = append(stack, i)
	}
	fmt.Println(stack)

	for len(stack) > 0 {
		var last int
		last, stack = stack[len(stack)-1], stack[:len(stack)-1]
		fmt.Println(stack, "->", last)
	}

	fmt.Println()

	// queue by slice
	fmt.Println("queue by slice")
	queue := []int{}
	for i := 1; i <= 5; i++ {
		queue = append(queue, i)
	}
	fmt.Println(queue)

	for len(queue) > 0 {
		var front int
		front, queue = queue[0], queue[1:]
		fmt.Println(queue, "->", front)
	}

	fmt.Println()

	// stack by linkedlist(package dataStruct)
	fmt.Println("stack by linkedlist")
	stack2 := dataStruct.NewStack()
	for i := 1; i <= 5; i++ {
		stack2.Push(i)
	}
	stack2.PrintStack()

	for !stack2.Empty() {
		val := stack2.Pop()

		fmt.Printf("[%d] ", val)
		stack2.PrintStack()
	}
	fmt.Println()

	fmt.Println()

	// queue by linkedlist(package dataStruct)
	fmt.Println("queue by linkedlist")
	queue2 := dataStruct.NewQueue()
	for i := 1; i <= 5; i++ {
		queue2.Push(i)
	}
	queue2.PrintQueue()

	for !queue2.Empty() {
		val := queue2.Pop()

		fmt.Printf("[%d] ", val)
		queue2.PrintQueue()
	}
	fmt.Println()

}
