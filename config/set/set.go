package set

import (
	"github.com/KasumiMercury/cp-cli/config/option"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func NewCmd() *cobra.Command {
	setCmd := &cobra.Command{
		Use:   "set <key> <value>",
		Short: "Set configuration",
		Long:  "Set configuration",
		Args:  cobra.ExactArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			set(cmd, args)
		},
	}

	return setCmd
}

func set(cmd *cobra.Command, args []string) {
	if !validOption(args[0]) {
		return
	}

	viper.Set(args[0], args[1])
	err := viper.WriteConfig()
	if err != nil {
		return
	}
}

func validOption(key string) bool {
	options := option.Options
	if _, ok := options[option.OptionKey(key)]; !ok {
		return false
	}

	return true
}
