// heap : MaxHeap(부모노드가 무조건 자식노드 보다 크거나 같다)과 MinHeap(부모노드가 무조건 자식노드보다 작거나 같다)
//		  따라서 최대값과 최솟값이 최상위 노드에 존재함(), 속도 : Push&POP- O(logN),
//		  추가적으로 PQ(priority queue, 우선순위큐)를 heap으로 구현 할 수 있음
// 		  PQ: 선입선출이 아닌 우선순위에 의해 선출 됨
//  	  따라서 heap을 이용하면 정렬을 할 수 있음! heap정렬의 속도 : O(NlogN)(원래 2NlogN개 인데 비고법으로 상수를 제거하므로 2는 생략함)
//		  heap은 tree로 구현하기가 어렵(같은 층간의 대소비교힘듬) 따라서 slice(Array)로 보통 구햔함
package main

import (
	"dataStruct"
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	h := &dataStruct.Heap{}

	numLst := MakeRandNumbers(50)

	for i := 0; i < len(numLst); i++ {
		if numLst[i] != 50 {
			h.Push(numLst[i])
		}
	}

	h.Print()

	for i := 0; i < 10; i++ {
		fmt.Println(h.Pop())
	}

}

func MakeRandNumbers(len int) []int {
	var rst []int
	var length = len

	for i := 0; i < length; i++ {
		for {
			n := rand.Intn(length)
			duplicated := false
			for j := 0; j < i; j++ {
				if rst[j] == n {
					duplicated = true
					break
				}
			}
			if !duplicated {
				rst = append(rst, n)
				break
			}
		}
	}
	return rst
}
