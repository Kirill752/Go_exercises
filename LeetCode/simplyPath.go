package main

import (
	"strings"
)

func simplifyPath(path string) string {
	dirs := strings.Split(path, "/")
	stack := []string{""}
	for i := 0; i < len(dirs); i++ {
		if dirs[i] == "/" || dirs[i] == "." || dirs[i] == "" {
			continue
		}
		if dirs[i] == ".." {
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			}
			continue
		}
		stack = append(stack, dirs[i])
	}
	if len(stack) <= 1 {
		return "/" + strings.Join(stack, "/")
	}
	if stack[0] != "" {
		return "/" + strings.Join(stack, "/")
	}
	return strings.Join(stack, "/")
}
