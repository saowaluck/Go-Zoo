package main

import "fmt"

func Len(s string) int {
	return len(s)
}

func main() {
	// Name funcs เป็น การเขียน func ใน GO
	l := Len("Hello world 👋")
	fmt.Printf("Hello world 👋 has len %v", l)
}
