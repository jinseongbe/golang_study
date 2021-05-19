package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func main() {
	var a, b, c, d, re int

	// 공백, 개행으로 나눠서 받음
	fmt.Println("자유롭게 숫자 네개 입력")
	re, _ = fmt.Scan(&a, &b, &c, &d)
	fmt.Printf("네수의 합은 : %d\n", a+b)
	fmt.Println("re :", re)
	// 여기서 n값은 입력 성공한 항목의 수 임!

	a, b = 0, 0

	// 사용자가 지정한 형식으로 입력받음
	fmt.Println("두수를 .으로 구분하여 입력(공백x)")
	re, _ = fmt.Scanf("%d.%d", &a, &b)
	fmt.Printf("두수의 합은 : %d\n", a+b)
	fmt.Println("re :", re)

	var aa, bb = 0, 0

	// 공백을 통해 입력받음
	fmt.Println("두수를 공백을 통해 입력")
	re, _ = fmt.Scanln(&aa, &bb)
	fmt.Printf("두수의 합은 : %d\n", aa+bb)
	fmt.Println("re :", re)

	// fmt.Fscan : bufio를 사용하여 빠른 입출력!
	Big := make([]int, 100000, 100000)
	fmt.Println("start!!")
	start := time.Now()
	for i := 0; i < 100000; i = i + 1 {
		fmt.Println(i)
		fmt.Scan(&Big[i])
	}
	fmt.Printf("fmt.Scan: %v\n", time.Since(start))

	start = time.Now()
	r := bufio.NewReader(os.Stdin)
	for i := 0; i < 100000; i = i + 1 {
		fmt.Fscan(r, &Big[i])
	}
	fmt.Printf("fmt.Fscan: %v\n", time.Since(start))
	fmt.Println("end!!")
}
