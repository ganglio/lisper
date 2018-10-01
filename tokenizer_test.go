package lisper

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTokenizer(t *testing.T) {
	assert.Equal(t, Tokenize("(* 2.1 (/ (+ 3.4 4.5) 5.6))"), Tokens{[]string{"*", "2.1", "(", "/", "(", "+", "3.4", "4.5", ")", "5.6", ")"}, 0})
}

func TestTokenizerNext(t *testing.T) {
	to := Tokenize("(1 1 1)")

	for s, l := to.Next(); !l; s, l = to.Next() {
		assert.Equal(t, s, "1")
	}
}
