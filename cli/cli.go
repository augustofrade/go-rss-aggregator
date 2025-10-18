package cli

import (
	"flag"
	"os"
	"os/exec"
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
