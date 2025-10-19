package cli

import (
	"flag"
	"os"
	"os/exec"

	"golang.org/x/term"
)

var (
	file = flag.String("file", "", "Local RSS content file to be parsed")
	url  = flag.String("url", "", "External RSS content to be fetched and parsed")
)

type ContentSource struct {
	Origin string
	Mode   string
}

func Init() ContentSource {
	flag.Parse()

	if *file != "" {
		return ContentSource{Mode: "file", Origin: *file}
	} else if *url != "" {
		return ContentSource{Mode: "url", Origin: *url}
	}

	return ContentSource{Mode: "default"}
}

func ClearTerminal() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func GetTerminalHeight(lines int) int {
	headerSize := 4
	menuHeight := lines
	if term.IsTerminal(0) {
		_, height, err := term.GetSize(0)
		if err != nil {
			panic(err)
		}
		menuHeight = height - headerSize
	}

	return menuHeight
}
