package cli

import (
	"flag"
	"os"
	"os/exec"
	"strings"

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

func Exec(command string, args ...string) {
	cmd := exec.Command(command, args...)
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func ClearTerminal() {
	Exec("clear")
}

func GetTerminalWidth() int {
	if term.IsTerminal(0) {
		width, _, err := term.GetSize(0)
		if err != nil {
			panic(err)
		}
		return width
	}
	return 60
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

func Separator() string {
	return strings.Repeat("-", GetTerminalWidth())
}
