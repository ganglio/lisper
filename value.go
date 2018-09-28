package lisper

import (
	"fmt"
)

const (
	v_int = iota
	v_float
	v_string
	v_bool
)

// Value is a generic value
type Value struct {
	v interface{}
	t int
}

// V creates a new value
func V(x interface{}) Value {
	switch v := x.(type) {
	case int:
		return Value{int64(v), v_int}
	case int64:
		return Value{int64(v), v_int}
	case float64:
		return Value{v, v_float}
	case string:
		return Value{v, v_string}
	case bool:
		return Value{v, v_bool}
	default:
		panic(fmt.Sprintf("Unsupported Value type [%T]", x))
	}
}
