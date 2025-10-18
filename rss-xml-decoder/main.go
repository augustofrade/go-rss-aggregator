package rssxmldecoder

import (
	"encoding/xml"
)

type RSS struct {
	XMLName xml.Name `xml:"rss"`
	Channel Channel  `xml:"channel"`
}

type Channel struct {
	XMLName xml.Name   `xml:"channel"`
	Title   string     `xml:"title"`
	Link    string     `xml:"link"`
	Items   []FeedItem `xml:"item"`
}

type FeedItem struct {
	XMLName     xml.Name `xml:"item"`
	Title       string   `xml:"title"`
	Link        string   `xml:"link"`
	Description string   `xml:"description"`
}

func Decode(data *[]byte) *[]FeedItem {
	var root RSS
	err := xml.Unmarshal(*data, &root)
	if err != nil {
		panic(err)
	}

	return &root.Channel.Items
}

func truncateString(s *string, length int) string {
	truncated := ""
	count := 0
	for _, char := range *s {
		truncated += string(char)
		count++
		if count >= length {
			break
		}
	}
	return truncated + "..."
}
