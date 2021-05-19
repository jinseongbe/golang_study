package main

import "fmt"

func main() {
	var a int
	var b int

	// a := 4 // 선업대입문 (a :=4 ) == (var a int = 4) == (var a = 4)

	a = 3
	b = 4

	fmt.Println(a + b)
}

// 32bit/64bit -> 기본연산이 되는 사이즈
//
// int		: 정수 4/8byte
// int8		: 정수 1byte		-> (2^7 - 1)  -> (-128 ~ 127)
// uint8	: 양의 정수 1byte	 -> (2^8 - 1)  -> (0 ~ 255)
// int16	: 정수 2byte		-> (2^15 - 1) -> (-32768 ~ 32767)
// uint16	: 양의 정수 2byte	 -> (2^16 - 1) -> (0 ~ 65535)
// int32	: 정수 4byte		-> (2^31 - 1) -> (-21억 ~ 21억)
// uint32	: 양의 정수 4byte 	 -> (2^32 - 1) -> (0 ~ 42억)
// int64	: 정수 8byte
// uint64	: 양의 정수 6byte
//
// float32	: 실수 4byte		-> 숫자부 7개까지, (@@@@@@@ * 10^(-6))
// float64	: 실수 8byte		-> 숫자부 15개까지, (@@@@@@@@@@@@@@@ * 10^(-16~26))
//
// bool		: 논리 true/false
// string 	: 문자열, 값에 따라 메모리크기가 달라진다, 보통 한글자에 1byte
//			: type rune = 한글자의 타입으로 1~3byte, string -> [] rune의 배열이라고 생각하면 됨
