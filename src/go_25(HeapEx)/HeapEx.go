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
)

func main() {
	h := &dataStruct.Heap{}

	// [-1, 3, -1, 5, 4, 7, 11, 6], 2번째 큰값
	// 이때 알고리즘은 N개의 배열에서 M번째를 찾는것인데,
	// 하나씩 할경우는 O(NM)
	// 정렬을 사용할 경우 O(NlogN)
	// Heap을 사용할 경우 O(NlogM) 으로 그래프를 그려보면 N,M이 커질수록 Heap의 속도가 가장 빠른것을 알 수 있음
	nums := []int{-1, 3, -1, 5, 4, 7, 11, 6}

	for i := 0; i < len(nums); i++ {
		h.Push(nums[i])
		if h.Count() > 2 {
			h.Pop()
		}
	}

	fmt.Print(h.Pop())

}
