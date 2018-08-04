package config

import (
	"fmt"
	"net/url"
	"os"
	"path"
	"sort"

	"github.com/mohae/deepcopy"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type conf struct {
	snake     *viper.Viper
	locations []Location
}

const (
	endpointKey = "endpoint"
	apiKey      = "api_key"
	langKey     = "lang"
	langValue   = "ru"
	unitsKey    = "units"
	unitsValue  = "si"
)

func (c *conf) Load(fpath string) error {
	if _, err := os.Open(fpath); err != nil {
		return err
	} else {
		c.snake = viper.New()
		c.snake.SetConfigFile(fpath)

		if err := c.snake.ReadInConfig(); err != nil {
			return err
		} else {
			defer c.loadLocations()
			c.snake.WatchConfig()
			return nil
		}
	}
}

func (c *conf) loadLocations() {
	locations := c.snake.GetStringMap(locationsKey)
	c.locations = make([]Location, 0, len(locations))
	for k, v := range locations {
		if cv, ok := v.(map[string]interface{}); !ok {
			log.Errorf("%s conf is corrupted", k)
			continue
		} else {
			loc := Location{}
			if err := loc.Fill(cv); err != nil {
				log.Error(err)
				continue
			} else {
				c.locations = append(c.locations, loc)
			}
		}
	}

	if base, err := url.Parse(c.snake.GetString(endpointKey)); err != nil {
		log.Error(err)
	} else {
		base.Path = path.Join(base.Path, c.snake.GetString(apiKey))
		log.Info(c.snake.GetString(endpointKey))
		for i, loc := range c.locations {
			var locURL url.URL
			locURL = *base
			locURL.Path = path.Join(locURL.Path, fmt.Sprintf("%f,%f", loc.Latitude, loc.Longitude))
			vals := url.Values{}
			vals.Set(langKey, langValue)
			vals.Set(unitsKey, unitsValue)
			locURL.RawQuery = vals.Encode()
			c.locations[i].Url = locURL.String()
		}
	}
	sort.Sort(locArr(c.locations))
}

func (c *conf) Locations() []Location {
	cp := deepcopy.Copy(c.locations)
	return cp.([]Location)
}

func (c *conf) Output() (string, string) {
	return c.snake.GetString("output.brief"), c.snake.GetString("output.summary")
}

func (c *conf) RawDir() string {
	return c.snake.GetString("raw_dir")
}
