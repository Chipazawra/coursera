package main

import (
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"sort"
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

	root := path
	lasts := make(map[string]bool)
	return filepath.Walk(path, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			fmt.Printf("prevent panic by handling failure accessing a path %q: %v\n", path, err)
			return err
		}

		if path == root {
			return nil
		}

		pseudo := "├───"

		if last, _ := readlastNameOfDir(filepath.Dir(path), printFiles); last == info.Name() {
			pseudo = "└───"
			lasts[last] = true
		}

		pseudoPrefix := ""

		for _, v := range strings.Split(filepath.Dir(path), string(os.PathSeparator))[1:] {
			if _, exist := lasts[v]; exist {
				pseudoPrefix += "\t"
			} else {
				pseudoPrefix += "│\t"
			}
		}

		if info.IsDir() {
			fmt.Fprintf(out, "%v%v%v\n", pseudoPrefix, pseudo, info.Name())
		} else if printFiles {
			if info.Size() == 0 {
				fmt.Fprintf(out, "%v%v%v (empty)\n", pseudoPrefix, pseudo, info.Name())
			} else {
				fmt.Fprintf(out, "%v%v%v (%vb)\n", pseudoPrefix, pseudo, info.Name(), info.Size())
			}
		}

		return nil
	})

}

func readlastNameOfDir(dirname string, readfiles bool) (string, error) {
	f, err := os.Open(dirname)
	if err != nil {
		return "", err
	}
	names, err := f.Readdirnames(-1)

	if !readfiles {
		for i, name := range names {
			filename := filepath.Join(dirname, name)
			fileInfo, err := os.Lstat(filename)
			if err != nil {
				panic(err)
			} else if !fileInfo.IsDir() {
				copy(names[i:], names[i+1:])
				names[len(names)-1] = ""
				names = names[:len(names)-1]
			}
		}
	}

	f.Close()
	if err != nil {
		return "", err
	}

	if len(names) == 0 {
		return "", nil
	}

	sort.Strings(names)
	return names[len(names)-1], nil
}
