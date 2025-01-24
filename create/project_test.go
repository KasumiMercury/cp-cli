package create

import (
	"os"
	"path/filepath"
	"testing"
)

func TestProjectDir(t *testing.T) {
	t.Parallel()

	tempDir := t.TempDir()
	projectDir := filepath.Join(tempDir, "temp")
	problemID := "example_id"

	if err := ProjectDir(projectDir, problemID); err != nil {
		t.Errorf("ProjectDir() error = %v", err)
	}

	if _, err := os.Stat(projectDir); os.IsNotExist(err) {
		t.Errorf("ProjectDir() error = %v", err)
	}
}

func Test_projectMain(t *testing.T) {
	t.Parallel()

	tempDir := t.TempDir()

	if err := projectMain(tempDir); err != nil {
		t.Errorf("projectMain() error = %v", err)
	}

	if _, err := os.Stat(tempDir); os.IsNotExist(err) {
		t.Errorf("projectMain() error = %v", err)
	}
}
