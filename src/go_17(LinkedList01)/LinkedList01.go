// 맨 끝을 찾는 첫번째 방법 O(N)
package main

import "fmt"

type Node struct {
	next *Node
	val  int
}

func main() {
	var root *Node // 맨 처음 시작하는 node는 root로 하고 기억해야함!
	root = &Node{val: 0}

	for i := 1; i < 10; i++ {
		root.AddNode(i)
	}

	node := root
	for node.next != nil {
		fmt.Printf("%d -> ", node.val)
		node = node.next
	}
	fmt.Printf("%d\n", node.val)
}

func (root *Node) AddNode(val int) {
	var tail *Node
	tail = root
	for tail.next != nil {
		tail = tail.next
	}

	node := &Node{val: val}
	tail.next = node
}
