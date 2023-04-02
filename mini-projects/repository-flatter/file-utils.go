package main 

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func MoveFile(sourcePath string, destDir string) error {
	// Open the source file for reading
	source, err := os.Open(sourcePath)
	if err != nil {
		return err
	}
	defer source.Close()

	if err := os.MkdirAll(destDir, os.ModePerm); err != nil {
		return err
	}

	// Create the destination file path
	destPath := filepath.Join(destDir, filepath.Base(sourcePath))

	// Open the destination file for writing
	dest, err := os.Create(destPath)
	if err != nil {
		return err
	}
	defer dest.Close()

	// Copy the contents of the source file to the destination file
	if _, err := io.Copy(dest, source); err != nil {
		return err
	}

	// Remove the source file
	if err := os.Remove(sourcePath); err != nil {
		return err
	}

	fmt.Printf("Moved file from %s to %s\n", sourcePath, destPath)

	return nil
}

