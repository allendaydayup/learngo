package condition

import (
	"fmt"
	"github.com/easierway/concurrent_map"
	"go_learning/src/ch5/series"
	"testing"
)

func TestSwitchMultiCase(t *testing.T) {
	for i := 0; i < 5; i++ {
		switch {
		case 0 <= i && i <= 2:
			t.Log("even")
		case i <= 4:
			t.Log("Odd")
		default:
			t.Log("it is not 0-3")
		}
	}

	a := series.Max(1, 20)
	fmt.Println(a)

	m := concurrent_map.CreateConcurrentMap(99)
	m.Set(concurrent_map.StrKey("key"), 10)
	t.Log(m.Get(concurrent_map.StrKey("key")))
}
