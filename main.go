package main

import (
	"github.com/augustofrade/go-rss-aggregator/cli"
	"github.com/augustofrade/go-rss-aggregator/configdir"
	rssxmldecoder "github.com/augustofrade/go-rss-aggregator/rss-xml-decoder"
)

func main() {
	configdir.Init()
	xmlBytes := cli.Init()
	channel := rssxmldecoder.Decode(xmlBytes)

	cli.ShowArticlesMenu(channel)
}
