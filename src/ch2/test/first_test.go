package test

import (
    "testing"
)

func TestFirstTry(t *testing.T) {
    a := 1
    b := 1
    for i := 0; i < 5; i++ {
        t.Log(" ", b)
        tmp := a
        a = b
        b = tmp + a
    }
}
