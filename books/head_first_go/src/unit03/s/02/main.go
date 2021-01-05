package main

import "fmt"

func main() {

	// %f를 이용해서 부동 소수점 숫자를 표현 할 수 있음
	fmt.Printf("A float: %f\n", 3.1415)

	// %d를 이용해서 십진수 정수를 표현 할 수 있음
	fmt.Printf("An integer: %d\n", 15)

	// %s를 이용해서 문자열을 표현 할 수 있음
	fmt.Printf("A string: %s\n", "hello")

	// %t를 이용해서 부울 형을 표현 할 수 있음
	fmt.Printf("A boolean: %t\n", false)

	// %v를 이용해서 값의 타입에 따라 적절히 형식화된 값을 표현 할 수 있음
	fmt.Printf("Values: %v %v %v\n", 1.2, "\t", true)

	// %#v를 이용해서 Go 프로그램 코드에 나타나는 그대로 표현 할 수 있음
	fmt.Printf("Values: %#v %#v %#v\n", 1.2, "\t", true)

	// %T를 이용해서 제공된 값의 타입을 표현 할 수 있음
	fmt.Printf("Type: %T %T %T\n", 1.2, "\t", true)

	// %%를 이용해서 백분율 기호를 표현 할 수 있음
	fmt.Printf("Percent sign: %%\n")
}
