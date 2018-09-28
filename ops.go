package lisper

// Add calculates the sum between two values
func Add(a Value, b Value) (r Value) {
	if a.t == vString || b.t == vString {
		panic("Unsupported operands")
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
func Sub(a Value, b Value) (r Value) {
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
func Mul(a Value, b Value) (r Value) {

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
func Div(a Value, b Value) (r Value) {

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
func Eq(a Value, b Value) Value {
	if a.t != b.t {
		return FALSE
	}

	if a.t == vFloat {
		return V(floatEquals(a.V.(float64), b.V.(float64)))
	}

	return V(a.V == b.V)
}

// Gt checks a > b
func Gt(a Value, b Value) Value {
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
func Ge(a Value, b Value) Value {
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
func Lt(a Value, b Value) Value {
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
func Le(a Value, b Value) Value {
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
func Ne(a Value, b Value) Value {
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
