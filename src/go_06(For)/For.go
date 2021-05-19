package main

import "fmt"

func main() {
	i := 0

	for i < 5 {
		fmt.Println(i)
		i++
	}
	fmt.Println("최종 i 값 : ", i)

	for j := 0; j < 5; j++ {
		fmt.Println("j = ", j)
	}

	var h int
	for h = 0; h < 5; h++ {
		fmt.Println("h = ", h)
	}

	var k int = 0
	for { // 아무것도 안 명시할 경우 참으로 간주, 다른 언어의 while과 동일, 무한루프
		k++
		if k == 5 {
			break
		} else if k == 3 {
			continue
		}

		fmt.Println("k = ", k)
	}
}
