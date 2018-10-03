package lisper

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValue(t *testing.T) {
	i := V(3)
	i64 := V(int64(3))
	f := V(3.14)
	s := V(`"aaa"`)
	b := V(true)
	v := V(V(3))
	iS := V("3")
	fS := V("2.3")

	sy := V(`a`)

	assert.Equal(t, i.t, vInt)
	assert.Equal(t, i64.t, vInt)
	assert.Equal(t, iS.t, vInt)
	assert.Equal(t, f.t, vFloat)
	assert.Equal(t, fS.t, vFloat)
	assert.Equal(t, s.t, vString)
	assert.Equal(t, s.V, "aaa")
	assert.Equal(t, b.t, vBool)
	assert.Equal(t, v.t, vInt)
	assert.Equal(t, sy.t, vSymbol)
}

func TestValuePanic(t *testing.T) {
	defer func() {
		assert.Equal(t, recover(), "Unsupported Value type [map[int]string]")
	}()

	_ = V(map[int]string{2: "pesco"})
}
