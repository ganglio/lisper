package lisper

const (
	aValue = iota
	aList
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
		return Atom{v, aValue}
	case List:
		return Atom{v, aList}
	default:
		return Atom{V(v), aValue}
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
