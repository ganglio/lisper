package lisper

// List is a list of Atoms
type List struct {
	V []Value `json:"l"`
}

// L creates a new List
func L(items ...interface{}) List {
	o := []Value{}
	for _, i := range items {
		switch x := i.(type) {
		case Value:
			o = append(o, x)
		default:
			o = append(o, V(x))
		}
	}

	return List{o}
}

// Append allows appending to the current List
func (l *List) Append(items ...interface{}) {
	if l.V == nil {
		l.V = []Value{}
	}

	for _, i := range items {
		switch v := i.(type) {
		case Value:
			l.V = append(l.V, v)
		default:
			l.V = append(l.V, V(v))
		}
	}
}

// C returns a channel emitting all the values in the list. Handy to iterate over the list.
func (l *List) C() <-chan Value {
	ch := make(chan Value)
	go func() {
		for _, v := range l.V {
			ch <- v
		}
		close(ch)
	}()
	return ch
}
