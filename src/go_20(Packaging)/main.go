package main

import (
	"dataStruct"
	"fmt"
)

func main() {
	list := &dataStruct.LinkedList{}

	list.AddNode(0)

	for i := 1; i < 10; i++ {
		list.AddNode(i)
	}

	list.PrintNodes()

	list.RemoveNode(list.Root.Next)

	list.PrintNodes()

	list.RemoveNode(list.Root)

	list.PrintNodes()

	list.RemoveNode(list.Tail)

	list.PrintNodes()

	fmt.Printf("tail:%d\n", list.Tail.Val)

	list.PrintReverse()

}
