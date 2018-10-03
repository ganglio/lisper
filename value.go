package lisper

import (
	"fmt"
	"strconv"
)

const (
	vNil = iota
	vBool
	vInt
	vFloat
	vString
	vList
)

// Value is a generic value
type Value struct {
	V interface{} `json:"v"`
	t int
}

// V creates a new value
func V(x interface{}) Value {

	switch v := x.(type) {

	case nil:
		return Value{nil, vNil}

	case int:
		return Value{int64(v), vInt}

	case int64:
		return Value{int64(v), vInt}

	case float64:
		return Value{v, vFloat}

	case Value:
		return v

	case List:
		return Value{v, vList}

	case string:
		if i, e := strconv.ParseInt(v, 10, 64); e == nil {
			return Value{i, vInt}
		}
		if f, e := strconv.ParseFloat(v, 64); e == nil {
			return Value{f, vFloat}
		}
		return Value{v, vString}

	case bool:
		return Value{v, vBool}

	default:
		panic(fmt.Sprintf("Unsupported Value type [%T]", x))
	}
}
