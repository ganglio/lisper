package lisper

import (
	"path/filepath"
	"reflect"
	"runtime"
	"strings"
)

type Op func(args ...Value) Value

func (o Op) Name() string {
	r := reflect.ValueOf(o).Pointer()
	nameFull := runtime.FuncForPC(r).Name()
	nameEnd := filepath.Ext(nameFull)
	name := strings.TrimPrefix(nameEnd, ".")
	return name
}

// Add calculates the sum between two values
func Add(args ...Value) (r Value) {

	if len(args) != 2 {
		panic("Invalid number of operands")
	}

	a, b := args[0], args[1]

	if a.t == vString || b.t == vString {
		return V(`"` + a.V.(string) + b.V.(string) + `"`)
	}

	switch a.t {
	case vInt:
		switch b.t {
		case vInt:
			r = V(a.V.(int64) + b.V.(int64))
		case vFloat:
			r = V(float64(a.V.(int64)) + b.V.(float64))
		}
	case vFloat:
		switch b.t {
		case vInt:
			r = V(a.V.(float64) + float64(b.V.(int64)))
		case vFloat:
			r = V(a.V.(float64) + b.V.(float64))
		}
	}

	return
}

// Sub calculates the difference between two values
func Sub(args ...Value) (r Value) {

	if len(args) != 2 {
		panic("Invalid number of operands")
	}

	a, b := args[0], args[1]

	if a.t == vString || b.t == vString {
		panic("Unsupported operands")
	}

	switch a.t {
	case vInt:
		switch b.t {
		case vInt:
			r = V(a.V.(int64) - b.V.(int64))
		case vFloat:
			r = V(float64(a.V.(int64)) - b.V.(float64))
		}
	case vFloat:
		switch b.t {
		case vInt:
			r = V(a.V.(float64) - float64(b.V.(int64)))
		case vFloat:
			r = V(a.V.(float64) - b.V.(float64))
		}
	}

	return
}

// Mul calculates the product between two values
func Mul(args ...Value) (r Value) {

	if len(args) != 2 {
		panic("Invalid number of operands")
	}

	a, b := args[0], args[1]

	if a.t == vString || b.t == vString {
		panic("Unsupported operands")
	}

	switch a.t {
	case vInt:
		switch b.t {
		case vInt:
			r = V(a.V.(int64) * b.V.(int64))
		case vFloat:
			r = V(float64(a.V.(int64)) * b.V.(float64))
		}
	case vFloat:
		switch b.t {
		case vInt:
			r = V(a.V.(float64) * float64(b.V.(int64)))
		case vFloat:
			r = V(a.V.(float64) * b.V.(float64))
		}
	}

	return
}

// Div calculates the division between two values
func Div(args ...Value) (r Value) {

	if len(args) != 2 {
		panic("Invalid number of operands")
	}

	a, b := args[0], args[1]

	if a.t == vString || b.t == vString {
		panic("Unsupported operands")
	}

	if b.t == vInt && b.V.(int64) == 0 {
		panic("Divide by zero")
	}

	if b.t == vFloat && floatEquals(b.V.(float64), 0) {
		panic("Divide by zero")
	}

	switch a.t {
	case vInt:
		switch b.t {
		case vInt:
			r = V(a.V.(int64) / b.V.(int64))
		case vFloat:
			r = V(float64(a.V.(int64)) / b.V.(float64))
		}
	case vFloat:
		switch b.t {
		case vInt:
			r = V(a.V.(float64) / float64(b.V.(int64)))
		case vFloat:
			r = V(a.V.(float64) / b.V.(float64))
		}
	}

	return
}

// Eq checks a == b
func Eq(args ...Value) Value {

	if len(args) != 2 {
		panic("Invalid number of operands")
	}

	a, b := args[0], args[1]

	if a.t != b.t {
		return FALSE
	}

	if a.t == vFloat {
		return V(floatEquals(a.V.(float64), b.V.(float64)))
	}

	return V(a.V == b.V)
}

// Gt checks a > b
func Gt(args ...Value) Value {

	if len(args) != 2 {
		panic("Invalid number of operands")
	}

	a, b := args[0], args[1]

	if a.t != b.t && (a.t == vString || b.t == vString) {
		return FALSE
	}

	if a.t == vString && b.t == vString {
		return V(a.V.(string) > b.V.(string))
	}

	var x, y float64

	if a.t == vFloat {
		x = a.V.(float64)
	} else {
		x = float64(a.V.(int64))
	}

	if b.t == vFloat {
		y = b.V.(float64)
	} else {
		y = float64(b.V.(int64))
	}

	return V(x > y)
}

