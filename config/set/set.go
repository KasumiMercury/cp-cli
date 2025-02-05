package set

import (
	"fmt"
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
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := set(args); err != nil {
				return err
			}

			return nil
		},
	}

	return setCmd
}

func set(args []string) error {
	if !validOption(args[0]) {
		return fmt.Errorf("%q is an invalid option", args[0])
	}

	if err := writeConfig(args[0], args[1]); err != nil {
		return err
	}

	return nil
}

func validOption(key string) bool {
	options := option.Options
	if _, ok := options[option.OptionKey(key)]; !ok {
		return false
	}

	return true
}

func writeConfig(key, value string) error {
	viper.Set(key, value)
	if err := viper.WriteConfig(); err != nil {
		return err
	}

	return nil
}
