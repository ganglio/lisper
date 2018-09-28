package lisper

const (
	a_value = iota
	a_list
)

// Atom is an atomic AST element it can be either a List or a Value
type Atom struct {
	V interface{} `json:"a"`
	t int
}

// A creates a new Atom
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

// List is a list of Atoms
type List struct {
	V []Atom `json:"l"`
}

// L creates a new List
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

// Append allows appending to the current List
func (l *List) Append(items ...interface{}) {
	if l.V == nil {
		l.V = []Atom{}
	}

	for _, i := range items {
		switch v := i.(type) {
		case Atom:
			l.V = append(l.V, v)
		default:
			l.V = append(l.V, A(v))
		}
	}
}
