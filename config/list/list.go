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
		Run: func(cmd *cobra.Command, _ []string) {
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

	strBuilder := strings.Builder{}

	for _, key := range keys {
		strBuilder.WriteString("\n")
		strBuilder.WriteString(string(key))
		strBuilder.WriteString(": ")
		strBuilder.WriteString(options[key].String())
	}

	cmd.Println(strBuilder.String())
}
