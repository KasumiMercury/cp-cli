package edit

import (
	"fmt"
	"os"
	"os/exec"
)

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
		return fmt.Errorf("cannot open %s with %s: %w", targetPath, editor, err)
	}

	return nil
}
