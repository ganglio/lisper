package lisper

const (
	a_value = iota
	a_list
)

type Atom struct {
	V interface{} `json:"a"`
	t int
}

func A(x interface{}) Atom {
	switch v := x.(type) {
	case Value:
		return Atom{v, a_value}
	case List:
		return Atom{v, a_list}
	default:
		return Atom{V(v), a_value}
	}
}

type List struct {
	V []Atom `json:"l"`
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
	if l.V == nil {
		l.V = []Atom{}
	}

	for _, i := range items {
		// switch v := i.(type) {
		// case List:
		// 	for _, k := range v.V {
		// 		l.V = append(l.V, k)
		// 	}
		// default:
		l.V = append(l.V, A(i))
		// }
	}
}
