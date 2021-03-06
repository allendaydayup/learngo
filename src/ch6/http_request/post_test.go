package http_request

import (
    "bytes"
    "fmt"
    "io/ioutil"
    "net/http"
    "reflect"
    "strings"
    "testing"
)

func TestPost(t *testing.T) {
    resp, err := http.Post("http://httpbin.org/post",
        "application/x-www-form-urlencoded",
        strings.NewReader("name=Detector&mobile=1xxxxxxxx"))
    if err != nil {
        fmt.Println(err)
        return
    }

    defer resp.Body.Close()
    headers := resp.Header
    // headers 打印报文头部信息
    for k, v := range headers {
        fmt.Printf("%v, %v\n", k, v) // %v 打印interfac{}的值
    }

    // 打印响应信息内容
    fmt.Printf("响应状态：%s,响应码： %d\n", resp.Status, resp.StatusCode)
    fmt.Printf("协议：%s\n", resp.Proto)
    fmt.Printf("响应内容长度： %d\n", resp.ContentLength)
    fmt.Printf("编码格式：%v\n", resp.TransferEncoding) // 未指定时为空
    fmt.Printf("是否压缩：%t\n", resp.Uncompressed)
    fmt.Println(reflect.TypeOf(resp.Body)) // *http.gzipReader
    fmt.Println(resp.Body)

    buf := bytes.NewBuffer(make([]byte, 0, 512))
    length, _ := buf.ReadFrom(resp.Body)
    fmt.Println(len(buf.Bytes()))
    fmt.Println(length)
    fmt.Println(string(buf.Bytes()))
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        fmt.Println(err)
        return
    }
    fmt.Println(string(body))
}
