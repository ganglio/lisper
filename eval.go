package lisper

import "fmt"

func Eval(s string) Value {
	t := Tokenize(s)
	p := Parse(&t)
	return eval(p)
}

func resolve(op Op, args ...Value) (r []Value) {
	switch op.Name() {
	case "Define":
		first := true
		for _, v := range args {
			if first {
				r = append(r, v)
				first = false
			} else {
				r = append(r, eval(L(v)))
			}
		}
	default:
		for _, v := range args {
			r = append(r, eval(L(v)))
		}
	}
	return
}

func eval(l List) (r Value) {

	fmt.Printf("SYM:%+v\n", Sym)

	ch := l.C()

	for v := range ch {
		switch v.t {
		case vList:
			r = eval(v.V.(List))
		case vSymbol:
			if op, ok := Env[v]; ok {
				ops := []Value{}
				for o := range ch {
					ops = append(ops, o)
				}
				r = op(resolve(op, ops...)...)
			} else {
				if r, ok = Sym[v]; !ok {
					r = v
				}
			}
		default:
			r = v
		}
	}

	return
}
