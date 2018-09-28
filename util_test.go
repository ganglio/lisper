package lisper

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTokenizer(t *testing.T) {
	assert.Equal(t, tokenizer("(* 2.1 (/ (+ 3.4 4.5) 5.6))"), []string{"(", "*", "2.1", "(", "/", "(", "+", "3.4", "4.5", ")", "5.6", ")", ")"})
}

func TestParser(t *testing.T) {

	to := tokenizer("(* 2.1 (/ (+ 3.4 4.5) 5.6))")
	p1 := parser(to)
	p2 := L(
		STAR,
		V(2.1),
		L(
			SLASH,
			L(
				PLUS,
				V(3.4),
				V(4.5),
			),
			V(5.6),
		),
	)

	assert.Equal(t, p1, p2)
}
