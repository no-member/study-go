package main

import "strings"

func RepeatStr(repeat_count int, target_str string) string {
	// I don't want to switch positions in side function
	return strings.Repeat(target_str, repeat_count)
}
