package main

import (
	"encoding/json"
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

func ReadData(data []byte) *directory {
	var dir *directory
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
func virusFile(dir *directory) int {
	re := regexp.MustCompile(`\.hack$`)
	for _, v := range dir.Files {
		if re.FindStringIndex(v) != nil {
			return numberof_files(dir)
		}
	}
	virus_folders := 0
	for _, folder := range dir.Folders {
		virus_folders += virusFile(folder)
	}
	return virus_folders
}
