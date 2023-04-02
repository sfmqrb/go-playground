package main

import (
	"flag"
	"fmt"
	"log"
	"path/filepath"
)

func main() {
	var srcDir string
	var distDir string
	var showHelp bool

	// Define command-line flags
	flag.StringVar(&srcDir, "src-dir", "", "source directory to move files from")
	flag.StringVar(&distDir, "dist-dir", "", "destination directory to move files to")
	flag.BoolVar(&showHelp, "help", false, "show help message")

	// Parse command-line flags
	flag.Parse()

	// Show help message if requested
	if showHelp {
		showUsage()
		return
	}

	// Check that source and destination directories were provided
	if srcDir == "" || distDir == "" {
		log.Fatal("Please provide source and destination directories using the --src-dir and --dist-dir flags.")
	}

	// Get absolute paths for source and destination directories
	srcDirAbs, err := filepath.Abs(srcDir)
	if err != nil {
		log.Fatalf("Error getting absolute path for source directory: %v", err)
	}
	distDirAbs, err := filepath.Abs(distDir)
	if err != nil {
		log.Fatalf("Error getting absolute path for destination directory: %v", err)
	}

	// Move files from source directory to destination directory
	err = YankFilesRecursively(srcDirAbs, distDirAbs)
	if err != nil {
		log.Fatalf("Error moving files: %v", err)
	}

	fmt.Println("Files moved successfully!")
}

func showUsage() {
	fmt.Println("Usage: mytool --src-dir=<source directory> --dist-dir=<destination directory>")
	flag.PrintDefaults()
}

