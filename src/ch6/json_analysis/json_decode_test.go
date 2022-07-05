package json_analysis

import (
    "encoding/json"
    "fmt"
    "io"
    "log"
    "strings"
    "testing"
)

// This example uses a Decoder to decode a stream of distinct JSON values.
func TestJsonDecoder(t *testing.T) {
    const jsonStream = `
    {"Name": "Ed", "Text": "Knock knock."}
    {"Name": "Sam", "Text": "Who's there?"}
    {"Name": "Ed", "Text": "Go fmt."}
    {"Name": "Sam", "Text": "Go fmt who?"}
    {"Name": "Ed", "Text": "Go fmt yourself!"}
`
    type Message struct {
        Name, Text string
    }
    dec := json.NewDecoder(strings.NewReader(jsonStream))
    for {
        var m Message
        if err := dec.Decode(&m); err == io.EOF {
            break
        } else if err != nil {
            log.Fatal(err)
        }
        fmt.Printf("%s: %s\n", m.Name, m.Text)
    }
}
