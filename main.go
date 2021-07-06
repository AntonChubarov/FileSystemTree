package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

const (
	lastPrefix    = "\u2514" + "\u2500" + "\u2500" + " "
	nonLastPrefix = "\u251C" + "\u2500" + "\u2500" + " "
	linePrefix    = "\u2502" + "   "
	emptyPrefix   = "    "
)

func main() {

	currentFolder, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	fmt.Println(currentFolder)

	PrintFolder(currentFolder, true, make([]string, 1))
}

func PrintFolder(path string, isLast bool, prefix []string) {
	// get folder name
	var ss []string
	if runtime.GOOS == "windows" {
		ss = strings.Split(path, "\\")
	} else {
		ss = strings.Split(path, "/")
	}
	folderName := ss[len(ss)-1]

	// modify folder tree-prefix and print
	if isLast {
		fmt.Println(PrefixToString(prefix) + lastPrefix + folderName)
		prefix = append(prefix, emptyPrefix)
	} else {
		fmt.Println(PrefixToString(prefix) + nonLastPrefix + folderName)
		prefix = append(prefix, linePrefix)
	}

	// get all folders and fso in current folder
	fso, err := os.ReadDir(path)
	if err != nil {
		log.Println(err)
	}

	// call print functions for each folders and/or file
	for i := range fso {
		f := fso[i]
		fsoName := f.Name()
		if f.IsDir() {
			PrintFolder(filepath.Join(path, fsoName), i == len(fso)-1, prefix)
		} else {
			PrintFile(fsoName, i == len(fso)-1, prefix)
		}
	}
}

func PrintFile(name string, isLast bool, prefix []string) {
	// print line for current file
	if isLast {
		fmt.Println(PrefixToString(prefix) + lastPrefix + name)
		prefix = append(prefix, emptyPrefix)
	} else {
		fmt.Println(PrefixToString(prefix) + nonLastPrefix + name)
		prefix = append(prefix, linePrefix)
	}
}

func PrefixToString(prefix []string) string {
	var s string
	for _, p := range prefix {
		s += p
	}
	return s
}
