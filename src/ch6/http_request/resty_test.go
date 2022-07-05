package http_request

import (
    "fmt"
    "github.com/go-resty/resty/v2"
    "testing"
)

type Library struct {
    Name   string
    Latest string
}

type Libraries struct {
    Results []*Library
}

func TestResty(t *testing.T) {
    client := resty.New()

    libraries := &Libraries{}
    client.R().SetResult(libraries).Get("https://api.cdnjs.com/libraries")
    fmt.Printf("%d libraries\n", len(libraries.Results))

    for _, lib := range libraries.Results {
        fmt.Println("first library:")
        fmt.Printf("name:%s latest:%s\n", lib.Name, lib.Latest)
        break
    }
}
