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

	assert.Equal(t, Add(V("a"), V("b")), V("ab"))
}

func TestOpAddFewOps(t *testing.T) {
	defer func() {
		r := recover()
		assert.NotEqual(t, r, nil)
	}()

	_ = Add(V(1))
}

func TestOpSub(t *testing.T) {
	defer func() {
		r := recover()
		assert.NotEqual(t, r, nil)
	}()

	assert.Equal(t, Sub(V(1), V(1)), V(0))
	assert.Equal(t, Sub(V(0.9), V(0.9)), V(0.0))
	assert.Equal(t, Sub(V(1), V(2)), V(-1))

	assert.Equal(t, Eq(Sub(V(1), V(1.1)), V(-0.1)), TRUE)
	assert.Equal(t, Eq(Sub(V(1.1), V(1)), V(0.1)), TRUE)

	_ = Sub(V("a"), V("b"))
}

func TestOpSubFewOps(t *testing.T) {
	defer func() {
		r := recover()
		assert.NotEqual(t, r, nil)
	}()

	_ = Sub(V(1))
}

func TestOpMul(t *testing.T) {
	defer func() {
		r := recover()
		assert.NotEqual(t, r, nil)
	}()

	assert.Equal(t, Mul(V(2), V(3)), V(6))
	assert.Equal(t, Eq(Mul(V(2.3), V(3.4)), V(7.82)), TRUE)
	assert.Equal(t, Eq(Mul(V(2), V(3.4)), V(6.8)), TRUE)
	assert.Equal(t, Eq(Mul(V(2.3), V(3)), V(6.9)), TRUE)

	_ = Mul(V("a"), V("b"))
}

func TestOpMulFewOps(t *testing.T) {
	defer func() {
		r := recover()
		assert.NotEqual(t, r, nil)
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
		r := recover()
		assert.NotEqual(t, r, nil)
	}()
	_ = Div(V(1), V(0))
}

func TestValDivByZeroFloat(t *testing.T) {
	defer func() {
		r := recover()
		assert.NotEqual(t, r, nil)
	}()
	_ = Div(V(1), V(0.0))
}

func TestValDivStrings(t *testing.T) {
	defer func() {
		r := recover()
		assert.NotEqual(t, r, nil)
	}()

	_ = Div(V("a"), V("b"))
}

func TestOpDivFewOps(t *testing.T) {
	defer func() {
		r := recover()
		assert.NotEqual(t, r, nil)
	}()

	_ = Div(V(1))
}

func TestOpEq(t *testing.T) {
	defer func() {
		r := recover()
		assert.NotEqual(t, r, nil)
	}()

	assert.Equal(t, Eq(V(1), V(2)), FALSE)
	assert.Equal(t, Eq(V(1), V(2.2)), FALSE)
	assert.Equal(t, Eq(V(1.2), V(1.2)), TRUE)
	assert.Equal(t, Eq(V("string1"), V("string1")), TRUE)

	_ = Eq(V(1))
}

func TestOpNe(t *testing.T) {
	defer func() {
		r := recover()
		assert.NotEqual(t, r, nil)
	}()

	assert.Equal(t, Ne(V(33), V(34)), TRUE)
	assert.Equal(t, Ne(V(33), V(34.3)), TRUE)
	assert.Equal(t, Ne(V("a"), V("b")), TRUE)
	assert.Equal(t, Ne(V(33.3), V(34.3)), TRUE)
	assert.Equal(t, Ne(V(33.3), V(34.3)), TRUE)

	_ = Ne(V(1))
}

func TestOpGt(t *testing.T) {
	defer func() {
		r := recover()
		assert.NotEqual(t, r, nil)
	}()

	assert.Equal(t, Gt(V("string"), V(33)), FALSE)
	assert.Equal(t, Gt(V(34), V(33)), TRUE)
	assert.Equal(t, Gt(V(34.3), V(33)), TRUE)
	assert.Equal(t, Gt(V(34), V(33.4)), TRUE)
	assert.Equal(t, Gt(V("bbb"), V("aaa")), TRUE)

	_ = Gt(V(1))
}

func TestOpGe(t *testing.T) {
	defer func() {
		r := recover()
		assert.NotEqual(t, r, nil)
	}()

	assert.Equal(t, Ge(V(33), V(33)), TRUE)
	assert.Equal(t, Ge(V(34), V(33)), TRUE)
	assert.Equal(t, Ge(V(34.3), V(34)), TRUE)
	assert.Equal(t, Ge(V(34), V(33.3)), TRUE)
	assert.Equal(t, Ge(V(34.3), V(33.3)), TRUE)
	assert.Equal(t, Ge(V("bbb"), V("aaa")), TRUE)
	assert.Equal(t, Ge(V(33), V("bbb")), FALSE)

	_ = Ge(V(1))
}

func TestOpLt(t *testing.T) {
	defer func() {
		r := recover()
		assert.NotEqual(t, r, nil)
	}()

	assert.Equal(t, Lt(V(33), V(34)), TRUE)
	assert.Equal(t, Lt(V(33.3), V(34)), TRUE)
	assert.Equal(t, Lt(V(33), V(34.3)), TRUE)
	assert.Equal(t, Lt(V(33.3), V(34.3)), TRUE)
	assert.Equal(t, Lt(V("aaa"), V("bbb")), TRUE)
	assert.Equal(t, Lt(V(33), V("bbb")), FALSE)

	_ = Lt(V(1))
}

func TestOpLe(t *testing.T) {
	defer func() {
		r := recover()
		assert.NotEqual(t, r, nil)
	}()

	assert.Equal(t, Le(V(33), V(33)), TRUE)
	assert.Equal(t, Le(V(33), V(34)), TRUE)
	assert.Equal(t, Le(V(33.3), V(34)), TRUE)
	assert.Equal(t, Le(V(33), V(34.3)), TRUE)
	assert.Equal(t, Le(V(33.3), V(34.3)), TRUE)
	assert.Equal(t, Le(V("aaa"), V("bbb")), TRUE)
	assert.Equal(t, Le(V(33), V("bbb")), FALSE)

	_ = Le(V(1))
}
