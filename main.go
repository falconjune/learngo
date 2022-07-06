package main //프로젝트를 컴파일 하고 싶다는 뜻이고 그것을 사용할 것이란 얘기
//라이브러리등을 만든다면 package가 필요 없을수도 있다.

import (
	"fmt" //formatting을 위한 package이다.
	"strings"

	"github.com/kolonist/learngo/banking"
)

//함수를 선언할때 인자의 타입과 return 값의 타입을 모두 정해져야 됨.
func multiply(a, b int) int {
	return a * b
}

//return 값을 여러 개를 반환할 수 있음.
func lenAndUpper(name string) (int, string) {
	return len(name), strings.ToUpper(name)
}

//return을 끝에 붙이지 않고 함수 return 값 타입 선언하는데에서 설정할수 있음.
func lenAndUpper2(name string) (length int, uppercase2 string) {
	defer fmt.Println(("I'm Done")) //func이 끝난 다음에 코드를 exec할 수 있음.
	length = len(name)              //위에서 타입을 선언해주었으므로 축약형이나 타입을 선언해줄 필요는 없음.
	uppercase2 = strings.ToUpper(name)
	return
}

//인자를 한 번 여러개 보낼수도 있다. 단순히 출력만 하고 값을 return 해주지 않음.
func repeatMe(name ...string) {
	fmt.Println(name)
}

func superAdd(numbers ...int) {
	for index, number := range numbers {
		fmt.Println(index, number)
	}
}

func superAdd2(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}

func superAdd3(numbers ...int) int {
	total := 0
	for i := 0; i < len(numbers); i++ {
		total += numbers[i]
	}
	return total
}

func canIdrink(age int) bool {
	if koreanAge := age + 2; koreanAge < 18 {
		return false
	}
	return true
}

func canIdrink2(age int) bool {
	switch koreanAge := age + 2; koreanAge {
	case 18:
		return false
	case 50:
		return true
	case 0:
		return false
	}
	return false //예외가 있을 수 있으므로 이를 작성해주어야 됨.
}

//go에는 c언어의 struct와 같은 구조체가 있음.
type person struct {
	name    string
	age     int
	favFood []string
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
	repeatMe("abc", "def", "ghi")
	superAdd(1, 2, 3, 4, 5)
	fmt.Println(superAdd2(1, 1, 3))
	fmt.Println(superAdd2(1, 1, 5))
	fmt.Println((canIdrink(14)))
	//Go에서도 pointer를 지원함.
	a := 1
	b := &a
	*b = 20
	fmt.Println(a)
	//Python의 map과 비슷한 것이 있음. 다만 아예 같은 것은 아님.
	kolonist := map[string]string{"name": "minjun", "age": "23"}
	for key, value := range kolonist {
		fmt.Println(key, value)
	}
	favFood1 := []string{"kimchi", "ramen"}
	kolonist2 := person{name: "MINJUN", age: 23, favFood: favFood1}
	fmt.Println(kolonist2.name)

	//다른 package에 있는 struct를 통해 구조체를 만들 수 있음. 안에 있는 것도 소문자로 시작하면 pri임.
	//그로 인해 struct가 대문자여서 접근 가능하더라도 내부 원소는 제어할 수 없음.

	//다만 외부인이 마음대로 바꿀 수 없게 하기 위해 private으로 변경한 다음 생성자 같은 함수를 따로 만든다.
	account := banking.NewAccount("minjun")
	fmt.Println(account)
	account.Deposit(10)
	fmt.Println(account.Balance())
}

//test123
