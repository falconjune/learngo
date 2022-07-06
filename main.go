package main //프로젝트를 컴파일 하고 싶다는 뜻이고 그것을 사용할 것이란 얘기
//라이브러리등을 만든다면 package가 필요 없을수도 있다.

import (
	"fmt" //formatting을 위한 package이다.
	"strings"
)

//함수를 선언할때 인자의 타입과 return 값의 타입을 모두 정해져야 됨.
func multiply(a, b int) int {
	return a * b
}

//return 값을 여러 개를 반환할 수 있음.
func lenAndUpper(name string) (int, string) {
	return len(name), strings.ToUpper(name)
}

// a3 := 4 변수를 선언하는데에는 축약형이 있는데 이는 func 안에서만 가능하고 밖에서는 축약형을 사용할 수 없음.

func main() {
	//const, var
	a2 := 1234 //축약형. 넣고자 하는 값을 기반으로 타입을 정해줌. command + hover를 통해 int형으로 선언된 것을 알 수 있음.
	const c1 string = "testc"
	var a1 string = "testv" //변수는 선언하면 무조건 사용되어야 함. 안그러면 에러
	fmt.Println(a1)         //export 하기 위해서 대문자로 시작되는 것.
	fmt.Println(a2)
	fmt.Println(multiply(2, 2))
	totalLength, upperName := lenAndUpper("kolonist") //2개의 값을 반환하므로 무조건 2개의 값을 모두 받아줘야 됨.
	totalLength, _ = lenAndUpper("teset_")            //만약 하나의 값만 받고 싶다면 _을 통해 값을 무시할 수 있다.
	//변수 초기화는 언제나 할 수 있지만, 한 번 선언한 변수는 타입을 변경할 수 없음.
	fmt.Println(totalLength, upperName)
}

//test123
