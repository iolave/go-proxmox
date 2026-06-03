package main

func main() {
	s := new(S)
	s.i.ph()
}

type Inner struct {
	s S
}

func NewInner(s S) *Inner {
	i := new(Inner)
	i.s = s
	return i
}

func (i Inner) ph() {

}

type S struct {
	host string
	i    *Inner
}
