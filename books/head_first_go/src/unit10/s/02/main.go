package main

import "fmt"

type Date struct {
	Year  int
	Month int
	Day   int
}

// 설정자 메서드에서는 포인터 리시버를 사용해야한다.
func (d *Date) SetYear(year int) {
	d.Year = year
}

func main() {
	date := Date{}
	date.SetYear(2019)
	fmt.Println(date.Year)
}
