package main

type LargeObject struct {
	a int
	b int
	c int
	g int
}

//go:noinline
func newLargeObject(a, b, c, g int) *LargeObject {
	res := &LargeObject{
		a: a,
		b: b,
		c: c,
		g: g,
	}

	return res
}

func main() {
	a := 1
	b := 1
	c := 2
	g := 46

	obj := newLargeObject(a, b, c, g)

	_ = obj
}
