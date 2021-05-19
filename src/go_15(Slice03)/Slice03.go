package main

import (
	"fmt"
)

func main() {
	s := make([]int, 5, 5)
	var t []int
	t = s // slice는 pointer(배열의), len, cap 의 3가지를 struct로 가지고 있고 assign할때 이 값이 그대로 copy되는것임
	// 따라서 slice는 같은 포인터주소값을 가지게 됨!

	fmt.Println(s, len(s), cap(s), &s[0])

	s = append(s, 1, 2, 5, 3, 4, 2, 7, 5) // cap이 5인 slice에 7개를 append해줌

	fmt.Println(s, len(s), cap(s), &s[0]) // 추가된 capacity의 용량은, append한 것과 원래 slice의 length중에서 큰것의 두배를 하여 새로 키워준다!

	fmt.Println(t, len(t), cap(t), &t[0])
}
