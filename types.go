package lisper

import (
	"fmt"
)

const (
	a_value = iota
	a_list
)

type Atom struct {
	v interface{}
	t int
}

func A(x interface{}) Atom {
	switch v := x.(type) {
	case Value:
		return Atom{v, a_value}
	case List:
		return Atom{v, a_list}
	default:
		panic(fmt.Sprintf("Unsupported Atom type [%T]", x))
	}
}

type List struct {
	v []Atom
}

func L(items ...interface{}) List {
	o := []Atom{}
	for _, i := range items {
		switch x := i.(type) {
		case Atom:
			o = append(o, x)
		default:
			o = append(o, A(x))
		}
	}

	return List{o}
}

func (l *List) Append(items ...interface{}) {
	if l.v == nil {
		l.v = []Atom{}
	}

	for _, i := range items {
		l.v = append(l.v, A(i))
	}
}
