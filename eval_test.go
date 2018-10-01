package lisper

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEval(t *testing.T) {
	assert.Equal(t, Eval("(+ 1 (* 3 5))"), V(16))
	assert.Equal(t, Eval("(+ 1 2)"), V(3))
}

// func TestEvalMulti(t *testing.T) {
// 	assert.Equal(t, Eval("(+ 1 2 3 4 5)"), V(15))
// }
