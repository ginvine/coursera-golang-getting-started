package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sort"
)

func main() {
	out := os.Stdout
	if !(len(os.Args) == 2 || len(os.Args) == 3) {
		panic("usage go run main.go . [-f]")
	}
	path := os.Args[1]
	printFiles := len(os.Args) == 3 && os.Args[2] == "-f"
	err := dirTree(out, path, printFiles)
	if err != nil {
		panic(err.Error())
	}
}

func dirTree(out io.Writer, path string, printFiles bool) error {
	return dirTreeHelper("", out, path, printFiles)
}

func dirTreeHelper(prefix string, out io.Writer, path string, printFiles bool) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	fileInfos, err := file.Readdir(0)

	sort.Slice(fileInfos, func(i, j int) bool { return fileInfos[i].Name() < fileInfos[j].Name() })

	if !printFiles {
		fileInfos = removeFiles(fileInfos)
	}

	for i, f := range fileInfos {
		p := "%s├───%s\n"
		if i == len(fileInfos)-1 {
			p = "%s└───%s\n"
		}

		s := f.Name()
		if !f.IsDir() {
			s = fmt.Sprintf("%s (%db)", s, f.Size())
			if f.Size() == 0 {
				s = fmt.Sprintf("%s (empty)", f.Name())
			}
		}

		fmt.Fprintf(out, p, prefix, s)

		if f.IsDir() {
			pr := prefix + "│\t"
			if i == len(fileInfos)-1 {
				pr = prefix + "\t"
			}
			dirTreeHelper(pr, out, filepath.Join(path, f.Name()), printFiles)
		}
	}
	return nil
}

func removeFiles(fi []os.FileInfo) []os.FileInfo {
	res := []os.FileInfo{}
	for _, f := range fi {
		if f.IsDir() {
			res = append(res, f)
		}
	}
	return res
}
