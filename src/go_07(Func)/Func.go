package main

import "fmt"

func add(x int, y int) int {
	return x + y
}

func fun1(x, y int) (int, int) {
	fun2(x, y)
	return y, x
}

func fun2(x, y int) {
	fmt.Println("func2 :", x, y)
}

func f1(x int) { // recursive function
	if x == 0 {
		return
	}
	fmt.Printf("f1(%d) before call f1(%d)\n", x, x-1)
	f1(x - 1)
	fmt.Printf("f1(%d) after call f1(%d)\n", x, x-1)
}

func sum(x int, s int) int {
	if x == 0 {
		return s
	}

	s += x
	return sum(x-1, s)
}

func fibo(x int) int {
	if x == 0 {
		return 1
	}

	if x == 1 {
		return 1
	}

	return fibo(x-1) + fibo(x-2)
}

func main() {

	for i := 0; i < 10; i++ {
		fmt.Printf("%d + %d = %d\n", i, i+2, add(i, i+2))
	}

	fmt.Println()

	a, b := fun1(2, 4)
	fmt.Println(a, b)

	fmt.Println()

	f1(10)

	fmt.Println()

	fmt.Println(sum(10, 0))
	// 모든 재귀함수는 for문으로 바꿀 수 있다, (피보나치수열(수학정의) 과 같은 경우에는 재귀호출이 더 편리하다!)
	// 그러나 재귀호출의 문재는 함수의 호출과정이 반복문보다 단계가 많다(즉 느리다!)
	s := 0
	for i := 1; i <= 10; i++ {
		s += i
		if i == 10 {
			fmt.Println(s)
		}
	}

	fmt.Println(fibo(6))

}
