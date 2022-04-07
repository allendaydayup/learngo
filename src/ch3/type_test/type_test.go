package type_test

import (
	"math"
	"testing"
)

type MyInt int64

func TestImplicit(t *testing.T) {
	var a int = 1
	var b int64
	b = int64(a)
	var c MyInt
	c = MyInt(b)
	d := math.MaxInt64
	e := math.MinInt64
	t.Log(a, b, c, d, e)
}
