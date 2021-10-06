package main

import (
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	out := os.Stdout
	if !(len(os.Args) == 2 || len(os.Args) == 3) {
		panic("usage go run main.go . [-f]")
	}
	path := os.Args[1]
	printFiles := len(os.Args) == 3 && os.Args[2] == "-f"
	err := DirTree(out, path, printFiles)
	if err != nil {
		panic(err.Error())
	}
}

func DirTree(out io.Writer, path string, printFiles bool) error {

	return filepath.Walk(path, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			fmt.Printf("prevent panic by handling failure accessing a path %q: %v\n", path, err)
			return err
		}

		count := strings.Count(path, string(os.PathSeparator)) + 2

		fmt.Printf("%v%v%v path - %v\n", strings.Repeat("\t", count-2), "├───", info.Name(), path)

		return nil
	})
}
