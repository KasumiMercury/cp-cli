package create

import (
	"fmt"
	"html/template"
	"log/slog"
	"os"
	"os/exec"
)

var cmd *exec.Cmd

func ProjectDir(id string) error {
	err := os.Mkdir("./"+id, os.ModePerm)
	if err != nil {
		return fmt.Errorf("cannot create project directory: %w", err)
	}

	cmd = exec.Command("go", "mod", "init", id)
	cmd.Dir = "./" + id

	err = cmd.Run()
	if err != nil {
		return fmt.Errorf("cannot create project directory: %w", err)
	}

	return nil
}

func ProjectMain(id string) error {
	t, err := template.ParseFiles("create/resource/main.tmpl")
	if err != nil {
		return fmt.Errorf("cannot parse template: %w", err)
	}

	f, err := os.Create("./" + id + "/main.go")
	if err != nil {
		return fmt.Errorf("cannot create file: %w", err)
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			slog.Error("cannot close file", "err", err)
		}
	}(f)

	if err := t.Execute(f, nil); err != nil {
		return fmt.Errorf("cannot execute template: %w", err)
	}
	return nil
}
