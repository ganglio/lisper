package lisper

import "fmt"

func Eval(s string) Value {
	t := Tokenize(s)
	p := Parse(&t)
	return eval(p)
}

func eval(l List) (r Value) {

	ch := l.C()

	for v := range ch {
		fmt.Printf("v: %#v %#v\n", v, PLUS)
		switch v.t {
		case vList:
			r = eval(v.V.(List))
		case vSymbol:
			if op, ok := Env[v]; ok {
				ops := []Value{}
				for o := range ch {
					ops = append(ops, eval(L(o)))
				}
				r = op(ops...)
			} else {
				r = Sym[v]
			}
		default:
			r = v
		}
	}

	return
}
