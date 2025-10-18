package main

import (
	"fmt"
	"os"
	"path"
	"runtime"
	"strings"

	"github.com/augustofrade/go-rss-aggregator/cli"
	rssxmldecoder "github.com/augustofrade/go-rss-aggregator/rss-xml-decoder"
	"github.com/manifoldco/promptui"
	"golang.org/x/term"
)

type CliOption struct {
	Label string
	Value string
}

func main() {
	handleAppDir()
	xmlBytes := cli.Init()
	articles := rssxmldecoder.Decode(xmlBytes)

	options := make([]CliOption, 0)

	for i, item := range *articles {
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
			Active:   "\U00002022 {{ .Label | cyan }}",
			Inactive: "  {{ .Label }}",
			Selected: "{{ .Label | red | cyan }}",
		},
		Size: menuHeight,
	}

	cli.Clear()
	fmt.Printf("Listing %d entries\n\n", len(options))
	index, _, err := prompt.Run()
	if err != nil {
		panic(err)
	}

	cli.Clear()

	selectedArticle := (*articles)[index]
	articleDescription := strings.TrimSpace(selectedArticle.Description)
	articleDescription = strings.ReplaceAll(articleDescription, "[&#8230;]", "[...]")
	articleDescription = strings.ReplaceAll(articleDescription, "&#160;", " ")
	fmt.Printf("[%s]     %s\n\n", selectedArticle.PubDate, selectedArticle.Title)
	fmt.Println(selectedArticle.Link)
	fmt.Printf("\n\n%s\n\n", articleDescription)
}

func handleAppDir() string {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}
	var configDir string
	switch runtime.GOOS {
	case "windows":
		configDir = "Documents"
	default:
		configDir = ".config"
	}
	appDir := path.Join(homeDir, configDir, "kevin")

	if _, err = os.Stat(appDir); os.IsNotExist(err) {
		err = os.Mkdir(appDir, 0755)
		if err != nil {
			panic(err)
		}
	}

	return appDir
}
