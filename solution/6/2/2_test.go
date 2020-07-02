package p622

import (
	"fmt"
	"testing"
)

type A struct {
	p int
}

func (a *A) String() string {
	return fmt.Sprintf("%d", a.p)
}
func (a *A) Set(p int) {
	a.p = p
}

type B struct {
	p int
}

func (b B) String() string {
	return fmt.Sprintf("%d", b.p)
}
func (b B) Set(p int) {
	b.p = p
}

func TestGoArray(t *testing.T) {
	a := A{p: 3}
	a.Set(4)
	fmt.Printf("%s\n", a.String())

	b := B{p: 3}
	b.Set(4)
	fmt.Printf("%s\n", b.String())
}
