package lisper

import (
	"fmt"
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
				fmt.Printf("op %#v\n", v)
				ops := []Value{}
				for o := range ch {
					fmt.Printf("arg %#v\n", o)
					ops = append(ops, eval(L(o)))
				}
				fmt.Printf("Ops %#v\n", ops)
				r = op(ops...)
			} else {
				r = v
			}

		}
	}

	return
}
