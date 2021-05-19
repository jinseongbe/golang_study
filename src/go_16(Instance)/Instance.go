package main

import "fmt"

func main() {

	a := Student{"aaa", 20, 10}
	fmt.Println(a)

	var b *Student
	b = &a // Reference type

	c := a // Value Type
	c.age = 500

	b.age = 100

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

	SetName(&a, "bbb")
	fmt.Println(a)

	SetName(b, "ccc")
	fmt.Println(a)

	// method로 사용
	c.SetName("CcCc") // 여기서 c를 Instance라고 함, 생명을 가지는 것?! 주체가 되는것(주어)!
	fmt.Println(c)

}

type Student struct {
	name  string
	age   int
	grade int
}

func SetName(t *Student, newName string) { // Student를 pointer로 받지 않으면 바뀌지 않음(copy)
	t.name = newName
}

// 위의 함수를 method 형태로 바꾸면
func (t *Student) SetName(newName string) { // Student를 pointer로 받지 않으면 바뀌지 않음(copy)
	t.name = newName
}
