package main // 프로그램의 시작점은 main!(instructer point)
// package	 : 기능등을 묶어 놓음(기능,모듈 등등)
// library	 : 책,지식 -> 참고지식등을 모아놓은것(필요한 기능들을 묶어 놓은것)
// module	 : 기능들을 묶어 놓았지만 library보다는 좀 더 포커싱해놓은것(ex:입력모듈)
// framework : 기능 묶음 절차(뼈대), 메뉴얼과 같이
// engine	 : 기능의 묶음 뿐만 아니라 필요한 다른 프로그램(툴)도 다 묶어 놓은것

import "fmt" // import : package에 있는 것을 가져온다
// fmt : 표준패키지(Go에서 만든)로 format의 약자

func main() {
	fmt.Println("Hello World!")
}

// terminal 에서 빌드 -> 해당폴터, go build
// terminal 에서 실행(인터프린트방식?!) -> go run 파일명
