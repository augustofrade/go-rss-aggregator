package cli

import (
	"fmt"
	"strings"

	rssxmldecoder "github.com/augustofrade/go-rss-aggregator/rss-xml-decoder"
	"github.com/manifoldco/promptui"
	"golang.org/x/term"
)

type CliOption struct {
	Label string
	Value string
}

func ShowArticlesMenu(channel *rssxmldecoder.Channel) {
	options := make([]CliOption, 0)

	for i, item := range channel.Articles {
		options = append(options, CliOption{Label: item.Title, Value: fmt.Sprint(i)})
	}

	menuHeight := len(options)
	if term.IsTerminal(0) {
		_, height, err := term.GetSize(0)
		if err != nil {
			panic(err)
		}
		menuHeight = height
	}

	prompt := promptui.Select{
		Label: "Choose an article",
		Items: options,
		Templates: &promptui.SelectTemplates{
			Label:    "{{ . }}",
			Active:   "\U00002022 {{ .Label | blue }}",
			Inactive: "  {{ .Label | white }}",
			Selected: "{{ .Label | red | cyan }}",
		},
		Size: menuHeight,
	}

	// cli.Clear()
	fmt.Printf("Listing %d entries\n\n", len(options))
	index, _, err := prompt.Run()
	if err != nil {
		panic(err)
	}

	Clear()

	selectedArticle := (channel.Articles)[index]
	articleDescription := strings.TrimSpace(selectedArticle.Description)
	articleDescription = strings.ReplaceAll(articleDescription, "[&#8230;]", "[...]")
	articleDescription = strings.ReplaceAll(articleDescription, "&#160;", " ")

	fmt.Printf("%s\n\n[%s]     %s\n\n", channel.Title, selectedArticle.PubDate, selectedArticle.Title)
	fmt.Println(selectedArticle.Link)
	fmt.Printf("\n\n%s\n\n", articleDescription)
}
