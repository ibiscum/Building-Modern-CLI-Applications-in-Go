package main

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
)

func walking() {
	workingDir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	dir1 := filepath.Join(workingDir, "dir1")
	err = filepath.WalkDir(dir1, func(path string, d fs.DirEntry, err error) error {
		if !d.IsDir() {
			contents, err := os.ReadFile(path)
			if err != nil {
				return err
			}
			fmt.Printf("%s -> %s\n", d.Name(), string(contents))
		}
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}
}
