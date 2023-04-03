package main

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"runtime/debug"
	"strings"
)

func YankFile(sourcePath string, destDir string, fileName string) error {
	// Open the source file for reading
	source, err := os.Open(sourcePath)
	if err != nil {
		debug.PrintStack()
		return err
	}
	defer source.Close()

	// Create the destination file path
	destPath := filepath.Join(destDir, fileName)

	// Open the destination file for writing
	dest, err := os.Create(destPath)
	if err != nil {
		debug.PrintStack()
		return err
	}
	defer dest.Close()

	// Copy the contents of the source file to the destination file
	if _, err := io.Copy(dest, source); err != nil {
		debug.PrintStack()
		return err
	}

	return nil
}


func YankFilesRecursively(srcDir string, dstDir string) error {
	// Get a list of all files and directories in the source directory
	files, err := ioutil.ReadDir(srcDir)
	if err != nil {
		return err
	}

	// Loop through each file and directory in the source directory
	for _, file := range files {
		srcPath := filepath.Join(srcDir, file.Name())
		// dstPath := filepath.Join(dstDir, strings.ReplaceAll(srcPath), "/", "-"))

		if file.IsDir() {
			// If the file is a directory, recursively move its contents to the destination directory
			if err := YankFilesRecursively(srcPath, dstDir); err != nil {
				return err
			}
		} else {
			// If the file is a regular file, move it to the destination directory with the new name
			fileName := strings.ReplaceAll(srcPath, "/", "@")
			if err := YankFile(srcPath, dstDir, fileName); err != nil {
				return err
			}
		}
	}

	return nil
}


func CreateJsonFileFromDir(dirPath, outputPath string) error {
	// Create a map to hold the file contents
	fileContents := make(map[string]string)

	// Read the contents of all files in the directory
	files, err := ioutil.ReadDir(dirPath)
	if err != nil {
		return err
	}

	for _, file := range files {
		if file.IsDir() {
			continue
		}

		filename := filepath.Join(dirPath, file.Name())
		fileData, err := ioutil.ReadFile(filename)
		if err != nil {
			return err
		}

		// Store the file contents in the map, with the file name as the key
		fileContents[file.Name()] = string(fileData)
	}

	// Encode the map as JSON
	jsonData, err := json.Marshal(fileContents)
	if err != nil {
		return err
	}

	// Write the JSON data to file
	err = ioutil.WriteFile(outputPath, jsonData, 0644)
	if err != nil {
		return err
	}

	return nil
}
