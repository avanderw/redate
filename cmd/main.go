package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

type FileProcessor struct {
	DryRun    bool
	Force     bool
	Recursive bool
}

func (fp *FileProcessor) ProcessFiles(inputPath, pattern string) error {
	if fp.Recursive {
		return fp.processFilesRecursive(inputPath, pattern)
	}
	return fp.processFilesNonRecursive(inputPath, pattern)
}

func (fp *FileProcessor) processFilesRecursive(inputPath, pattern string) error {
	return filepath.Walk(inputPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() {
			return fp.processFile(path, info, pattern)
		}

		return nil
	})
}

func (fp *FileProcessor) processFilesNonRecursive(inputPath, pattern string) error {
	files, err := os.ReadDir(inputPath)
	if err != nil {
		return err
	}

	for _, file := range files {
		if !file.IsDir() {
			info, err := file.Info()
			if err != nil {
				return err
			}
			err = fp.processFile(filepath.Join(inputPath, file.Name()), info, pattern)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func (fp *FileProcessor) processFile(path string, info os.FileInfo, pattern string) error {
	matched, err := filepath.Match(pattern, info.Name())
	if err != nil {
		return err
	}

	if matched {
		modTime := info.ModTime().Format("20060102T150405")
		newName := fp.updateFileName(info.Name(), modTime)

		// Skip if the filename wouldn't change
		if newName == info.Name() {
			fmt.Printf("Skipped: %s (no change needed)\n", path)
			return nil
		}

		newPath := filepath.Join(filepath.Dir(path), newName)

		if fp.DryRun {
			fmt.Printf("Would rename: %s -> %s\n", path, newPath)
			return nil
		}

		shouldRename := fp.Force

		// Prompt for confirmation if not in force mode
		if !fp.Force {
			fmt.Printf("Rename \"%s\" to \"%s\"? [y/N]: ", filepath.Base(path), filepath.Base(newPath))
			reader := bufio.NewReader(os.Stdin)
			response, err := reader.ReadString('\n')
			if err != nil {
				return err
			}
			response = strings.TrimSpace(response)

			shouldRename = strings.ToLower(response) == "y"
		}

		if shouldRename {
			if !fp.DryRun {
				err := os.Rename(path, newPath)
				if err != nil {
					return err
				}
			}
			fmt.Printf("✓ Renamed to \"%s\"\n", filepath.Base(newPath))
		} else {
			fmt.Printf("⨯ Skipped \"%s\"\n", filepath.Base(path))
		}
	}

	return nil
}

func (fp *FileProcessor) updateFileName(fileName, newDate string) string {
	re := regexp.MustCompile(`^\d{8}T\d{6}_`)
	if re.MatchString(fileName) {
		existingDate := re.FindString(fileName)
		if existingDate == newDate+"_" {
			return fileName // No change needed
		}
		return newDate + "_" + fileName[len(existingDate):]
	}
	return newDate + "_" + fileName
}

func main() {
	// Define all command line flags
	recursive := flag.Bool("r", false, "Recursively process files in subdirectories")
	recursiveLong := flag.Bool("recursive", false, "Recursively process files in subdirectories")

	dryRun := flag.Bool("d", false, "Show what would be renamed without actually renaming")
	dryRunLong := flag.Bool("dry-run", false, "Show what would be renamed without actually renaming")

	force := flag.Bool("f", false, "Rename without prompting for confirmation")
	forceLong := flag.Bool("force", false, "Rename without prompting for confirmation")

	flag.Parse()

	if flag.NArg() < 1 {
		log.Fatalf("Error: File pattern is required")
	}

	pattern := flag.Arg(0)
	inputPath := "."

	// Combine short and long flag options
	isRecursive := *recursive || *recursiveLong
	isDryRun := *dryRun || *dryRunLong
	isForce := *force || *forceLong

	fp := FileProcessor{
		DryRun:    isDryRun,
		Force:     isForce,
		Recursive: isRecursive,
	}

	err := fp.ProcessFiles(inputPath, pattern)
	if err != nil {
		log.Fatalf("Error processing files: %v", err)
	}

	fmt.Println("Files processed successfully.")
}
