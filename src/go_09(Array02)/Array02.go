package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	arr1 := [5]int{1, 2, 3, 4, 5}
	arr2 := [5]int{}
	arr3 := [5]int{}

	for i := 0; i < len(arr1); i++ {
		arr2[i] = arr1[i]
		arr3[i] = arr1[(len(arr1)-1)-i] // 2n의 알고리즘 복잡도
	}

	fmt.Println(arr2)
	fmt.Println(arr3)
	fmt.Println()

	// 역순으로 자리바꾸기
	arr4 := [10]int{1, 2, 3, 4, 5, 6}
	fmt.Println(arr4)
	for i := 0; i < len(arr4)/2; i++ {
		arr4[i], arr4[len(arr4)-1-i] = arr4[len(arr4)-1-i], arr4[i] // n/2의 알고리즘 복잡도
	}
	fmt.Println(arr4)

	// 정렬 알고리즘 중 하나(RADIX), 다양한 알고리즘이 존재함
	var arr5 [50]int

	for i := 0; i < len(arr5); i++ {
		arr5[i] = rand.Intn(70)
	}

	fmt.Println(arr5)

	temp := [70]int{}
	for i := 0; i < len(arr5); i++ {
		idx := arr5[i]
		temp[idx]++
	}

	idx2 := 0
	for i := 0; i < len(temp); i++ {
		for j := 0; j < temp[i]; j++ {
			arr5[idx2] = i
			idx2++
		}
		// fmt.Println(arr5)
	}

	fmt.Println(arr5)

}
