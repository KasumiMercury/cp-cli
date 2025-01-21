package create

import "github.com/spf13/viper"

func MemoryCurrentProject(pID string, site string) error {
	current := map[string]string{
		"id":   pID,
		"site": site,
	}

	viper.Set("current_project", current)

	if err := viper.WriteConfig(); err != nil {
		return err
	}

	return nil
}
