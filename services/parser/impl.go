package parser

import (
	"encoding/json"
	"io"
)

type parser struct{}

func (p *parser) Exec(r io.Reader) (*Weather, error) {
	var w Weather
	dec := json.NewDecoder(r)
	if err := dec.Decode(&w); err != nil {
		return nil, err
	} else {
		return &w, nil
	}
}
