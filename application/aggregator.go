package application

import (
	"fmt"
	"log"
	"os"
	"strings"
	"sync"

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

	var wg sync.WaitGroup

	for _, url := range urls {
		wg.Go(func() {
			currentFeed, err := agg.handleSingleFeed(&url)
			if err != nil {
				log.Printf("Failed fetching \"%s\"\n%s\n", url, err)
			}
			agg.feeds = append(agg.feeds, currentFeed)
		})
	}

	wg.Wait()
	fmt.Printf("Found %d articles\n\n", len(agg.feeds))
}

func (agg *Aggregator) handleSingleFeed(url *string) (*rssxmldecoder.Channel, error) {
	fmt.Println("Handling ", *url)
	currentFeed, err := fetchExternalFile(url)
	if err != nil {
		return nil, err
	}

	return rssxmldecoder.Decode(currentFeed)
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
