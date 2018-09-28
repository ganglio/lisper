package lisper

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValue(t *testing.T) {
	i := V(3)
	f := V(3.14)
	s := V("aaa")
	b := V(true)

	assert.Equal(t, i.t, vInt)
	assert.Equal(t, f.t, vFloat)
	assert.Equal(t, s.t, vString)
	assert.Equal(t, b.t, vBool)
}

func TestValuePanic(t *testing.T) {
	defer func() {
		r := recover()
		assert.NotEqual(t, r, nil)
	}()

	_ = V(map[int]string{2: "pesco"})
}
