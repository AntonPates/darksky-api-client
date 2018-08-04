package client

import "io"

type Client interface {
	Get(href string) (io.Reader, error)
}

var _ Client = New()

func New() Client {
	return &cli{}
}
