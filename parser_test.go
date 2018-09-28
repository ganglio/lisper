package lisper

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParserSimple(t *testing.T) {

	to := Tokenize("(+ 1 2)")
	p1 := Parse(&to)
	p2 := L(PLUS, V(1), V(2))

	assert.Equal(t, p1, p2)
}

func TestParserComplex(t *testing.T) {

	to := Tokenize("(* 2.1 (/ (+ 3.4 4.5) 5.6))")
	p1 := Parse(&to)
	p2 := L(STAR, V(2.1), L(SLASH, L(PLUS, V(3.4), V(4.5)), V(5.6)))

	assert.Equal(t, p1, p2)
}

func TestParserPanic(t *testing.T) {
	defer func() {
		r := recover()
		assert.NotEqual(t, r, nil)
	}()

	_ = Parse(&Tokens{[]string{}, 0})
}
