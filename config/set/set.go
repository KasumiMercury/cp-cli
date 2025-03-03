package set

import (
	"errors"
	"fmt"

	"github.com/KasumiMercury/cp-cli/config/option"
	"github.com/spf13/cobra"
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

var (
	ErrInvalidOption = errors.New("invalid option")
	ErrFailedConfig  = errors.New("failed to configure")
)

func set(args []string) error {
	registry := option.NewRegistry()

	opt, err := registry.Get(args[0])
	if err != nil {
		return fmt.Errorf("%w: %w", ErrInvalidOption, err)
	}

	if err := opt.Validate(args[1]); err != nil {
		return fmt.Errorf("%w: %w", ErrInvalidOption, err)
	}

	if err := opt.SetValue(args[1]); err != nil {
		return fmt.Errorf("%w: %w", ErrFailedConfig, err)
	}

	return nil
}
