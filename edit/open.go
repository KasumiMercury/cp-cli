package edit

import (
	"errors"
	"os"
	"os/exec"
)

var ErrFailedOpenEditor = errors.New("failed to open editor")

func OpenEditor(targetPath string) error {
	editor := os.Getenv("EDITOR")
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
