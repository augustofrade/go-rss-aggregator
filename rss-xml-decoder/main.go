package rssxmldecoder

import (
	"bytes"
	"encoding/xml"
)

type Rss struct {
	XMLName xml.Name   `xml:"rss"`
	Channel RssChannel `xml:"channel"`
}

type Atom struct {
	XMLName xml.Name    `xml:"feed"`
	Title   string      `xml:"title"`
	Link    AtomLink    `xml:"link"`
	Entries []AtomEntry `xml:"entry"`
}

type AtomLink struct {
	Href string `xml:"href,attr"`
}

type AtomEntry struct {
	XMLName   xml.Name `xml:"entry"`
	Title     string   `xml:"title"`
	Published string   `xml:"published"`
	Summary   string   `xml:"summary"`
	Link      AtomLink `xml:"link"`
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

func Decode(data *[]byte) (*Feed, error) {
	if isAtom(data) {
		return decodeAndMapAtom(data)
	}
	return decodeAndMapRSS(data)
}

func isAtom(data *[]byte) bool {
	decoder := xml.NewDecoder(bytes.NewReader(*data))
	for {
		t, _ := decoder.Token()
		if t == nil {
			return false
		}
		if se, ok := t.(xml.StartElement); ok {
			switch se.Name.Local {
			case "feed":
				return true
			case "rss":
				return false
			}
		}
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

func decodeAndMapAtom(data *[]byte) (*Feed, error) {
	var root Atom
	err := xml.Unmarshal(*data, &root)
	if err != nil {
		return nil, err
	}

	feed := Feed{
		Title:    root.Title,
		Link:     root.Link.Href,
		Articles: make([]FeedArticle, 0),
	}

	for _, item := range root.Entries {
		feed.Articles = append(feed.Articles, FeedArticle{
			Title:       item.Title,
			Link:        item.Link.Href,
			PubDate:     item.Published,
			Description: item.Summary,
		})
	}

	return &feed, nil
}
