package cli

import (
	"fmt"
	"os"
	"strings"

	rssxmldecoder "github.com/augustofrade/go-rss-aggregator/rss-xml-decoder"
	"github.com/manifoldco/promptui"
)

type CliOption struct {
	Label string
	Value string
}

var termWidth int = GetTerminalWidth()

func ShowArticlesMenu(channel *rssxmldecoder.Feed) {
	options := make([]CliOption, 0)

	for i, item := range channel.Articles {
		truncatedTitle := truncateText(item.Title, termWidth)
		options = append(options, CliOption{Label: truncatedTitle, Value: fmt.Sprint(i)})
	}

	for {
		ClearTerminal()
		fmt.Println(channel.Title)
		fmt.Println(Separator())
		fmt.Printf("\nListing %d entries\n\n", len(options))
		fmt.Println()

		selectedIndex := displayGenericSelect("Choose an article:", options)
		if selectedIndex == -1 {
			return
		}

		ClearTerminal()

		selectedArticle := (channel.Articles)[selectedIndex]
		articleDescription := strings.TrimSpace(selectedArticle.Description)
		articleDescription = strings.ReplaceAll(articleDescription, "[&#8230;]", "[...]")
		articleDescription = strings.ReplaceAll(articleDescription, "&#160;", " ")

		fmt.Printf("%s\n\n[%s]     %s\n\n", channel.Title, selectedArticle.PubDate, selectedArticle.Title)
		fmt.Println(selectedArticle.Link)
		fmt.Printf("\n\n%s\n\n", articleDescription)

		showArticleMenuOptions(&selectedArticle)
	}
}

func ShowFeedsMenu(channels []*rssxmldecoder.Feed) {

	options := make([]CliOption, 0)
	for i, channel := range channels {
		truncatedTitle := truncateText(channel.Title, termWidth)
		optionLabel := fmt.Sprintf("[%d]  %s", len(channel.Articles), truncatedTitle)
		options = append(options, CliOption{Label: optionLabel, Value: fmt.Sprint(i)})
	}
	for {
		ClearTerminal()
		fmt.Printf("%d Feeds\n\n", len(options))
		selectedIndex := displayGenericSelect("Choose a feed:", options)
		if selectedIndex == -1 {
			return
		}
		selectedFeed := channels[selectedIndex]

		ShowArticlesMenu(selectedFeed)
	}
}

func showArticleMenuOptions(article *rssxmldecoder.FeedArticle) {
	options := []CliOption{
		{Label: "Back", Value: "back"},
		{Label: "Open in browser", Value: "open"},
	}

	index := displayGenericSelect("Choose an option", options)
	if index == -1 {
		os.Exit(0)
		return
	}

	if options[index].Value == "open" {
		Exec("open", article.Link)
	}
}

func displayGenericSelect(promptLabel string, options []CliOption) int {

	prompt := promptui.Select{
		Label: promptLabel,
		Items: options,
		Templates: &promptui.SelectTemplates{
			Label:    "{{ . }}",
			Active:   "\U00002022 {{ .Label | blue }}",
			Inactive: "  {{ .Label | white }}",
			Selected: "{{ .Label | cyan }}",
		},
		Size: 20,
	}

	index, _, err := prompt.Run()
	if err != nil {
		if err == promptui.ErrInterrupt {
			ClearTerminal()
			return -1
		}
		panic(err)
	}

	return index
}

func truncateText(label string, max int) string {
	if len(label) > max {
		return label[:max-10] + "..."
	}
	return label
}
