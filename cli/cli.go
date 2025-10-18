package cli

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"strings"
)

var (
	file = flag.String("file", "", "Local RSS content file to be parsed")
	url  = flag.String("url", "", "External RSS content to be fetched and parsed")
)

func Init() *[]byte {
	flag.Parse()

	var err error
	var body []byte

	if *file != "" {
		err = handleLocalFileParse(&body)
	} else if *url != "" {
		err = handleUrlParse(&body)
	} else {
		panic("--file or --url required")
	}

	if err != nil {
		panic(err)
	}

	return &body
}

func handleLocalFileParse(body *[]byte) error {
	fmt.Println("Using file: ", *file)

	var err error
	*body, err = os.ReadFile(*file)

	return err
}

func handleUrlParse(body *[]byte) error {
	fmt.Println("Fetching from URL: ", *url)

	resp, err := http.Get(*url)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	*body, err = io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	contentType := resp.Header.Get("content-type")
	if !strings.Contains(contentType, "xml") {
		return fmt.Errorf("invalid content type: %s", contentType)
	}

	return nil
}

func Clear() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}
