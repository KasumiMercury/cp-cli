package path

import (
	"fmt"
	"path/filepath"
)

func ProjectDirPath(base string, problemID string) (string, error) {
	baseAbs, err := filepath.Abs(base)
	if err != nil {
		return "", fmt.Errorf("failed to get base absolute path: %w", err)
	}
	dirPath := filepath.Join(baseAbs, problemID)

	return dirPath, nil
}
