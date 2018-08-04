package converter

import (
	"fmt"
	"os"
	"testing"

	"io/ioutil"

	"github.com/AntonPates/darksky-api-client/services/parser"
	"github.com/mohae/deepcopy"
	"github.com/stretchr/testify/assert"
)

func TestBriefConvertToHtml(t *testing.T) {
	f, err := os.Open("tmp.json")
	assert.Nil(t, err)
	var p = parser.New()
	w, err := p.Exec(f)
	assert.Nil(t, err)

	ws := make([]*parser.Weather, 5)

	for i := range ws {
		cpy := deepcopy.Copy(w)
		ws[i] = cpy.(*parser.Weather)
		ws[i].Location.Name = "Листвянка"
		ws[i].Location.InName = "в Листвянке"
		ws[i].Location.Anchor = "listvyanka"

	}

	c := converter{}
	out, err := c.CurrentWeather(ws)
	if err != nil {
		fmt.Println(err)
	}
	assert.Nil(t, err)
	b := make([]byte, 0, 1000)
	out.Read(b)

	fmt.Printf("output is \n %s", string(b))

}

func TestSummaryConvertToHtml(t *testing.T) {
	f, err := os.Open("tmp.json")
	assert.Nil(t, err)
	var p = parser.New()
	w, err := p.Exec(f)
	assert.Nil(t, err)

	ws := make([]*parser.Weather, 5)

	for i := range ws {
		cpy := deepcopy.Copy(w)
		ws[i] = cpy.(*parser.Weather)
		ws[i].Location.Name = "Листвянка"
		ws[i].Location.InName = "в Листвянке"
		ws[i].Location.Anchor = "listvyanka"

	}

	c := converter{}
	out, err := c.Summary(ws)
	if err != nil {
		fmt.Println(err)
	}
	assert.Nil(t, err)
	//b := make([]byte, 1000)
	//n, err := out.Read(b)
	//assert.Nil(t, err)
	ioutil.WriteFile("/Users/antonpates/gowork/src/github.com/AntonPates/darksky-api-client/out/weather.html",
		out.Bytes(), 0644)
	//fmt.Printf("output is \n %s", string(out.Bytes()))

}
