package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	a := make([]int, 8, 12)

	for i := 0; i < 8; i++ {
		a[i] = rand.Intn(8)
	}

	fmt.Printf("%v : %v,\t %v address : %p,\t capacity : %v\n", "a", a, "a", a, cap(a))
	fmt.Println()

	b := a[4:6]
	c := a[3:]
	d := a[:4]
	fmt.Printf("%v : %v,\t\t %v address : %p,\t capacity : %v\n", "b", b, "b", b, cap(b))
	fmt.Printf("%v : %v,\t %v address : %p,\t capacity : %v\n", "c", c, "c", c, cap(c))
	fmt.Printf("%v : %v,\t\t %v address : %p,\t capacity : %v\n", "d", d, "d", d, cap(d))

	fmt.Println()

	for i := 0; i < 3; i++ {
		var back int
		a, back = RemoveBack(a)
		fmt.Println(back)
	}
	fmt.Printf("%v : %v,\t %v address : %p,\t capacity : %v\n", "a", a, "a", a, cap(a))
	fmt.Println()

	for i := 0; i < 3; i++ {
		var front int
		a, front = RemoveFront(a)
		fmt.Println(front)
	}
	fmt.Printf("%v : %v,\t %v address : %p,\t capacity : %v\n", "a", a, "a", a, cap(a))

}

func RemoveBack(a []int) ([]int, int) {
	return a[:len(a)-1], a[len(a)-1]
}

func RemoveFront(a []int) ([]int, int) {
	return a[1:], a[0]
}
