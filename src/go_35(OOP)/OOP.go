// Object-Oriented Programming
// interface : Object끼리의 관계를 따로 정의해준것(decoupling) - 이 과정에서 확장성을 얻을 수 있음

package main

import "fmt"

func main() {
	bread := &Bread{}
	// jam := &StrawberryJam{}
	jam := &OrangeJam{}
	bread.PutJam(jam)

	fmt.Println(bread)
}

type Jam interface {
	GetOneSpoon() SpoonOfJam
}

type SpoonOfJam interface {
	String() string
}

type Bread struct {
	val string
}

type StrawberryJam struct {
}

type OrangeJam struct {
}

type SpoonOfStrawberryJam struct {
}

type SpoonOfOrangeJam struct {
}

func (s *SpoonOfStrawberryJam) String() string {
	return "+ strawberry"
}

func (s *SpoonOfOrangeJam) String() string {
	return "+ Orange"
}

func (j *StrawberryJam) GetOneSpoon() SpoonOfJam {
	return &SpoonOfStrawberryJam{}
}

func (j *OrangeJam) GetOneSpoon() SpoonOfJam {
	return &SpoonOfOrangeJam{}
}

// func (b *Bread) PutJam(jam *StrawberryJam) {
func (b *Bread) PutJam(jam Jam) {
	spoon := jam.GetOneSpoon()
	b.val += spoon.String()
}

func (b *Bread) String() string {
	return "bread " + b.val
}
