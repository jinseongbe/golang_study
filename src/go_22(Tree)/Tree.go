// DFS : Depth First Search
// BFS : Breadth First Search

package main

import (
	"dataStruct"
	"fmt"
)

func main() {
	// tree 만들기
	tree := dataStruct.Tree{}

	val := 1
	tree.AddNode(val)
	val++

	for i := 0; i < 3; i++ {
		tree.Root.AddNode(val)
		val++
	}

	for i := 0; i < len(tree.Root.Childs); i++ {
		for j := 0; j < 2; j++ {
			tree.Root.Childs[i].AddNode(val)
			val++
		}
	}

	// DFS (재귀호출로 구현한 DFS)
	tree.DFS1()

	fmt.Println()

	// DFS (stack으로 구현한 DFS)
	tree.DFS2()

	fmt.Println()

	// BFS
	tree.BFS()

	fmt.Println()

}
