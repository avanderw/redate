package fileprocessor

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type FileProcessor struct{}

func (fp *FileProcessor) ProcessFiles(inputPath, prefix, pattern string) error {
	err := filepath.Walk(inputPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Match the file with the provided pattern
		matched, err := filepath.Match(pattern, info.Name())
		if err != nil {
			return err
		}

		if matched && !info.IsDir() {
			modTime := info.ModTime().Format("20060102T150405")
			newName := fmt.Sprintf("%s_%s_%s", modTime, prefix, info.Name())
			newPath := filepath.Join(filepath.Dir(path), newName)

			// Prompt for confirmation
			fmt.Printf("Rename: %s -> %s? (y/n): ", path, newPath)
			reader := bufio.NewReader(os.Stdin)
			response, err := reader.ReadString('\n')
			if err != nil {
				return err
			}
			response = strings.TrimSpace(response)

			if strings.ToLower(response) == "y" {
				err := os.Rename(path, newPath)
				if err != nil {
					return err
				}
				fmt.Printf("Renamed: %s -> %s\n", path, newPath)
			} else {
				fmt.Printf("Skipped: %s\n", path)
			}
		}

		return nil
	})

	return err
}
