package lisper

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEval(t *testing.T) {
	assert.Equal(t, Eval("(+ 1 (* 3 5))"), V(16))
	assert.Equal(t, Eval("(+ 1 2)"), V(3))
	assert.Equal(t, Eval("(+ (+ 1 2) (+ 3 4))"), V(10))
}

func TestBuiltin(t *testing.T) {
	assert.Equal(t, Eval("(gt? 4 3)"), TRUE)
	assert.Equal(t, Eval("(if (gt? 4 3) (* 3 5) (+ 2 6)"), V(15))
	assert.Equal(t, Eval(`((define a 3) (define b 3) (eq? a b))`), TRUE)
	assert.Equal(t, Eval(`( (define a 1) (define a 3) a)`), V(3))
	assert.Equal(t, Eq(Eval(`((define a 1) (* (* a a) PI))`), Sym[PI]), TRUE)
	assert.Equal(t, Eval(`(zz)`), V("zz"))
}
