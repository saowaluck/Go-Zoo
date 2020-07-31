package interfacemethod

type Count() int

func (c *Count) Incr() int {
	*c = *c + 1
	return int(*c)
}


type Counter interface {
	Incr() int
}

func onAPIHit(c Counter) {
	for i := 0; i < n; i++ {
		c.Incr()
	}
}

func main() {
	dummyCounter := Count(0)
	onAPIHit(100, &dummyCounter)
	fmt.Println(dummyCounter)
}
