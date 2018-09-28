package lisper

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAtom(t *testing.T) {
	assert.Equal(t, A(V(1)), A(V(1)))
}

func TestListAppend(t *testing.T) {
	l := List{}

	l.Append(V(1), V(2), V(33), A(33))

	assert.Equal(t, l, L(V(1), V(2), V(33), A(33)))
}
