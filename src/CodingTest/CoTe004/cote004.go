package main

import (
	"fmt"
)

func main() {
	var a, b, c, d, re int

	// 공백, 개행으로 나눠서 받음
	fmt.Println("자유롭게 숫자 네개 입력")
	re, _ = fmt.Scan(&a, &b, &c, &d)
	fmt.Printf("네수의 합은 : %d\n", a+b)
	fmt.Println("re :", re)

	// var n int
	// re, _ = fmt.Scan(n)

}
