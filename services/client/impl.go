package client

import (
	"bytes"
	"io"
	"net/http"
	"net/url"

	log "github.com/sirupsen/logrus"
)

type cli struct{}

func (c *cli) Get(href string) (io.Reader, error) {
	if u, err := url.Parse(href); err != nil {
		log.Error(err)
		return nil, err
	} else {
		//TODO: replace with fasthttp
		if resp, err := http.Get(u.String()); err != nil {
			log.Error(err)
			return nil, err
		} else {
			defer resp.Body.Close()
			b := make([]byte, 0)
			buf := bytes.NewBuffer(b)
			if _, err := io.Copy(buf, resp.Body); err != nil {
				log.Error(err)
				return nil, err
			}
			return buf, nil
		}
	}
}
