package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
)

func main() {
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	dirs := DirsWith(wd, "\\.go$")

	for _, d := range dirs {
		if !strings.Contains(d, "/vendor/") {
			relPath := strings.Replace(d, wd, ".", 1)
			cmd := exec.Command("go", "install", relPath)
			out, err := cmd.CombinedOutput()
			outStr := string(out)
			if err != nil {
				if !strings.Contains(outStr, "no non-test Go files") && !strings.Contains(outStr, "build constraints exclude all Go files") {
					fmt.Println(fmt.Sprintf("building %s finished with error %s", relPath, err.Error()))
					fmt.Printf(string(out))
					os.Exit(1)
				}
			} else {
				fmt.Println(fmt.Sprintf("building %s finished ok", relPath))
			}
		}
	}
}

// Contains checks if a string is contained in a slice
func Contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

// DirsWith finds all the directories in the root (including the root) containing files matching regex
func DirsWith(root, mask string) []string {
	var dirs []string
	filepath.Walk(root, func(path string, f os.FileInfo, _ error) error {
		if !f.IsDir() {
			r, err := regexp.MatchString(mask, f.Name())
			if err == nil && r {
				d := filepath.Dir(path)
				if !Contains(dirs, d) {
					dirs = append(dirs, filepath.Dir(path))
				}
			}
		}
		return nil
	})
	return dirs
}
