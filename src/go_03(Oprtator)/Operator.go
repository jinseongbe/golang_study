package main

import "fmt"

func main() {
	a := 4 // 선업대입문 (a :=4 ) == (var a int = 4) == (var a = 4)
	b := 8

	fmt.Printf("a&b = %v\n", a&b)
	fmt.Printf("a|b = %v\n", a|b)
	fmt.Println("a^b =", a^b)

	c := 21
	fmt.Printf("c = %v, c의 십의 자리 = %v, c의 1의자리 = %v\n", c, (c/10)%10, c%10)

	d := 4 // 0100, ^4 = 1011, 0100 & 1011 = 0!
	fmt.Println("d&^d", d&^d)

	e := 4
	fmt.Println("e    : (00000100) =", e)
	fmt.Println("e<<1 : (00001000) =", e<<1)
	fmt.Println("e<<3 : (00100000) =", e<<3)

}

// 산술연산자 : +, -, *. /, %
// 비트연산자 : &, |, ^, ^ -> (a^b일땐 XOR, ^b알땐 NOT)
// 논리연산자 : <, >, ==, !=, <=, >=, &&, ||, !
// 단항연산자 : +(양), -(음), !, ^, *, &, <-, ++, --
// 기타연산자 : <<, >>, &^(clear), +=, -=, *=, /=
