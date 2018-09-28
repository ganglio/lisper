package lisper

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValOp(t *testing.T) {
	assert.Equal(t, Add(V(1), V(2)), V(3))
	assert.Equal(t, Add(V(1.1), V(2)), V(3.1))
	assert.Equal(t, Add(V(1), V(2.2)), V(3.2))

	assert.Equal(t, Sub(V(1), V(1)), V(0))
	assert.Equal(t, Sub(V(0.9), V(0.9)), V(0.0))
	assert.Equal(t, Sub(V(1), V(2)), V(-1))

	assert.Equal(t, Mul(V(2), V(3)), V(6))
	assert.Equal(t, Eq(Mul(V(2.3), V(3.4)), V(7.82)), TRUE)

	assert.Equal(t, Div(V(4), V(2)), V(2))

	assert.Equal(t, Eq(Div(V(2.3), V(3.4)), V(0.676470588235294)), TRUE)
}

func TestValTest(t *testing.T) {
	assert.Equal(t, Eq(V(1), V(2)), FALSE)
	assert.Equal(t, Eq(V(1), V(2.2)), FALSE)
	assert.Equal(t, Eq(V(1.2), V(1.2)), TRUE)
	assert.Equal(t, Eq(V("string1"), V("string1")), TRUE)

	assert.Equal(t, Gt(V("string"), V(33)), FALSE)
	assert.Equal(t, Gt(V(34), V(33)), TRUE)
	assert.Equal(t, Gt(V("bbb"), V("aaa")), TRUE)
	assert.Equal(t, Ge(V(33), V(33)), TRUE)
	assert.Equal(t, Lt(V(33), V(34)), TRUE)
	assert.Equal(t, Le(V(33), V(33)), TRUE)
	assert.Equal(t, Ne(V(33), V(34)), TRUE)
}

func BenchmarkOp(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_ = Gt(Mul(V(2.1), Div(Add(V(3.4), V(4.5)), V(5.6))), V(5.4))
	}
}

func BenchmarkPlain(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_ = 2.1 * (3.4 + 4.5) / 5.6 //Mul(V(2.1), Div(Add(V(3.4), V(4.5)), V(5.6)))
	}
}
