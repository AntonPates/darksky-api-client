package parser

import "github.com/AntonPates/darksky-api-client/services/config"

type Weather struct {
	Location  config.Location
	Currently HourlyData `json:"currently"`
	Daily     struct {
		Data    []DailyData `json:"data"`
		Icon    string      `json:"icon"`
		Summary string      `json:"summary"`
	} `json:"daily"`
	Flags struct {
		Isd_stations []string `json:"isd-stations"`
		Sources      []string `json:"sources"`
		Units        string   `json:"units"`
	} `json:"flags"`
	Hourly struct {
		Data    []HourlyData `json:"data"`
		Icon    string       `json:"icon"`
		Summary string       `json:"summary"`
	} `json:"hourly"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Offset    int     `json:"offset"`
	Timezone  string  `json:"timezone"`
}

type DailyData struct {
	ApparentTemperatureHigh     float64 `json:"apparentTemperatureHigh"`
	ApparentTemperatureHighTime int     `json:"apparentTemperatureHighTime"`
	ApparentTemperatureLow      float64 `json:"apparentTemperatureLow"`
	ApparentTemperatureLowTime  int     `json:"apparentTemperatureLowTime"`
	ApparentTemperatureMax      float64 `json:"apparentTemperatureMax"`
	ApparentTemperatureMaxTime  int     `json:"apparentTemperatureMaxTime"`
	ApparentTemperatureMin      float64 `json:"apparentTemperatureMin"`
	ApparentTemperatureMinTime  int     `json:"apparentTemperatureMinTime"`
	Icon                        string  `json:"icon"`
	MoonPhase                   float64 `json:"moonPhase"`
	PrecipIntensityMax          float64 `json:"precipIntensityMax"`
	PrecipIntensityMaxTime      int     `json:"precipIntensityMaxTime"`
	SunriseTime                 int     `json:"sunriseTime"`
	SunsetTime                  int     `json:"sunsetTime"`
	TemperatureHigh             float64 `json:"temperatureHigh"`
	TemperatureHighTime         int     `json:"temperatureHighTime"`
	TemperatureLow              float64 `json:"temperatureLow"`
	TemperatureLowTime          int     `json:"temperatureLowTime"`
	TemperatureMax              float64 `json:"temperatureMax"`
	TemperatureMaxTime          int     `json:"temperatureMaxTime"`
	TemperatureMin              float64 `json:"temperatureMin"`
	TemperatureMinTime          int     `json:"temperatureMinTime"`
	UvIndexTime                 int     `json:"uvIndexTime"`
	WindGustTime                int     `json:"windGustTime"`

	HourlyData
}

type HourlyData struct {
	ApparentTemperature float64 `json:"apparentTemperature"`
	CloudCover          float64 `json:"cloudCover"`
	DewPoint            float64 `json:"dewPoint"`
	Humidity            float64 `json:"humidity"`
	Icon                string  `json:"icon"`
	Ozone               float64 `json:"ozone"`
	PrecipAccumulation  float64 `json:"precipAccumulation"`
	PrecipIntensity     float64 `json:"precipIntensity"`
	PrecipProbability   float64 `json:"precipProbability"`
	PrecipType          string  `json:"precipType"`
	Pressure            float64 `json:"pressure"`
	Summary             string  `json:"summary"`
	Temperature         float64 `json:"temperature"`
	Time                int     `json:"time"`
	UvIndex             int     `json:"uvIndex"`
	Visibility          float64 `json:"visibility"`
	WindBearing         int     `json:"windBearing"`
	WindGust            float64 `json:"windGust"`
	WindSpeed           float64 `json:"windSpeed"`
}
