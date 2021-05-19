// Map (= Dictionary, HashTable), Key와 Value로 이루어진 자료구조
// Map을 BST형태로 많이씀(O(logN)이므로 빨라서),
//
// HashMap : 같은입력-같은출력, 다른입력-다른출력 이나, 다른입력에 다른출력은 구현이 쉽지 않아서 대부분 다르다면 수용
//			 속도는 Hash 함수에서 바로 가져오므로 O(1) 상수속도(find,add,remove)를 가짐, 중요한것은 Hash 함수를 어떻게 구현하는가가 중요!
//			 Hash Function 조건
//			 1. 출력값의 범위가 정해져있다 (ex 0~ 1)
//			 2. 같은 입력 - 같은 출력
// 			 3. 다른 입력 - 다른 출력 (이는 구현하기 쉽지 않아서, 조금은 에외)
// 			 hashFunction 예시) Modular(%, OneWayFunction), sin, cos
// 			 Map의 단점 : 속도는 빠르지만 정렬은 불가, 정렬이 필요할때는 SortedMap을 사용( O(logN) )

package main

// 예제, Rolling Hash
// S0, S1, ... Sn
// Hi = ( H(i-1) * A + Si ) % B
// H0 = (S0) % B
// B는 소수를 사용하는것이 좋음(범위가 넓어짐)

import (
	"dataStruct"
	"fmt"
)

func main() {

	// Hash Function
	fmt.Println("paradox = ", dataStruct.Hash("paradox"))
	fmt.Println("paradox = ", dataStruct.Hash("paradox"))
	fmt.Println("Paradox = ", dataStruct.Hash("Paradox"))

	fmt.Println()

	// Map
	m := dataStruct.CreatMap()

	m.Add("dox", "012342587")
	m.Add("sdf", "039962284")
	m.Add("dow", "165418634")
	m.Add("sowp", "888878444")

	fmt.Println("dox = ", m.Get("dox"))
	fmt.Println("dfe = ", m.Get("dfe"))
	fmt.Println("sdf = ", m.Get("sdf"))
	fmt.Println("dow = ", m.Get("dow"))
	fmt.Println("sowp = ", m.Get("sowp"))

}
