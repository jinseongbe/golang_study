// golang에 내장되어 있는 map
package main

import "fmt"

type Key struct {
	v int
}

type Value struct {
	v int
}

func main() {
	// var m map[string]string
	// var m1 map[int]string 이런식으로 자유롭게 사용가능
	// var m2 map[Key]Value
	// var m3 map[Key]*Value

	// map은 초기화를 해준 다음에 사용해야함!
	var m map[string]string
	m = make(map[string]string)
	// 선언과 초기화를 같이 할 수 도 있음
	m1 := make(map[int]string)
	m2 := make(map[int]int)

	// Add
	m["abc"] = "0123123"
	m1[29] = "tN"

	// Read
	fmt.Println("m[\"abc\"] :", m["abc"])
	fmt.Println("m1[29] :", m1[29])
	// 없는 값을 넣으면 golang에서는 기본값을 지정해줌 (빈문자열, 0, false)
	fmt.Println("m1[100] :", m1[100])
	fmt.Println("m2[10] :", m2[10])
	fmt.Println()
	// 위의 성질을 사용하여, C언어의 set 기능을 사용 할 수 있음
	m3 := make(map[int]bool)
	m3[4] = true
	fmt.Println(m3[4], m3[100])
	fmt.Println()
	// 문제는 기본값인지 내가 입력해준 값인지 구분이 안됨!! 주의해야함!
	// 따라서 golang에서는 값이 있는지 없는지에 대한 기능을 제공함(해당값이 있는지 없는지에 대한 true,false를 같이 반환해줌)
	m2[0] = 0
	m2[1] = 0
	v0, chk0 := m2[0]
	v1, chk1 := m2[1]
	v2, chk2 := m2[77]
	fmt.Println(v0, v1, v2)
	fmt.Println(chk0, chk1, chk2)
	fmt.Println()

	// Delete
	delete(m2, 0)
	v0, chk0 = m2[0]
	fmt.Println(v0, v1, v2)
	fmt.Println(chk0, chk1, chk2)
	fmt.Println()

	// 순회
	m2[10] = 345
	m2[11] = 583
	m2[12] = 296
	m2[13] = 502
	m2[14] = 304
	for key, value := range m2 {
		fmt.Println(key, "=", value) // 순차적으로 반환되는것은 아님!
	}

	// range Test
	lst := make([]int, 10)
	for i := 0; i < len(lst); i++ {
		lst[i] = (i-10)*i + 100%(i*7+13)
	}
	fmt.Println(lst)

	for val := range lst {
		fmt.Println(val, lst[val])
	}

}
