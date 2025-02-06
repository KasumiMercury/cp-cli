package option

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/viper"
)

type OptionKey string

const (
	EditorKey         = OptionKey("editor")
	ProjectDirKey     = OptionKey("project-dir")
	CurrentProjectKey = OptionKey("current-project")
)

type Item struct {
	String   func() string
	Validate func(string) error
}

var Options = map[OptionKey]Item{
	EditorKey: {
		String: func() string {
			v := viper.GetString("editor")
			if v == "" {
				return "not configured"
			}

			return v
		},
	},
	ProjectDirKey: {
		String: func() string {
			v := viper.GetString("project_dir")
			if v == "" {
				return "not configured"
			}

			return v
		},
		Validate: func(v string) error {
			if f, err := os.Stat(v); os.IsNotExist(err) || !f.IsDir() {
				return fmt.Errorf("%s is not a directory", v)
			}

			return nil
		},
	},
	CurrentProjectKey: {
		String: func() string {
			project := viper.GetStringMapString("current_project")
			if project == nil {
				return "not configured"
			}

			sb := strings.Builder{}
			sb.WriteString("\n")

			sb.WriteString("\t")
			sb.WriteString("id: ")
			sb.WriteString(project["id"])
			sb.WriteString("\n")
			sb.WriteString("\t")
			sb.WriteString("site: ")
			sb.WriteString(project["site"])

			return sb.String()
		},
	},
}
