package config

import (
	"github.com/KasumiMercury/cp-cli/config/option"
	"github.com/spf13/viper"
	"os"
	"runtime"
)

var defaultEditor = ""

func init() {
	if c := viper.GetString(option.EditorKey); c != "" {
		defaultEditor = c

		return
	}

	if v := os.Getenv("VISUAL"); v != "" {
		defaultEditor = v
	} else if e := os.Getenv("EDITOR"); e != "" {
		defaultEditor = e
	} else if runtime.GOOS == "windows" {
		defaultEditor = "notepad"
	} else if runtime.GOOS == "linux" {
		defaultEditor = "vim"
	}
}

func GetDefaultEditor() string {
	return defaultEditor
}
