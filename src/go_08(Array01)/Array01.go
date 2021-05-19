package main

import "fmt"

func main() {

	var A [10]int
	C := [10]int{6, 7, 8, 9, 10}

	for i := 0; i < 10; i++ {
		A[i] = i
	}

	for i := 0; i < len(C); i++ {
		C[i] = i * i
	}

	fmt.Println(A)
	fmt.Println(C)

	// 문자열도 배열의 일종임, golang은 utf8(1~3byte)을 사용함(아스키코드가 아님!!)
	S := "hello World" // [11]byte(byte == uint8, 0 ~ 255)
	fmt.Println(S)

	for i := 0; i < len(S); i++ {
		fmt.Print(S[i], ", ")
		fmt.Print(string(S[i]), ", ")
	}
	fmt.Println()

	SS := "hello 월드" // [11]byte(byte == uint8, 0 ~ 255)
	// SS := []byte("hello월드")
	fmt.Println(SS)

	for i := 0; i < len(SS); i++ {
		fmt.Print(SS[i], ", ")
		fmt.Print(string(SS[i]), ", ") // 한글은 한글자에 3byte를 차지 즉 236 155 148이 '월'이라는 한글자를 의미함
	}
	fmt.Println()
	// 배열 -> 메모리
	// 길이 -> 항목길이 * 개수
	// 문자열 -> 배열, 각 글자는 1 ~ 3 byte

	// 글자를 한글자씩 찍고싶을경우(문자열을 []byte가 아닌 []rune으로 사용하면 됨, rune(golang에서 지원하는 utf8을 지원하는 Type))
	SS2 := []rune(SS)
	fmt.Println("len(SS2) = ", len(SS2))
	fmt.Println("len(SS) = ", len(SS))

	for i := 0; i < len(SS2); i++ {
		fmt.Print(SS2[i], ", ")
		fmt.Print(string(SS2[i]), ", ") // 한글은 한글자에 3byte를 차지 즉 236 155 148이 '월'이라는 한글자를 의미함
	}

	fmt.Println()
}
