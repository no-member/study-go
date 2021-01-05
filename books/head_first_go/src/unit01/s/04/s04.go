package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println(reflect.TypeOf(43))
	fmt.Println(reflect.TypeOf(3.12343))
	fmt.Println(reflect.TypeOf(true))
	fmt.Println(reflect.TypeOf("Hello, Go!"))
}
