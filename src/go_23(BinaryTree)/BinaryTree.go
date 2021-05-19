// BST : binary search tree, O(logN)
// 주의할점! 최소신장트리로 만들어야 높은 효율을 가짐!(층이 적어야함)
// 해결법 : AVL Tree(tree를 회전시켜서 바꾸는 방법)

package main

import (
	"dataStruct"
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	tree := dataStruct.NewBinaryTree(50)

	numLst := MakeNumbers()

	for i := 0; i < len(numLst); i++ {
		if numLst[i] != 50 {
			tree.Root.AddNode(numLst[i])
		}
	}

	tree.Print()

	fmt.Println()
	fmt.Println()

	// if문의 초기값을 설정하고 사용하는 방법!
	if found, cnt := tree.Search(77); found {
		fmt.Println("found 77, cnt :", cnt)
	} else {
		fmt.Println("not found 77, cnt :", cnt)
	}

}

func MakeNumbers() [100]int {
	var rst [100]int

	for i := 0; i < 100; i++ {
		for {
			n := rand.Intn(100)
			duplicated := false
			for j := 0; j < i; j++ {
				if rst[j] == n {
					duplicated = true
					break
				}
			}
			if !duplicated {
				rst[i] = n
				break
			}
		}
	}
	return rst
}
