// slice(배열)와 list의 장담점
// 배열의 장점, 1. cache miss rate가 낮다(베열은 연속된 메모리값을 사용하므로, cache 뭉텅이(4kbyte)에 게속 사용할만한 메모리가 있을 확룰이 높음)
// 			 2. random access 속도가 빠르다, 주소값과 데이터 크기를 알고있기 떄문에 바로 불러 올 수 있다.( O(1) )
// 배열의 단점, 1. 데이터를 제거 할 때 맨 앞과 뒤의 값을 제거할때는 O(1)로 속도가 빠르지만, 중간값을 제거할때는 O(N)의 속도로 느리다. (새로운 배열에 복사하는 방식이기 때문)
//			 2. 데이터를 추가할때(append) 배열의 capacity를 초과하면 O(N)의 속도로 느리다.
//
// list의 장점,1. 데이터를 추가하거나 제거(중간)할때, DoubleLinkedList를 사용하면, O(1)의 속도로 빠르다.(SingleLinkedList는 그렇지 않다: O(N) )
// list의 단점,1. randem access 속도가 느리다. O(N)으로 root부터 찾는 데이터까지 계속 반복해야 하므로,
// 			 2. cache miss rate가 높다. list는 메모리상에서 따로따로 존재하기 때문이다.

package main

import "fmt"

func main() {
	list := &LinkedList{}

	list.AddNode(0)

	for i := 1; i < 10; i++ {
		list.AddNode(i)
	}

	list.PrintNodes()

	list.RemoveNode(list.root.next)

	list.PrintNodes()

	list.RemoveNode(list.root)

	list.PrintNodes()

	list.RemoveNode(list.tail)

	list.PrintNodes()

	fmt.Printf("tail:%d\n", list.tail.val)

	list.PrintReverse()

}

type Node struct {
	next *Node
	prev *Node
	val  int
}

type LinkedList struct {
	root *Node
	tail *Node
}

func (l *LinkedList) AddNode(val int) {
	if l.root == nil {
		l.root = &Node{val: val}
		l.tail = l.root
		return
	}

	l.tail.next = &Node{val: val}
	prev := l.tail
	l.tail = l.tail.next
	l.tail.prev = prev
}

func (l *LinkedList) RemoveNode(node *Node) {
	if node == l.root {
		l.root = l.root.next
		l.root.prev = nil
		node.next = nil
		return
	}

	prev := node.prev

	if node == l.tail {
		prev.next = nil
		l.tail.prev = nil
		l.tail = prev
	} else {
		node.next.prev = prev
		prev.next = prev.next.next
	}
	node.prev = nil
	node.next = nil
}

func (l *LinkedList) PrintNodes() {
	node := l.root
	for node.next != nil {
		fmt.Printf("%d -> ", node.val)
		node = node.next
	}
	fmt.Printf("%d\n", node.val)
}

func (l *LinkedList) PrintReverse() {
	node := l.tail
	for node.prev != nil {
		fmt.Printf("%d -> ", node.val)
		node = node.prev
	}
	fmt.Printf("%d\n", node.val)
}
