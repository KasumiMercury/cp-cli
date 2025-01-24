package create

import (
	"os"
	"path"
	"reflect"
	"testing"

	"github.com/spf13/viper"
)

func TestMemoryCurrentProject(t *testing.T) {
	tempDir := t.TempDir()
	viper.AddConfigPath(tempDir)
	viper.SetConfigType("yaml")
	viper.SetConfigName(".cp-cli")

	_, _ = os.Create(path.Join(tempDir, ".cp-cli.yaml"))

	tempProjectID := "example_id"
	tempProjectSite := "example.com"

	if err := MemoryCurrentProject(tempProjectID, tempProjectSite); err != nil {
		t.Errorf("failed to open project %v", err)
	}

	wantCurrent := map[string]string{
		"id":   tempProjectID,
		"site": tempProjectSite,
	}

	actualCurrentProject := viper.Get("current_project")
	if !reflect.DeepEqual(actualCurrentProject, wantCurrent) {
		t.Errorf("current_project = %v, want %v", actualCurrentProject, wantCurrent)
	}
}
