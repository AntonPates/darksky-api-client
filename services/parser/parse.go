package parser

import "io"

type Parser interface {
	Exec(reader io.Reader) (*Weather, error)
}

var _ Parser = New()

func New() Parser {
	return &parser{}
}
