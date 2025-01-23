package create

import (
	"errors"
	"fmt"
	"html/template"
	"log/slog"
	"os"
	"os/exec"
	"path/filepath"
)

var cmd *exec.Cmd

var ErrFailedCreateProject = errors.New("failed to create project")

func ProjectDir(pID string) error {
	err := os.Mkdir("./"+pID, os.ModePerm)
	if err != nil {
		return fmt.Errorf("%w(%s): %w", ErrFailedCreateProject, "directory", err)
	}

	cmd = exec.Command("go", "mod", "init", pID)
	cmd.Dir = "./" + pID

	err = cmd.Run()
	if err != nil {
		return fmt.Errorf("%w(%s): %w", ErrFailedCreateProject, "workspace", err)
	}

	return nil
}

func ProjectMain(pID string) (string, error) {
	tmpl, err := template.ParseFiles("create/resource/main.tmpl")
	if err != nil {
		panic("failed to parse template")
	}

	genMainPath := filepath.Join("./"+pID, "main.go")
	genMain, err := os.Create(genMainPath)
	if err != nil {
		return "", fmt.Errorf("%w(%s): %w", ErrFailedCreateProject, "main.go", err)
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			slog.Error("cannot close file", "err", err)
		}
	}(genMain)

	if err := tmpl.Execute(genMain, nil); err != nil {
		return "", fmt.Errorf("%w(%s): %w", ErrFailedCreateProject, "execute template", err)
	}

	return genMainPath, nil
}
