package create

import "github.com/spf13/viper"

func MemoryCurrentProject(pID string) error {
	viper.Set("CURRENT_PROBLEM", pID)

	if err := viper.SafeWriteConfig(); err != nil {
		return err
	}

	return nil
}
