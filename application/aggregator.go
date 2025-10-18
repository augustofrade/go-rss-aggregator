package application

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/augustofrade/go-rss-aggregator/configdir"
	rssxmldecoder "github.com/augustofrade/go-rss-aggregator/rss-xml-decoder"
)

type Aggregator struct {
	feeds []*rssxmldecoder.Channel
}

func (agg *Aggregator) Handle() error {
	urls, err := agg.getFeedsURLs()
	if err != nil {
		return err
	}

	agg.fetchFeeds(urls)

	return nil
}

func (agg *Aggregator) fetchFeeds(urls []string) {
	parsedData := make([]*rssxmldecoder.Channel, 0)

	for _, url := range urls {
		currentFeed, err := agg.handleSingleFeed(&url)
		if err != nil {
			log.Printf("Failed fetching \"%s\"\n%s\n", url, err)
		}
		parsedData = append(parsedData, currentFeed)
	}
}

func (agg *Aggregator) handleSingleFeed(url *string) (*rssxmldecoder.Channel, error) {
	fmt.Println("Handling ", *url)
	currentFeed, err := fetchExternalFile(url)
	if err != nil {
		return nil, err
	}

	return rssxmldecoder.Decode(currentFeed), nil
}

func (agg *Aggregator) getFeedsURLs() ([]string, error) {
	data, err := os.ReadFile(configdir.FeedFilePath())
	if err != nil {
		return nil, err
	}

	urls := make([]string, 0)
	for url := range strings.SplitSeq(string(data), "\n") {
		url = strings.TrimSpace(url)
		if url != "" {
			urls = append(urls, url)
		}
	}

	return urls, nil
}
