package create

import (
	"fmt"

	"github.com/spf13/viper"
)

func MemoryCurrentProject(pID string, site string) error {
	current := map[string]string{
		"id":   pID,
		"site": site,
	}

	viper.Set("current_project", current)

	if err := viper.WriteConfig(); err != nil {
		return fmt.Errorf("cannot write config: %w", err)
	}

	return nil
}

func GetCurrentProject() (string, string) {
	current := viper.GetStringMapString("current_project")

	return current["id"], current["site"]
}
