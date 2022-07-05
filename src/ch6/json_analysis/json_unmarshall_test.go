package json_analysis

import (
    "encoding/json"
    "fmt"
    "testing"
    "time"
)

func TestJsonAnalysis(t *testing.T) {
    type Fruit struct {
        Name     string `json:"name"`
        PriceTag string `json:"priceTag"`
    }

    type FruitBasket struct {
        Name    string           `json:"name"`
        Fruit   map[string]Fruit `json:"fruit"`
        Id      int64            `json:"id"`
        Created time.Time        `json:"created"`
    }
    jsonData := []byte(`
    {
        "Name": "Standard",
        "Fruit" : {
              "1": {
                    "name": "Apple",
                    "priceTag": "$1"
              },
              "2": {
                    "name": "Pear",
                    "priceTag": "$1.5"
              }
        },
        "id": 999,
        "created": "2018-04-09T23:00:00Z"
    }`)

    var basket FruitBasket
    err := json.Unmarshal(jsonData, &basket)
    if err != nil {
        fmt.Println(err)
    }
    for _, item := range basket.Fruit {
        fmt.Println(item.Name, item.PriceTag)
    }
    fmt.Println(jsonData)
}

func TestUnknownJson(t *testing.T) {
    jsonData := []byte(`{"Name":"Eve","Age":6,"Parents":["Alice","Bob"]}`)

    var v interface{}
    json.Unmarshal(jsonData, &v)
    data := v.(map[string]interface{})

    for k, v := range data {
        switch v := v.(type) {
        case string:
            fmt.Println(k, v, "(string)")
        case float64:
            fmt.Println(k, v, "(float64)")
        case []interface{}:
            fmt.Println(k, "(array):")
            for i, u := range v {
                fmt.Println("    ", i, u)
            }
        default:
            fmt.Println(k, v, "(unknown)")
        }
    }
}
