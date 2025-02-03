package option

import (
	"github.com/spf13/viper"
	"strings"
)

type OptionKey string

const (
	EditorKey         = OptionKey("editor")
	ProjectDirKey     = OptionKey("project-dir")
	CurrentProjectKey = OptionKey("current-project")
)

type Item struct {
	String func() string
}

var Options = map[OptionKey]Item{
	EditorKey: {
		String: func() string {
			v := viper.GetString("editor")
			if v == "" {
				return "not configured"
			}

			return v
		}},
	ProjectDirKey: {
		String: func() string {
			v := viper.GetString("project_dir")
			if v == "" {
				return "not configured"
			}

			return v
		}},
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
		}},
}
