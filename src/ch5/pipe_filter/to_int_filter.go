package pipe_filter

import (
    "errors"
    "strconv"
)

var ToIntFilterWrongFormatError = errors.New("input data should be []string")

type ToInFilter struct {
}

func NewToInFilter() *ToInFilter {
    return &ToInFilter{}
}

func (tif *ToInFilter) Process(data Request) (Response, error) {
    parts, ok := data.([]string)
    if !ok {
        return nil, ToIntFilterWrongFormatError
    }
    ret := []int{}
    for _, part := range parts {
        s, err := strconv.Atoi(part)
        if err != nil {
            return nil, err
        }
        ret = append(ret, s)
    }
    return ret, nil
}
