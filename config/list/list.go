package list

import (
	"maps"
	"slices"
	"sort"
	"strings"

	"github.com/KasumiMercury/cp-cli/config/option"
	"github.com/spf13/cobra"
)

func NewCmd() *cobra.Command {
	showCmd := &cobra.Command{
		Use:   "list",
		Short: "list current configuration",
		Long:  "list current configuration",
		Run: func(cmd *cobra.Command, args []string) {
			show(cmd)
		},
	}

	return showCmd
}

func show(cmd *cobra.Command) {
	options := option.Options
	keys := slices.Collect(maps.Keys(options))
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
