package create

import (
	"errors"
	"fmt"
	"html/template"
	"os"
	"os/exec"
	"path/filepath"
)

var cmd *exec.Cmd

var ErrFailedCreateProject = errors.New("failed to create project")

func ProjectDir(path string, problemID string) error {
	err := os.Mkdir(path, os.ModePerm)
	if err != nil {
		return fmt.Errorf("%w(%s): %w", ErrFailedCreateProject, "directory", err)
	}

	cmd = exec.Command("go", "mod", "init", problemID)
	cmd.Dir = path

	err = cmd.Run()
	if err != nil {
		return fmt.Errorf("%w(%s): %w", ErrFailedCreateProject, "workspace", err)
	}

	if err := projectMain(path); err != nil {
		return err
	}

	return nil
}

func projectMain(path string) error {
	tmpl, err := template.ParseFiles("create/resource/main.tmpl")
	if err != nil {
		panic("failed to parse template")
	}

	genMainPath := filepath.Join(path, "main.go")
	genMain, err := os.Create(genMainPath)
	if err != nil {
		return fmt.Errorf("%w(%s): %w", ErrFailedCreateProject, "main.go", err)
	}
	defer genMain.Close()

	if err := tmpl.Execute(genMain, nil); err != nil {
		return fmt.Errorf("%w(%s): %w", ErrFailedCreateProject, "execute template", err)
	}

	if err := genMain.Sync(); err != nil {
		return fmt.Errorf("%w: failed to sync %s: %w", ErrFailedCreateProject, "main.go", err)
	}

	return nil
}
