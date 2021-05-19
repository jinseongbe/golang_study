package main

import "fmt"

func main() {
	var a int = 100
	var b int = 200
	var p *int

	p = &a

	fmt.Println(a)
	fmt.Println(*p)
	fmt.Println(p)

	p = &b

	fmt.Println(b)
	fmt.Println(*p)
	fmt.Println(p)
	fmt.Println()

	// 포인터 예제 01
	var x int = 0

	Increase1(x)
	fmt.Println(x)

	Increase2(&x)
	fmt.Println(x)

	fmt.Println()

	// 포인터 예제 02
	var s1 Student = Student{name: "jackjack", age: 23, class: "math", grade: "A"}

	s1.PrintGrade()
	s1.InputGrade("science", "c")
	s1.PrintGrade()

}

func Increase1(x int) {
	x++
}

func Increase2(x *int) {
	*x = *x + 1
}

type Student struct {
	name string
	age  int

	grade string
	class string
}

func (s *Student) PrintGrade() {
	fmt.Println(s.class, s.grade)
}

// c,c++과 다르게 golang에서는 s->class가 아니라 s가 포인터인지 아닌지에 따라 알아서 s.class로 쓰면 선택해서 사용해줌
func (s *Student) InputGrade(class string, grade string) {
	s.class = class
	s.grade = grade
}
