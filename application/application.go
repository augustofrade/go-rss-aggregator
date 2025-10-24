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
	fmt.Printf("Reading feed from file: %s...\n\n", *path)
	bodyBytes, err := fetchLocalFile(path)
	if err != nil {
		return err
	}

	rssChannel, err := rssxmldecoder.Decode(bodyBytes)
	if err != nil {
		return err
	}

	cli.ShowArticlesMenu(rssChannel)

	return nil
}

func HandleExternalUrl(url *string) error {
	fmt.Printf("Reading feed from URL: %s...\n\n", *url)
	bodyBytes, err := fetchExternalFile(url)
	if err != nil {
		return err
	}

	rssChannel, err := rssxmldecoder.Decode(bodyBytes)
	if err != nil {
		return err
	}

	cli.ShowArticlesMenu(rssChannel)

	return nil
}

func fetchLocalFile(file *string) (*[]byte, error) {

	body, err := os.ReadFile(*file)

	return &body, err
}

func fetchExternalFile(url *string) (*[]byte, error) {
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
