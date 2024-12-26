package create

import (
	"fmt"
	"html/template"
	"log/slog"
	"os"
	"os/exec"
)

var cmd *exec.Cmd

func ProjectDir(pID string) error {
	err := os.Mkdir("./"+pID, os.ModePerm)
	if err != nil {
		return fmt.Errorf("cannot create project directory: %w", err)
	}

	cmd = exec.Command("go", "mod", "init", pID)
	cmd.Dir = "./" + pID

	err = cmd.Run()
	if err != nil {
		return fmt.Errorf("cannot create project directory: %w", err)
	}

	return nil
}

func ProjectMain(pID string) error {
	tmpl, err := template.ParseFiles("create/resource/main.tmpl")
	if err != nil {
		return fmt.Errorf("cannot parse template: %w", err)
	}

	genMain, err := os.Create("./" + pID + "/main.go")
	if err != nil {
		return fmt.Errorf("cannot create file: %w", err)
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			slog.Error("cannot close file", "err", err)
		}
	}(genMain)

	if err := tmpl.Execute(genMain, nil); err != nil {
		return fmt.Errorf("cannot execute template: %w", err)
	}

	return nil
}
