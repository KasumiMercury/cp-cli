package set

import (
	"errors"
	"fmt"

	"github.com/KasumiMercury/cp-cli/config/option"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	ErrInvalidOption = errors.New("invalid option")
	ErrKeyNotFound   = errors.New("key not found")
	ErrInvalidKey    = errors.New("invalid key")
)

func NewCmd() *cobra.Command {
	setCmd := &cobra.Command{
		Use:   "set <key> <value>",
		Short: "Set configuration",
		Long:  "Set configuration",
		Args:  cobra.ExactArgs(2),
		RunE: func(_ *cobra.Command, args []string) error {
			if err := set(args); err != nil {
				return err
			}

			return nil
		},
	}

	return setCmd
}

func set(args []string) error {
	key := option.Key(args[0])

	if !validOption(key) {
		return fmt.Errorf("%w: %w: %s", ErrInvalidOption, ErrKeyNotFound, args[0])
	}

	validFunc := option.Options[key].Validate
	if validFunc == nil {
		return nil
	}

	if validFunc(args[1]) != nil {
		return fmt.Errorf("%w: %w: %s", ErrInvalidOption, ErrInvalidKey, args[1])
	}

	if err := writeConfig(args[0], args[1]); err != nil {
		return err
	}

	return nil
}

func validOption(key option.Key) bool {
	options := option.Options
	if _, ok := options[key]; !ok {
		return false
	}

	return true
}

func writeConfig(key, value string) error {
	viper.Set(key, value)
	if err := viper.WriteConfig(); err != nil {
		return fmt.Errorf("error writing config: %w", err)
	}

	return nil
}