// Ge checks a >= b
func Ge(args ...Value) Value {

	if len(args) != 2 {
		panic("Invalid number of operands")
	}

	a, b := args[0], args[1]

	if a.t != b.t && (a.t == vString || b.t == vString) {
		return FALSE
	}

	if a.t == vString && b.t == vString {
		return V(a.V.(string) >= b.V.(string))
	}

	var x, y float64

	if a.t == vFloat {
		x = a.V.(float64)
	} else {
		x = float64(a.V.(int64))
	}

	if b.t == vFloat {
		y = b.V.(float64)
	} else {
		y = float64(b.V.(int64))
	}

	return V(x >= y)
}

// Lt checks a < b
func Lt(args ...Value) Value {

	if len(args) != 2 {
		panic("Invalid number of operands")
	}

	a, b := args[0], args[1]

	if a.t != b.t && (a.t == vString || b.t == vString) {
		return FALSE
	}

	if a.t == vString && b.t == vString {
		return V(a.V.(string) < b.V.(string))
	}

	var x, y float64

	if a.t == vFloat {
		x = a.V.(float64)
	} else {
		x = float64(a.V.(int64))
	}

	if b.t == vFloat {
		y = b.V.(float64)
	} else {
		y = float64(b.V.(int64))
	}

	return V(x < y)
}

// Le checks a <= b
func Le(args ...Value) Value {

	if len(args) != 2 {
		panic("Invalid number of operands")
	}

	a, b := args[0], args[1]

	if a.t != b.t && (a.t == vString || b.t == vString) {
		return FALSE
	}

	if a.t == vString && b.t == vString {
		return V(a.V.(string) <= b.V.(string))
	}

	var x, y float64

	if a.t == vFloat {
		x = a.V.(float64)
	} else {
		x = float64(a.V.(int64))
	}

	if b.t == vFloat {
		y = b.V.(float64)
	} else {
		y = float64(b.V.(int64))
	}

	return V(x <= y)
}

// Ne checks a != b
func Ne(args ...Value) Value {

	if len(args) != 2 {
		panic("Invalid number of operands")
	}

	a, b := args[0], args[1]

	if a.t != b.t {
		return TRUE
	}

	if a.t == vString && b.t == vString {
		return V(a.V.(string) != b.V.(string))
	}

	var x, y float64

	if a.t == vFloat {
		x = a.V.(float64)
	} else {
		x = float64(a.V.(int64))
	}

	if b.t == vFloat {
		y = b.V.(float64)
	} else {
		y = float64(b.V.(int64))
	}

	return V(!floatEquals(x, y))
}

// First returns the head of a list
func First(args ...Value) Value {
	if len(args) != 1 {
		panic("Invalid number of operands")
	}

	a := args[0]

	if a.t != vList {
		panic("Argument must be a list")
	}

	if len(a.V.(List).V) < 1 {
		return NIL
	}

	return a.V.(List).V[0]
}

// Rest returns the tail of a list
func Rest(args ...Value) Value {
	if len(args) != 1 {
		panic("Invalid number of operands")
	}

	a := args[0]

	if a.t != vList {
		panic("Argument must be a list")
	}

	if len(a.V.(List).V) < 2 {
		return NIL
	}

	l := List{[]Value{}}

	for _, v := range a.V.(List).V[1:] {
		l.V = append(l.V, v)
	}

	return V(l)
}

// If (if cond case_true case_false). Checks if cond is true, if it is return case_true else returns case_false if available or V(nil)
func If(args ...Value) Value {
	if len(args) > 3 || len(args) < 2 {
		panic("Invalid number of operands")
	}

	c, t := args[0], args[1]

	if c == TRUE {
		return t
	} else if len(args) == 3 {
		return args[2]
	}

	return NIL
}

// Define allows the definition of a new symbol
func Define(args ...Value) Value {
	if len(args) != 2 {
		panic("Invalid number of operands")
	}

	a, b := args[0], args[1]

	if a.t != vSymbol {
		panic("Invalid operand type")
	}

	Sym[a] = b
	return NIL
}
