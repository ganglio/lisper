package lisper

func Parse(tokens *Tokens) (l List) {
	if len(tokens.t) == 0 {
		panic("unexpected EOF")
	}

	for t, last := tokens.Next(); !last; t, last = tokens.Next() {
		if t == "(" {
			l.Append(Parse(tokens))
		} else if t == ")" {
			return
		} else {
			l.Append(t)
		}
	}
	return
}

func floatEquals(a, b float64) bool {
	var eps = 0.00000001
	if (a-b) < eps && (b-a) < eps {
		return true
	}
	return false
}
