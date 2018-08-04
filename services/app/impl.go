package app

import (
	"os"
	"path/filepath"

	"io"

	"sync"

	"github.com/AntonPates/darksky-api-client/services/client"
	"github.com/AntonPates/darksky-api-client/services/config"
	"github.com/AntonPates/darksky-api-client/services/converter"
	"github.com/AntonPates/darksky-api-client/services/parser"
	log "github.com/sirupsen/logrus"
)

type app struct {
	config    config.Config
	client    client.Client
	parser    parser.Parser
	converter converter.Converter
}

func (a *app) Bootstrap(configPath string) {
	a.config = config.New()
	if err := a.config.Load(configPath); err != nil {
		log.Fatal(err)

	}
	a.client = client.New()
	a.parser = parser.New()
	a.converter = converter.New()

}

func (a *app) Config() config.Config {
	return a.config
}

func (a *app) Client() client.Client {
	return a.client
}

func (a *app) Parser() parser.Parser {
	return a.parser
}

func (a *app) BriefWF() {

}

func (a *app) DownloadWF() error {
	log.Info("Start download wf for all locations")
	locs := a.config.Locations()
	var wg sync.WaitGroup
	for i := range locs {
		wg.Add(1)
		go func(href, anchor, name string) {
			defer wg.Done()
			log.Infof("Get WF for %s - %s", anchor, href)
			if r, err := a.client.Get(href); err != nil {
				return
			} else {
				dir := a.config.RawDir()
				if f, err := os.Create(filepath.Join(dir, anchor+".json")); err != nil {
					log.Error(err)
					return
				} else {
					io.Copy(f, r)
					f.Close()
				}
			}
		}(locs[i].Url, locs[i].Anchor, locs[i].Name)
	}
	wg.Wait()
	return nil
}

func (a *app) rawWFs() []*parser.Weather {
	log.Info("Get wf for all locations")
	locs := a.config.Locations()
	wfs := make([]*parser.Weather, 0, len(locs))
	for i := range locs {
		dir := a.config.RawDir()
		if f, err := os.Open(filepath.Join(dir, locs[i].Anchor+".json")); err != nil {
			log.Error(err)
			continue
		} else {
			wf, err := a.parser.Exec(f)
			if err != nil {
				continue
			}
			wf.Location = locs[i]
			wfs = append(wfs, wf)
		}
	}
	return wfs
}

func (a *app) OutputWFs() {
	wfs := a.rawWFs()
	brief, summary := a.config.Output()
	r, err := a.converter.CurrentWeather(wfs)
	if err != nil {
		log.Error(err)
		return
	}

	f, err := os.Create(brief)
	if err != nil {
		log.Error(err)
		return
	}
	defer f.Close()
	io.Copy(f, r)

	r, err = a.converter.Summary(wfs)
	if err != nil {
		log.Error(err)
		return
	}
	s, err := os.Create(summary)
	if err != nil {
		log.Error(err)
		return
	}
	defer s.Close()
	io.Copy(s, r)

}
