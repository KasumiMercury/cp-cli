package editor

import (
	"errors"
	"io"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/KasumiMercury/cp-cli/config"
)

var (
	ErrInvalidCommand   = errors.New("invalid command")
	ErrFailedOpenEditor = errors.New("failed to open editor")
)

func Open(editorCmd string, targetDir string, targetFile string,
	stdin io.Reader, stdout io.Writer, stderr io.Writer,
) error {
	if editorCmd == "" {
		editorCmd = config.GetDefaultEditor()
	}

	if len(editorCmd) == 0 {
		return ErrInvalidCommand
	}

	args := strings.Split(editorCmd, " ")

	targetPath := filepath.Join(targetDir, targetFile)
	args = append(args, targetPath)

	// #nosec G204
	cmd := exec.Command(args[0], args[1:]...)
	cmd.Stdin = stdin
	cmd.Stdout = stdout
	cmd.Stderr = stderr

	if err := cmd.Run(); err != nil {
		return ErrFailedOpenEditor
	}

	return nil
}
