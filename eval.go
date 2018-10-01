package lisper

import (
// "fmt"
)

func Eval(s string) Value {
	t := Tokenize(s)
	p := Parse(&t)
	return eval(p)
}

func eval(l List) (r Value) {

	ch := l.C()

	for v := range ch {
		switch v.t {
		case vList:
			r = eval(v.V.(List))
		default:
			if op, ok := Env[v]; ok {
				o1 := eval(L(<-ch))
				o2 := eval(L(<-ch))
				r = op(o1, o2)
			} else {
				r = v
			}

		}
	}

	return
}
