package converter

import (
	"bytes"
	"fmt"
	temp2 "html/template"
	"io"
	"strings"
	"text/template"
	"time"

	"strconv"

	"github.com/AntonPates/darksky-api-client/services/parser"
)

type converter struct {
}

const pressureMagic = 54
const mmRSinOnePascal = 0.007501

var tempFunc = func(val float64) temp2.HTML {
	val += 0.5
	intVal := int(val)
	var out string
	if intVal >= 0 {
		out = fmt.Sprintf(`<span style="color: red;">%d&#176;</span>`, intVal)
	} else {
		out = fmt.Sprintf(`<span style="color: navy;">%d&#176;</span>`, intVal)
	}
	return temp2.HTML(out)
}

var pressureFunc = func(hectoPascals float64) string {
	val := hectoPascals * 100 * mmRSinOnePascal
	val -= pressureMagic
	intVal := int64(val)
	return strconv.FormatInt(intVal, 10)
}
var humidityFunc = func(val float64) string {
	val *= 100
	intVal := int64(val)
	return strconv.FormatInt(intVal, 10)
}

var dateFunc = func(ts int) string {
	t := time.Unix(int64(ts), 0)
	var month string
	switch t.Month() {
	case time.January:
		month = "Января"
	case time.February:
		month = "Февраля"
	case time.March:
		month = "Марта"
	case time.April:
		month = "Апреля"
	case time.May:
		month = "Мая"
	case time.June:
		month = "Июня"
	case time.July:
		month = "Июля"
	case time.August:
		month = "Августа"
	case time.September:
		month = "Сентября"
	case time.October:
		month = "Октября"
	case time.November:
		month = "Ноября"
	case time.December:
		month = "Декабря"
	default:
		month = t.Month().String()
	}
	out := fmt.Sprintf("%d %s", t.Day(), month)
	return out
}

var timeFunc = func(ts int) string {
	t := time.Unix(int64(ts), 0)
	t = t.Add(time.Hour * 5)
	return t.Format("01.02 15:04")
}

var isFuture = func(ts int) bool {
	t := time.Unix(int64(ts), 0)
	now := time.Now()
	n := int(now.Unix())
	n -= now.Hour()*3600 - now.Minute()*60 - now.Second()
	threshold := time.Unix(int64(n), 0)
	return t.After(threshold)
}

var windDirectionFunc = func(degrees int) (rumb string) {
	switch {
	case (degrees >= 0 && degrees <= 12) || (degrees >= 348 && degrees <= 360):
		rumb = "северный"
	case degrees >= 13 && degrees <= 34:
		rumb = "северный - северо-восточный"
	case degrees >= 35 && degrees <= 56:
		rumb = "северо-восточный"
	case degrees >= 57 && degrees <= 78:
		rumb = "восточный - северо-восточный"
	case degrees >= 79 && degrees <= 102:
		rumb = "восточный"
	case degrees >= 103 && degrees <= 123:
		rumb = "восточный - юго-восточный"
	case degrees >= 124 && degrees <= 147:
		rumb = "юго-восточный"
	case degrees >= 148 && degrees <= 168:
		rumb = "южный - юго-восточный"
	case degrees >= 169 && degrees <= 192:
		rumb = "южный"
	case degrees >= 193 && degrees <= 213:
		rumb = "южный - юго-западный"
	case degrees >= 214 && degrees <= 237:
		rumb = "юго-западный"
	case degrees >= 238 && degrees <= 258:
		rumb = "западный - юго-западный"
	case degrees >= 259 && degrees <= 282:
		rumb = "западный"
	case degrees >= 259 && degrees <= 282:
		rumb = "западный"
	case degrees >= 283 && degrees <= 303:
		rumb = "западный - северо-западный"
	case degrees >= 304 && degrees <= 327:
		rumb = "северо-западный"
	case degrees >= 328 && degrees <= 347:
		rumb = "северный - северо-западный"
	}
	return
}

var windSpeedFunc = func(speed float64) int {
	return int(speed + 1)
}

var iconDescriptionFunc = func(desc string) string {
	desc = strings.Replace(desc, "-", "_", -1)
	return strings.ToUpper(desc)
}

func (c *converter) Summary(ws []*parser.Weather) (*bytes.Buffer, error) {
	funcMap := template.FuncMap{
		"ToLower":         strings.ToLower,
		"TempColor":       tempFunc,
		"Humidity":        humidityFunc,
		"Pressure":        pressureFunc,
		"WindDirection":   windDirectionFunc,
		"WindSpeed":       windSpeedFunc,
		"DateFunc":        dateFunc,
		"IconDescription": iconDescriptionFunc,
		"mod":             func(i, j int) bool { return i%j == 0 },
		"future":          isFuture,
		"TimeFunc":        timeFunc,
	}

	if tpl, err := template.New("summary").Funcs(funcMap).Parse(summaryTpl); err != nil {
		return nil, err
	} else {
		b := make([]byte, 0, 1000)
		buf := bytes.NewBuffer(b)
		if err := tpl.Execute(buf, ws); err != nil {
			return nil, err
		} else {
			return buf, nil
		}
	}
}

func (c *converter) CurrentWeather(ws []*parser.Weather) (io.Reader, error) {
	funcMap := template.FuncMap{
		"ToLower":   strings.ToLower,
		"TempColor": tempFunc,
	}

	if tpl, err := template.New("currently").Funcs(funcMap).Parse(currentlyTpl); err != nil {
		return nil, err
	} else {
		b := make([]byte, 0, 1000)
		buf := bytes.NewBuffer(b)
		if err := tpl.Execute(buf, ws); err != nil {
			return nil, err
		} else {
			return buf, nil
		}
	}
}
