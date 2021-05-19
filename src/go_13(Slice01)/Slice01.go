// 동적배열, (c++의 vector와 같은 개념)
// golang에서는 동적배열의 공간을 늘릴때 2배수로 늘린다, 1-2-4-8

package main

import (
	"fmt"
)

func main() {
	var a1 []int
	a2 := []int{1, 2, 3, 4, 5}
	a3 := make([]int, 3)
	a4 := make([]int, 4, 8) // make([]type, length, capacity), capacity는 확보한 공간 length는 실제로 사용한 길이

	fmt.Println(a1)
	fmt.Println()
	fmt.Println(a2)
	fmt.Println()
	fmt.Println(a3)
	fmt.Println(len(a3))
	fmt.Println(cap(a3))
	fmt.Println()
	fmt.Println(a4)
	fmt.Println(len(a4))
	fmt.Println(cap(a4))
	fmt.Println()

	b := make([]int, 0, 2)
	for i := 0; i < 10; i++ {
		b = append(b, i)
		var Pb *int = &b[0]
		fmt.Println("b =", b)
		fmt.Println("len(b) =", len(b), "cap(b) =", cap(b), "addresss = ", Pb)
		fmt.Println()
	}
	fmt.Println()

	// 굉장히 중요함!!(다른 변수에 append하는 경우)
	c1 := make([]int, 2, 2)
	c2 := append(c1, 1)
	fmt.Printf("%p, %p\n", c1, c2)
	fmt.Println(c1, c2, "\n")

	d1 := make([]int, 2, 4) // 같은 주소를 공유하고 있음!! 주의해야함!
	d2 := append(d1, 1)
	fmt.Printf("%p, %p\n", d1, d2)
	fmt.Println(d1, d2)
	d2[0], d2[2] = d2[2], d2[0] // d2만 바꿨지만 d1도 같이 바뀜!!(메모리주소를 공유하기 때문!!) 주의해야함!!!!
	fmt.Println(d1, d2, "\n")

	// 이런 실수를 하기 싫다면, 복사해서 사용하는 것이 안전하다!
	e1 := make([]int, 2, 4)
	e1[0] = 1
	e1[1] = 2

	e2 := make([]int, len(e1))
	for i := 0; i < len(e1); i++ {
		e2[i] = e1[i]
	}
	e2 = append(e2, 3)
	e2[0] = 9
	fmt.Printf("%p, %p\n", e1, e2)
	fmt.Println(e1, e2)

}
