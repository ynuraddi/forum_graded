package utils

import (
	"fmt"
	"os"
	"path/filepath"
)

func ParseDirectoryFiles(directory, template string) (string, error) {
	files, err := filepath.Glob(filepath.Join(directory, template))
	if err != nil {
		return "", fmt.Errorf("%v: %s", files, err)
	}

	var content string

	for _, file := range files {
		fileContent, err := os.ReadFile(file)
		if err != nil {
			return "", err
		}

		content += string(fileContent)
	}

	return content, nil
}
