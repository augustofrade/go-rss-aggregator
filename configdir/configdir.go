package configdir

import (
	"os"
	"path"
	"runtime"
)

var appDir string
var feedFile string = "feeds"

type ConfigDir struct {
	DirPath      string
	FeedFilePath string
}

func Init() ConfigDir {
	setAppDirValue()
	createAppDirIfNotExists()
	createFeedFileIfNotExists()
	return ConfigDir{DirPath: appDir, FeedFilePath: feedFile}
}

func FeedFilePath() string {
	return feedFile
}

func createFeedFileIfNotExists() {
	feedFile = path.Join(appDir, feedFile)
	if _, err := os.Stat(feedFile); os.IsNotExist(err) {
		os.WriteFile(feedFile, make([]byte, 0), 0644)
	}
}

func setAppDirValue() {
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
	appDir = path.Join(homeDir, configDir, "kevin")
}

func createAppDirIfNotExists() {
	if _, err := os.Stat(appDir); os.IsNotExist(err) {
		err = os.Mkdir(appDir, 0755)
		if err != nil {
			panic(err)
		}
	}
}
