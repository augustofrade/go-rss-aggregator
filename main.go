package main

import (
	"github.com/augustofrade/go-rss-aggregator/cli"
	rssxmldecoder "github.com/augustofrade/go-rss-aggregator/rss-xml-decoder"
)

func main() {
	xmlBytes := cli.Init()

	rssxmldecoder.Decode(xmlBytes)
}
