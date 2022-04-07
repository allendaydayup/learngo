package main

import (
    "fmt"
    "os"
)

func main() {
    if len(os.Args) > 1 {
        fmt.Println("Hello World!", os.Args[2])
    } else {
        fmt.Println("没有参数")
    }
}
