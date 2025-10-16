package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	rssxmldecoder "github.com/augustofrade/go-rss-aggregator/rss-xml-decoder"
)

func main() {
	if len(os.Args[1:]) == 0 {
		log.Fatalln("Expected RSS URL")
		return
	}
	resp, err := http.Get(os.Args[1])
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	contentType := resp.Header.Get("content-type")
	if !strings.Contains(contentType, "xml") {
		panic(fmt.Sprintf("Invalid content type: %s", contentType))
	}

	rssxmldecoder.Decode(&body)
}
