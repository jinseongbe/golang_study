package main

import "fmt"

// golang에서는 다른 언어에서 class가 structure로 사용됨

type Person struct {
	name string
	age  int
}

// 현대의 언어들은 structure에 기능을 추가 할 수 있음!
// 다른 언어와 다르게 golang은 매소드가 구조체 밖에서 정의됨! 굉장히 특이함!
func (p Person) PrintName() {
	fmt.Println(p.name)
}

// 사실 객체에 속하지 않은 함수로도 구현 가능, 그러나 이 경우에는 객체에 포함되어 있지 않는것이므로 위에것과는 다름!
func PrintName(p Person) {
	fmt.Println(p.name)
}

type Student struct {
	name     string
	class    int
	sungjuck Sungjuck
}

type Sungjuck struct {
	name  string
	grade string
}

func (s Student) ViewSungjuck() {
	fmt.Println(s.sungjuck)
}

func main() {

	var p Person
	p.name = "smith"
	p.age = 10
	p1 := Person{"jackjack", 15}
	p2 := Person{name: "kail"}
	p3 := Person{}

	fmt.Println(p, p1, p2, p3)

	p.PrintName()
	PrintName(p)
	p1.PrintName()

	fmt.Println()

	var s Student
	s.name = "ChulSu"
	s.class = 1

	s.sungjuck.name = "math"
	s.sungjuck.grade = "c"

	s.ViewSungjuck()

}
