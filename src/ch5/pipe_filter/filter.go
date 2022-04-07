package pipe_filter

// input
type Request interface {
}

// output
type Response interface {
}

// Pipe-Filter structure
type Filter interface {
	Process(data Request) (Response, error)
}
