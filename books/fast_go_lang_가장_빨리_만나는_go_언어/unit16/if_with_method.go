package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	var b []byte
	var err error

	b, err = ioutil.ReadFile("./hello.txt")

	if err == nil {
		fmt.Printf("%s", b)
	}

	fmt.Println()

	if b_inside_if, err_inside_if := ioutil.ReadFile("./hello.txt"); err_inside_if == nil {
		fmt.Printf("%s", b_inside_if)
	}
}
