package app

import (
	"github.com/AntonPates/darksky-api-client/services/client"
	"github.com/AntonPates/darksky-api-client/services/config"
	"github.com/AntonPates/darksky-api-client/services/parser"
)

type App interface {
	Bootstrap(configPath string)
	Config() config.Config
	Client() client.Client
	Parser() parser.Parser
	DownloadWF() error
	OutputWFs()
}

var _ App = New()

func New() App {
	return &app{}
}
