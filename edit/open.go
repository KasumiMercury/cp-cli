package edit

import (
	"errors"
	"github.com/spf13/viper"
	"os"
	"os/exec"
	"path/filepath"
)

var ErrFailedOpenEditor = errors.New("failed to open editor")

func OpenEditor(targetDir string, targetFile string) error {
	targetPath := filepath.Join(targetDir, targetFile)

	editor := viper.GetString("editor")
	if editor == "" {
		editor = "vim"
	}

	cmd := exec.Command(editor, targetPath)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return ErrFailedOpenEditor
	}

	return nil
}
