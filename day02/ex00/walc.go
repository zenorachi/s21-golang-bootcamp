package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func WalkFunc(path string, info os.FileInfo, err error) error {
	if os.IsPermission(err) {
		return filepath.SkipDir
	} else if err != nil {
		return err
	}

	if info.Mode().IsRegular() && fl.ext != "" {
		fileExtension := filepath.Ext(path)
		if fileExtension == ("." + fl.ext) {
			fmt.Println(path)
		}
	} else {
		if s, _ := filepath.EvalSymlinks(path); fl.sl && s != path {
			if _, errExist := os.Stat(s); errExist == nil {
				fmt.Println(path, "->", s)
			} else {
				fmt.Println(path, "->", "[broken]")
			}
		}

		if fl.f && info.Mode().IsRegular() {
			fmt.Println(path)
		}

		if fl.d && info.Mode().IsDir() {
			fmt.Println(path)
		}
	}

	return nil
}
