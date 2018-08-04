package converter

import (
	"io"

	"bytes"

	"github.com/AntonPates/darksky-api-client/services/parser"
)

type Converter interface {
	Summary(ws []*parser.Weather) (*bytes.Buffer, error)
	CurrentWeather(ws []*parser.Weather) (io.Reader, error)
}

var _ Converter = New()

func New() Converter {
	return &converter{}
}
