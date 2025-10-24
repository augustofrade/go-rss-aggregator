package rssxmldecoder

import (
	"encoding/xml"
	"errors"
)

type Rss struct {
	XMLName xml.Name   `xml:"rss"`
	Channel RssChannel `xml:"channel"`
}

type RssChannel struct {
	XMLName  xml.Name  `xml:"channel"`
	Title    string    `xml:"title"`
	Link     string    `xml:"link"`
	Articles []RssItem `xml:"item"`
}

type RssItem struct {
	XMLName     xml.Name `xml:"item"`
	Title       string   `xml:"title"`
	Link        string   `xml:"link"`
	PubDate     string   `xml:"pubDate"`
	Description string   `xml:"description"`
}

type Feed struct {
	Title    string
	Link     string
	Articles []FeedArticle
}

type FeedArticle struct {
	Title       string
	Link        string
	PubDate     string
	Description string
}

func Decode(data *[]byte, feedType string) (*Feed, error) {
	switch feedType {
	case "rss":
		return decodeAndMapRSS(data)
	default:
		return nil, errors.New("invalid feed type, expected rss or atom")
	}
}

func decodeAndMapRSS(data *[]byte) (*Feed, error) {
	var root Rss
	err := xml.Unmarshal(*data, &root)
	if err != nil {
		return nil, err
	}

	feed := Feed{
		Title:    root.Channel.Title,
		Link:     root.Channel.Link,
		Articles: make([]FeedArticle, 0),
	}

	for _, item := range root.Channel.Articles {
		feed.Articles = append(feed.Articles, FeedArticle{
			Title:       item.Title,
			Link:        item.Link,
			PubDate:     item.PubDate,
			Description: item.Description,
		})
	}

	return &feed, nil
}
