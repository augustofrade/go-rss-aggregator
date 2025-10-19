package cli

import (
	"fmt"
	"strings"

	rssxmldecoder "github.com/augustofrade/go-rss-aggregator/rss-xml-decoder"
	"github.com/manifoldco/promptui"
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
	prompt := promptui.Select{
		Label: "Choose an article",
		Items: options,
		Templates: &promptui.SelectTemplates{
			Label:    "{{ . }}",
			Active:   "\U00002022 {{ .Label | blue }}",
			Inactive: "  {{ .Label | white }}",
			Selected: "{{ .Label | red | cyan }}",
		},
		Size: GetTerminalHeight(len(options)),
	}

	ClearTerminal()
	fmt.Printf("Listing %d entries\n\n", len(options))
	index, _, err := prompt.Run()
	if err != nil {
		panic(err)
	}

	ClearTerminal()

	selectedArticle := (channel.Articles)[index]
	articleDescription := strings.TrimSpace(selectedArticle.Description)
	articleDescription = strings.ReplaceAll(articleDescription, "[&#8230;]", "[...]")
	articleDescription = strings.ReplaceAll(articleDescription, "&#160;", " ")

	fmt.Printf("%s\n\n[%s]     %s\n\n", channel.Title, selectedArticle.PubDate, selectedArticle.Title)
	fmt.Println(selectedArticle.Link)
	fmt.Printf("\n\n%s\n\n", articleDescription)

	showArticleMenuOptions(&selectedArticle)
}

func ShowFeedsMenu(channels []*rssxmldecoder.Channel) {

	options := make([]CliOption, 0)
	for i, channel := range channels {
		optionLabel := fmt.Sprintf("[%d]  %s", len(channel.Articles), channel.Title)
		options = append(options, CliOption{Label: optionLabel, Value: fmt.Sprint(i)})
	}

	for {
		ClearTerminal()

		prompt := promptui.Select{
			Label: "Choose a feed",
			Items: options,
			Templates: &promptui.SelectTemplates{
				Label:    "{{ . }}",
				Active:   "\U00002022 {{ .Label | blue }}",
				Inactive: "  {{ .Label | white }}",
				Selected: "{{ .Label | cyan }}",
			},
			Size: GetTerminalHeight(len(options)),
		}

		index, _, err := prompt.Run()
		if err != nil {
			if err == promptui.ErrInterrupt {
				return
			}

			panic(err)
		}

		selectedFeed := channels[index]

		ShowArticlesMenu(selectedFeed)
	}
}

func showArticleMenuOptions(article *rssxmldecoder.FeedItem) {
	options := []CliOption{
		{Label: "Back", Value: "back"},
		{Label: "Open in browser", Value: "open"},
	}

	prompt := promptui.Select{
		Label: "Choose an option",
		Items: options,
		Templates: &promptui.SelectTemplates{
			Label:    "{{ . }}",
			Active:   "\U00002022 {{ .Label | blue }}",
			Inactive: "  {{ .Label | white }}",
			Selected: "{{ .Label | cyan }}",
		},
		Size: GetTerminalHeight(len(options)),
	}

	index, _, _ := prompt.Run()

	switch options[index].Value {
	case "open":
		Exec("open", article.Link)
	default:
		fmt.Println("default")
	}
}
