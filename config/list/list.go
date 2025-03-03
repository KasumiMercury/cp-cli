package list

import (
	"github.com/KasumiMercury/cp-cli/config/option"
	"github.com/spf13/cobra"
	"sort"
	"strings"
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
	registry := option.NewRegistry()
	options := registry.GetAll()

	// sort options by key
	// to make the output more predictable
	keys := make([]string, 0, len(options))

	for k := range options {
		keys = append(keys, k)
	}

	sort.Strings(keys)

	strBuilder := strings.Builder{}
	strBuilder.Grow(len(options) * 24)

	for i, k := range keys {
		opt := options[k]

		if i > 0 {
			strBuilder.WriteString("\n")
		}

		strBuilder.WriteString(k)
		strBuilder.WriteString(": ")
		strBuilder.WriteString(opt.String())
	}

	cmd.Println(strBuilder.String())
}
