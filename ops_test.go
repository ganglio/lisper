package lisper

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOpAdd(t *testing.T) {
	assert.Equal(t, Add(V(1), V(2)), V(3))
	assert.Equal(t, Add(V(1.1), V(2)), V(3.1))
	assert.Equal(t, Add(V(1), V(2.2)), V(3.2))
	assert.Equal(t, Eq(Add(V(1.1), V(2.2)), V(3.3)), TRUE)

	assert.Equal(t, Add(V(`"a"`), V(`"b"`)), V(`"ab"`))
}

func TestOpAddFewOps(t *testing.T) {
	defer func() {
		assert.Equal(t, recover(), "Invalid number of operands")
	}()

	_ = Add(V(1))
}

func TestOpSub(t *testing.T) {
	defer func() {
		assert.Equal(t, recover(), "Unsupported operands")
	}()

	assert.Equal(t, Sub(V(1), V(1)), V(0))
	assert.Equal(t, Sub(V(0.9), V(0.9)), V(0.0))
	assert.Equal(t, Sub(V(1), V(2)), V(-1))

	assert.Equal(t, Eq(Sub(V(1), V(1.1)), V(-0.1)), TRUE)
	assert.Equal(t, Eq(Sub(V(1.1), V(1)), V(0.1)), TRUE)

	_ = Sub(V(`"a"`), V(`"b"`))
}

func TestOpSubFewOps(t *testing.T) {
	defer func() {
		assert.Equal(t, recover(), "Invalid number of operands")
	}()

	_ = Sub(V(1))
}

func TestOpMul(t *testing.T) {
	defer func() {
		assert.Equal(t, recover(), "Unsupported operands")
	}()

	assert.Equal(t, Mul(V(2), V(3)), V(6))
	assert.Equal(t, Eq(Mul(V(2.3), V(3.4)), V(7.82)), TRUE)
	assert.Equal(t, Eq(Mul(V(2), V(3.4)), V(6.8)), TRUE)
	assert.Equal(t, Eq(Mul(V(2.3), V(3)), V(6.9)), TRUE)

	_ = Mul(V(`"a"`), V(`"b"`))
}

func TestOpMulFewOps(t *testing.T) {
	defer func() {
		assert.Equal(t, recover(), "Invalid number of operands")
	}()

	_ = Mul(V(1))
}

func TestOpDiv(t *testing.T) {
	assert.Equal(t, Div(V(4), V(2)), V(2))
	assert.Equal(t, Eq(Div(V(2.3), V(3.4)), V(0.676470588235294)), TRUE)
	assert.Equal(t, Eq(Div(V(4.2), V(2)), V(2.1)), TRUE)
	assert.Equal(t, Eq(Div(V(4), V(2.2)), V(1.8181818181818181)), TRUE)
}

func TestValDivByZeroInt(t *testing.T) {
	defer func() {
		assert.Equal(t, recover(), "Divide by zero")
	}()

	_ = Div(V(1), V(0))
}

func TestValDivByZeroFloat(t *testing.T) {
	defer func() {
		assert.Equal(t, recover(), "Divide by zero")
	}()

	_ = Div(V(1), V(0.0))
}

func TestValDivStrings(t *testing.T) {
	defer func() {
		assert.Equal(t, recover(), "Unsupported operands")
	}()

	_ = Div(V(`"a"`), V(`"b"`))
}

func TestOpDivFewOps(t *testing.T) {
	defer func() {
		assert.Equal(t, recover(), "Invalid number of operands")
	}()

	_ = Div(V(1))
}

func TestOpEq(t *testing.T) {
	defer func() {
		assert.Equal(t, recover(), "Invalid number of operands")
	}()

	assert.Equal(t, Eq(V(1), V(2)), FALSE)
	assert.Equal(t, Eq(V(1), V(2.2)), FALSE)
	assert.Equal(t, Eq(V(1.2), V(1.2)), TRUE)
	assert.Equal(t, Eq(V(`"string1"`), V(`"string1"`)), TRUE)
	assert.Equal(t, Eq(V("s1"), V("s1")), TRUE)

	_ = Eq(V(1))
}

func TestOpNe(t *testing.T) {
	defer func() {
		assert.Equal(t, recover(), "Invalid number of operands")
	}()

	assert.Equal(t, Ne(V(33), V(34)), TRUE)
	assert.Equal(t, Ne(V(33), V(34.3)), TRUE)
	assert.Equal(t, Ne(V(`"a"`), V(`"b"`)), TRUE)
	assert.Equal(t, Ne(V(33.3), V(34.3)), TRUE)
	assert.Equal(t, Ne(V(33.3), V(34.3)), TRUE)

	_ = Ne(V(1))
}

