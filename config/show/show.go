package show

import (
	"github.com/KasumiMercury/cp-cli/config/option"
	"github.com/spf13/cobra"
	"golang.org/x/exp/maps"
	"sort"
	"strings"
)

func NewCmd() *cobra.Command {
	showCmd := &cobra.Command{
		Use:   "show",
		Short: "show current configuration",
		Long:  "show current configuration",
		Run: func(cmd *cobra.Command, args []string) {
			show(cmd)
		},
	}

	return showCmd
}

func show(cmd *cobra.Command) {
	options := option.Options
	keys := maps.Keys(options)
	sort.SliceStable(keys, func(i, j int) bool {
		return keys[i] < keys[j]
	})

	sb := strings.Builder{}

	for _, key := range keys {
		sb.WriteString("\n")
		sb.WriteString(string(key))
		sb.WriteString(": ")
		sb.WriteString(options[key].String())
	}

	cmd.Println(sb.String())
}
