package _select

import (
    "fmt"
    "testing"
    "time"
)

func service() string {
    time.Sleep(time.Millisecond * 50)
    return "Done"
}

func AsyncService() chan string {
    //retCh := make(chan string)
    retCh := make(chan string, 1)
    go func() {
        ret := service()
        fmt.Println("return result.")
        retCh <- ret
        fmt.Println("servie exited.")
    }()
    return retCh
}
func TestSelect(t *testing.T) {
    select {
    case ret := <-AsyncService():
        t.Log(ret)
    case <-time.After(time.Millisecond * 10):
        t.Error("time out")
    }
}
