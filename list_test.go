package lisper

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestListAppend(t *testing.T) {
	l := List{}

	l.Append(V(1), V(2), V(33))

	assert.Equal(t, l, L(V(1), V(2), V(33)))
}
