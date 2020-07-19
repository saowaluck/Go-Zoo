package main

import (
	"fmt"
	"strings"
)

func retriveParameterAsSlice(names ...string) string {
	return strings.Join(names, ", ")
}

func main() {
	fullName := retriveParameterAsSlice("thor", "alit", "un")
	fmt.Println(fullName)
}