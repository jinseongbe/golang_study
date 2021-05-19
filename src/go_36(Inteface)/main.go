package main

import (
	"fmt"
	"strconv"
)

type StructA struct {
}

type StructB struct {
}

func (a *StructA) AAA(x int) int {
	return x * x
}

func (a *StructA) BBB(x int) string {
	return "X = " + strconv.Itoa(x)
}

func (b *StructB) AAA(x int) int {
	return x * 2
}

func main() {
	var c InterfaceA
	c = &StructA{}

	fmt.Println(c.BBB(3))

	// var d InterfaceA
	// d = &StructB{}
	//cannot use &StructB literal (type *StructB) as type InterfaceA in assignment: *StructB does not implement InterfaceA (missing BBB method)
	// StructB에 BBB method가 구현되어 있지 않으므로 interface를 사용하지 못하는 에러 발생!
}
