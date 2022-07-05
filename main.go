package main //프로젝트를 컴파일 하고 싶다는 뜻이고 그것을 사용할 것이란 얘기
//라이브러리등을 만든다면 package가 필요 없을수도 있다.

import (
	"fmt" //formatting을 위한 package이다.

	"github.com/kolonist/learngo/something"
)

func main() {
	fmt.Println("hello world!") //export 하기 위해서 대문자로 시작되는 것.
	something.SayHello()
}
