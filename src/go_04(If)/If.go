package main

import (
	"bufio"
	"fmt"
	"os"      // 입출력
	"strconv" // 문자열은 숫자로 바꿔줌
	"strings"
)

func main() {

	var a bool

	a = true

	if a {
		fmt.Println("true!")
	} else if !a {
		fmt.Println("false")
	} else {
		fmt.Println("never")
	}

	// 예제 : calculator

	fmt.Println("숫자를 입력하세요")
	reader := bufio.NewReader(os.Stdin)
	line, _ := reader.ReadString('\n') // \n이 나올때까지 읽는다, '_' 해당 변수는 처리하지 않겠다는 의미(ReadString이 String과 error을 반환하는데 String만 사용하므로)
	line = strings.TrimSpace(line)     // 빈칸같은것을 제거
	n1, _ := strconv.Atoi(line)        // 문자를 숫자로 바꿔줌

	line, _ = reader.ReadString('\n')
	line = strings.TrimSpace(line)
	n2, _ := strconv.Atoi(line)

	fmt.Printf("입력하신 숫자는 %d, %d 입니다", n1, n2)

	fmt.Println("연산자를 입력하세요")
	line, _ = reader.ReadString('\n')
	line = strings.TrimSpace(line)

	if line == "+" {
		fmt.Printf("%d + %d = %d", n1, n2, n1+n2)
	} else if line == "-" {
		fmt.Printf("%d - %d = %d", n1, n2, n1-n2)
	} else if line == "*" || line == "x" {
		fmt.Printf("%d * %d = %d", n1, n2, n1*n2)
	} else if line == "/" {
		fmt.Printf("%d / %d = %d", n1, n2, n1/n2)
	} else {
		fmt.Printf("잘못 입력하셨습니다.")
	}
}
