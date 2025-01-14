package main

import (
	"encoding/json"
	"fmt"
	"os"
	"regexp"
)

type directory struct {
	Dir     string
	Folders []*directory
	Files   []string
}

func ReadJson(source string) directory {
	var dir directory
	file, _ := os.ReadFile(source)
	json.Unmarshal(file, &dir)
	return dir
}

func ReadData(data []byte) directory {
	var dir directory
	json.Unmarshal(data, &dir)
	return dir
}

func numberof_files(dir_ptr *directory) int {
	totalFilesCount := len(dir_ptr.Files)
	for _, folder := range dir_ptr.Folders {
		totalFilesCount += numberof_files(folder)
	}
	return totalFilesCount
}
func virusFile(dir directory, virusCount int, virusDetected bool) int {
	re := regexp.MustCompile(".hack$")
	for _, v := range dir.Files {
		if re.FindStringIndex(v) != nil {
			virusDetected = true
			virusCount += len(dir.Files)
			fmt.Printf("virus_files is %v\n", dir.Files)
		}
	}
	if virusDetected {
		virusCount = numberof_files(&dir)
	} else {
		for _, folder := range dir.Folders {
			virusCount = virusFile(*folder, virusCount, virusDetected)
		}
	}
	return virusCount
}
