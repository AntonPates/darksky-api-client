package main

import (
	"flag"

	"os"

	"github.com/AntonPates/darksky-api-client/services/app"
	"github.com/sirupsen/logrus"
)

var configFile = flag.String("config", "config.yaml", "config path")

func main() {
	flag.Parse()
	app := app.New()
	app.Bootstrap(*configFile)
	err := app.DownloadWF()
	if err != nil {
		logrus.Error(err)
		os.Exit(1)
	}
	app.OutputWFs()

	return
}
