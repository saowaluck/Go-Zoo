package main

import "fmt"

type Count int

func (c Count) IncreaseCount() int {
	c = c + 1
	return int(c)
}

func (c *Count) IncreaseCountAsPointer() int {
	*c = *c + 1
	return int(*c)
}

func main() {
	// it not increase number because of this c the new copy (but not reuse the variable in the memory)
	var c Count
	count := c.IncreaseCount()
	fmt.Println(count)

	// it increse number because of the c is a pointer it's can points to the original of c
	count2 := c.IncreaseCountAsPointer()
	count2 = c.IncreaseCountAsPointer()
	count2 = c.IncreaseCountAsPointer()
	fmt.Println(count2)

}
