package option

import (
	"errors"
	"os"
	"strings"

	"github.com/spf13/viper"
)

type Key string

const (
	EditorKey         = Key("editor")
	WorkspaceKey      = Key("workspace")
	CurrentProjectKey = Key("current-project")
)

const NoConfigMessage = "not configured"

var ErrNotDirectory = errors.New("not a directory")

type Item struct {
	String   func() string
	Validate func(string) error
}

var Options = map[Key]Item{
	EditorKey: {
		String: func() string {
			v := viper.GetString("editor")
			if v == "" {
				return NoConfigMessage
			}

			return v
		},
	},
	WorkspaceKey: {
		String: func() string {
			v := viper.GetString("project_dir")
			if v == "" {
				return NoConfigMessage
			}

			return v
		},
		Validate: func(v string) error {
			if f, err := os.Stat(v); os.IsNotExist(err) || !f.IsDir() {
				return ErrNotDirectory
			}

			return nil
		},
	},
	CurrentProjectKey: {
		String: func() string {
			project := viper.GetStringMapString("current_project")
			if project == nil {
				return NoConfigMessage
			}

			strBuilder := strings.Builder{}
			strBuilder.WriteString("\n")

			strBuilder.WriteString("\t")
			strBuilder.WriteString("id: ")
			strBuilder.WriteString(project["id"])
			strBuilder.WriteString("\n")
			strBuilder.WriteString("\t")
			strBuilder.WriteString("site: ")
			strBuilder.WriteString(project["site"])

			return strBuilder.String()
		},
	},
}
