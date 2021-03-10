package main

import (
	"os"
	"path/filepath"
)

func init() {

	logpath := filepath.Join(".", "logs")
	if _, err := os.Stat(logpath); os.IsNotExist(err) {
		os.MkdirAll(logpath, os.ModePerm)
	}

	jobpath := filepath.Join(".", "jobs")
	if _, err := os.Stat(jobpath); os.IsNotExist(err) {
		os.MkdirAll(jobpath, os.ModePerm)
	}

	GetConfig()
}
