package lisper

// Add calculates the sum between two values
func Add(a Value, b Value) (r Value) {
	if a.t == v_string || b.t == v_string {
		panic("Unsupported operands")
	}

	switch a.t {
	case v_int:
		switch b.t {
		case v_int:
			r = V(a.v.(int64) + b.v.(int64))
		case v_float:
			r = V(float64(a.v.(int64)) + b.v.(float64))
		}
	case v_float:
		switch b.t {
		case v_int:
			r = V(a.v.(float64) + float64(b.v.(int64)))
		case v_float:
			r = V(a.v.(float64) + b.v.(float64))
		}
	}

	return
}

// Sub calculates the difference between two values
func Sub(a Value, b Value) (r Value) {
	if a.t == v_string || b.t == v_string {
		panic("Unsupported operands")
	}

	switch a.t {
	case v_int:
		switch b.t {
		case v_int:
			r = V(a.v.(int64) - b.v.(int64))
		case v_float:
			r = V(float64(a.v.(int64)) - b.v.(float64))
		}
	case v_float:
		switch b.t {
		case v_int:
			r = V(a.v.(float64) - float64(b.v.(int64)))
		case v_float:
			r = V(a.v.(float64) - b.v.(float64))
		}
	}

	return
}

// Mul calculates the product between two values
func Mul(a Value, b Value) (r Value) {

	if a.t == v_string || b.t == v_string {
		panic("Unsupported operands")
	}

	switch a.t {
	case v_int:
		switch b.t {
		case v_int:
			r = V(a.v.(int64) * b.v.(int64))
		case v_float:
			r = V(float64(a.v.(int64)) * b.v.(float64))
		}
	case v_float:
		switch b.t {
		case v_int:
			r = V(a.v.(float64) * float64(b.v.(int64)))
		case v_float:
			r = V(a.v.(float64) * b.v.(float64))
		}
	}

	return
}

// Div calculates the division between two values
func Div(a Value, b Value) (r Value) {

	if a.t == v_string || b.t == v_string {
		panic("Unsupported operands")
	}

	if b.t == v_int && b.v.(int64) == 0 {
		panic("Divide by zero")
	}

	if b.t == v_float && floatEquals(b.v.(float64), 0) {
		panic("Divide by zero")
	}

	switch a.t {
	case v_int:
		switch b.t {
		case v_int:
			r = V(a.v.(int64) / b.v.(int64))
		case v_float:
			r = V(float64(a.v.(int64)) / b.v.(float64))
		}
	case v_float:
		switch b.t {
		case v_int:
			r = V(a.v.(float64) / float64(b.v.(int64)))
		case v_float:
			r = V(a.v.(float64) / b.v.(float64))
		}
	}

	return
}

// Eq checks a == b
func Eq(a Value, b Value) Value {
	if a.t != b.t {
		return FALSE
	}

	if a.t == v_float {
		return V(floatEquals(a.v.(float64), b.v.(float64)))
	}

	return V(a.v == b.v)
}

// Gt checks a > b
func Gt(a Value, b Value) Value {
	if a.t != b.t && (a.t == v_string || b.t == v_string) {
		return FALSE
	}

	if a.t == v_string && b.t == v_string {
		return V(a.v.(string) > b.v.(string))
	}

	var x, y float64

	if a.t == v_float {
		x = a.v.(float64)
	} else {
		x = float64(a.v.(int64))
	}

	if b.t == v_float {
		y = b.v.(float64)
	} else {
		y = float64(b.v.(int64))
	}

	return V(x > y)
}

// Ge checks a >= b
func Ge(a Value, b Value) Value {
	if a.t != b.t && (a.t == v_string || b.t == v_string) {
		return FALSE
	}

	if a.t == v_string && b.t == v_string {
		return V(a.v.(string) >= b.v.(string))
	}

	var x, y float64

	if a.t == v_float {
		x = a.v.(float64)
	} else {
		x = float64(a.v.(int64))
	}

	if b.t == v_float {
		y = b.v.(float64)
	} else {
		y = float64(b.v.(int64))
	}

	return V(x >= y)
}

// Lt checks a < b
func Lt(a Value, b Value) Value {
	if a.t != b.t && (a.t == v_string || b.t == v_string) {
		return FALSE
	}

	if a.t == v_string && b.t == v_string {
		return V(a.v.(string) < b.v.(string))
	}

	var x, y float64

	if a.t == v_float {
		x = a.v.(float64)
	} else {
		x = float64(a.v.(int64))
	}

	if b.t == v_float {
		y = b.v.(float64)
	} else {
		y = float64(b.v.(int64))
	}

	return V(x < y)
}

// Le checks a <= b
func Le(a Value, b Value) Value {
	if a.t != b.t && (a.t == v_string || b.t == v_string) {
		return FALSE
	}

	if a.t == v_string && b.t == v_string {
		return V(a.v.(string) <= b.v.(string))
	}

	var x, y float64

	if a.t == v_float {
		x = a.v.(float64)
	} else {
		x = float64(a.v.(int64))
	}

	if b.t == v_float {
		y = b.v.(float64)
	} else {
		y = float64(b.v.(int64))
	}

	return V(x <= y)
}

// Ne checks a != b
func Ne(a Value, b Value) Value {
	if a.t != b.t {
		return TRUE
	}

	if a.t == v_string && b.t == v_string {
		return V(a.v.(string) != b.v.(string))
	}

	var x, y float64

	if a.t == v_float {
		x = a.v.(float64)
	} else {
		x = float64(a.v.(int64))
	}

	if b.t == v_float {
		y = b.v.(float64)
	} else {
		y = float64(b.v.(int64))
	}

	return V(!floatEquals(x, y))
}
