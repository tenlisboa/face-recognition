package config

import (
	"log"
	"os"
	"path/filepath"
)

func getDirName() string {
	path, err := os.Getwd()
	if err != nil {
		log.Fatalf("Can't get the current dir: %v", err)
	}

	return path
}

var (
	ModelsDir = filepath.Join(getDirName(), "models")
	ImagesDir = filepath.Join(getDirName(), "images")
)
