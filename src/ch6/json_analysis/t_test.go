package json_analysis

import (
    "encoding/json"
    "fmt"
    "testing"
)

type Person struct {
    Friends []string
}

func TestA(t *testing.T) {
    var f1 []string
    f2 := make([]string, 0)

    json1, _ := json.Marshal(Person{f1})
    json2, _ := json.Marshal(Person{f2})

    fmt.Printf("%s\n", json1)
    fmt.Printf("%s\n", json2)
}