func TestOpGt(t *testing.T) {
	defer func() {
		assert.Equal(t, recover(), "Invalid number of operands")
	}()

	assert.Equal(t, Gt(V(`"string"`), V(33)), FALSE)
	assert.Equal(t, Gt(V(34), V(33)), TRUE)
	assert.Equal(t, Gt(V(34.3), V(33)), TRUE)
	assert.Equal(t, Gt(V(34), V(33.4)), TRUE)
	assert.Equal(t, Gt(V(`"bbb"`), V(`"aaa"`)), TRUE)

	_ = Gt(V(1))
}

func TestOpGe(t *testing.T) {
	defer func() {
		assert.Equal(t, recover(), "Invalid number of operands")
	}()

	assert.Equal(t, Ge(V(33), V(33)), TRUE)
	assert.Equal(t, Ge(V(34), V(33)), TRUE)
	assert.Equal(t, Ge(V(34.3), V(34)), TRUE)
	assert.Equal(t, Ge(V(34), V(33.3)), TRUE)
	assert.Equal(t, Ge(V(34.3), V(33.3)), TRUE)
	assert.Equal(t, Ge(V(`"bbb"`), V(`"aaa"`)), TRUE)
	assert.Equal(t, Ge(V(33), V(`"bbb"`)), FALSE)

	_ = Ge(V(1))
}

func TestOpLt(t *testing.T) {
	defer func() {
		assert.Equal(t, recover(), "Invalid number of operands")
	}()

	assert.Equal(t, Lt(V(33), V(34)), TRUE)
	assert.Equal(t, Lt(V(33.3), V(34)), TRUE)
	assert.Equal(t, Lt(V(33), V(34.3)), TRUE)
	assert.Equal(t, Lt(V(33.3), V(34.3)), TRUE)
	assert.Equal(t, Lt(V(`"aaa"`), V(`"bbb"`)), TRUE)
	assert.Equal(t, Lt(V(33), V(`"bbb"`)), FALSE)

	_ = Lt(V(1))
}

func TestOpLe(t *testing.T) {
	defer func() {
		assert.Equal(t, recover(), "Invalid number of operands")
	}()

	assert.Equal(t, Le(V(33), V(33)), TRUE)
	assert.Equal(t, Le(V(33), V(34)), TRUE)
	assert.Equal(t, Le(V(33.3), V(34)), TRUE)
	assert.Equal(t, Le(V(33), V(34.3)), TRUE)
	assert.Equal(t, Le(V(33.3), V(34.3)), TRUE)
	assert.Equal(t, Le(V(`"aaa"`), V(`"bbb"`)), TRUE)
	assert.Equal(t, Le(V(33), V(`"bbb"`)), FALSE)

	_ = Le(V(1))
}

func TestFirst(t *testing.T) {
	defer func() {
		assert.Equal(t, recover(), "Invalid number of operands")
	}()

	assert.Equal(t, First(V(L(V(1), V(2)))), V(1))
	assert.Equal(t, First(V(L())), NIL)

	_ = First(
		V(L(V(1))),
		V(L(V(2))),
	)
}

func TestFirstNotList(t *testing.T) {
	defer func() {
		assert.Equal(t, recover(), "Argument must be a list")
	}()

	_ = First(V(1))
}

func TestRest(t *testing.T) {
	defer func() {
		assert.Equal(t, recover(), "Invalid number of operands")
	}()

	assert.Equal(t, Rest(V(L(V(1), V(2)))), V(L(V(2))))
	assert.Equal(t, Rest(V(L(V(1)))), NIL)

	_ = Rest(V(L(V(1))), V(L(V(2))))
}

func TestRestNotList(t *testing.T) {
	defer func() {
		assert.Equal(t, recover(), "Argument must be a list")
	}()

	_ = Rest(V(1))
}

func TestIf(t *testing.T) {
	defer func() {
		assert.Equal(t, recover(), "Invalid number of operands")
	}()

	assert.Equal(t, If(TRUE, V(1), V(2)), V(1))
	assert.Equal(t, If(FALSE, V(1), V(2)), V(2))
	assert.Equal(t, If(FALSE, V(1)), NIL)

	_ = If(TRUE, V(1), V(2), V(3))
}

func TestDefine(t *testing.T) {
	_ = Define(V("r"), V(33))
	assert.Equal(t, Sym[V("r")], V(33))

	_ = Define(V("r"), V(34))
	assert.Equal(t, Sym[V("r")], V(34))
}

func TestDefineArgCount(t *testing.T) {
	defer func() {
		assert.Equal(t, recover(), "Invalid number of operands")
	}()

	_ = Define(V("a"))
}

func TestDefineOpType(t *testing.T) {
	defer func() {
		assert.Equal(t, recover(), "Invalid operand type")
	}()

	_ = Define(V(3), V(4))
}
