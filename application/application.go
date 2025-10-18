package application

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/augustofrade/go-rss-aggregator/cli"
	rssxmldecoder "github.com/augustofrade/go-rss-aggregator/rss-xml-decoder"
)

func HandleLocalFile(path *string) error {
	bodyBytes, err := fetchLocalFile(path)
	if err != nil {
		return err
	}

	rssChannel := rssxmldecoder.Decode(bodyBytes)
	cli.ShowArticlesMenu(rssChannel)

	return nil
}

func HandleExternalUrl(path *string) error {
	bodyBytes, err := fetchExternalFile(path)
	if err != nil {
		return err
	}

	rssChannel := rssxmldecoder.Decode(bodyBytes)
	cli.ShowArticlesMenu(rssChannel)

	return nil
}

func fetchLocalFile(file *string) (*[]byte, error) {
	fmt.Printf("Reading feed from file: %s...\n\n", *file)

	body, err := os.ReadFile(*file)

	return &body, err
}

func fetchExternalFile(url *string) (*[]byte, error) {
	fmt.Printf("Reading feed from URL: %s...\n\n", *url)

	resp, err := http.Get(*url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	contentType := resp.Header.Get("content-type")
	if !strings.Contains(contentType, "xml") {
		return nil, fmt.Errorf("invalid content type: %s", contentType)
	}

	return &body, nil
}
