package main

import "fmt"

func main() {
	x := 30

	switch x {
	case 31:
		fmt.Println("31")
	case 30:
		fmt.Println("30")
	case 29:
		fmt.Println("29")
	}

	switch { // 비워져 있을 경우 항상 참으로 간주함
	case x < 40:
		fmt.Println("x는 40보다 작다")
	case x < 35: // 해당하는 case가 2개 이상일 경우 가장 첫번째것 하나만 실행됨!
		fmt.Println("x는 35보다 작다")
	case x < 30:
		fmt.Println("x는 30보다 작다")
	case x < 25:
		fmt.Println("x는 25보다 크다")
	}

}
