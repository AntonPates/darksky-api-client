package parser

import (
	"os"
	"testing"

	"fmt"

	"github.com/stretchr/testify/assert"
)

func TestDecodingOfDarkSkyResponse(t *testing.T) {
	f, err := os.Open("tmp.json")
	assert.Nil(t, err)
	var p parser
	w, err := p.Exec(f)
	assert.Nil(t, err)
	fmt.Printf("%#v", w)
	assert.NotNil(t, w)
	assert.Equal(t, w.Latitude, 52.2896)
	assert.Equal(t, w.Currently.ApparentTemperature, -14.73)
	assert.Equal(t, w.Currently.Icon, "partly-cloudy-day")
}
