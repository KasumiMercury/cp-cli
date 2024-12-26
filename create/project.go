package create

import (
	"fmt"
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
