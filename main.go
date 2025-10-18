package main

import (
	"fmt"

	"github.com/augustofrade/go-rss-aggregator/application"
	"github.com/augustofrade/go-rss-aggregator/cli"
	"github.com/augustofrade/go-rss-aggregator/configdir"
)

func main() {
	configdir.Init()
	applicationMode := cli.Init()
	origin := &applicationMode.Origin

	var err error

	switch applicationMode.Mode {
	case "file":
		err = application.HandleLocalFile(origin)
	case "url":
		err = application.HandleExternalUrl(origin)
	default:
		fmt.Println("Default mode")
	}

	if err != nil {
		fmt.Println(err)
	}
}
