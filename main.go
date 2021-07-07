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
		fmt.Println(PrefixToString(prefix) + lastPrefix + folderName + " " + DirInfo(path))
		prefix = append(prefix, emptyPrefix)
	} else {
		fmt.Println(PrefixToString(prefix) + nonLastPrefix + folderName + " " + DirInfo(path))
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
			PrintFile(f, i == len(fso)-1, prefix)
		}
	}
}

func PrintFile(file os.DirEntry, isLast bool, prefix []string) {
	// print line for current file
	info, err := file.Info()
	if err != nil {
		log.Println(err)
	}
	if isLast {
		fmt.Println(PrefixToString(prefix) + lastPrefix + info.Name() + " " + FileSizeInfo(file))
		prefix = append(prefix, emptyPrefix)
	} else {
		fmt.Println(PrefixToString(prefix) + nonLastPrefix + info.Name() + " " + FileSizeInfo(file))
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

func DirInfo(path string) string {
	folderCount, fileCount := DirCount(path)
	output := "(" + fmt.Sprintf("%v", folderCount) + " folders, " + fmt.Sprintf("%v", fileCount) + " files)"
	return output
}

func DirCount(path string) (folderCount int, fileCount int) {
	fso, err := os.ReadDir(path)
	if err != nil {
		log.Println(err)
	}
	for _, f := range fso {
		if f.IsDir() {
			folderCount++
		} else {
			fileCount++
		}
	}
	return
}

func FileSizeInfo(file os.DirEntry) (output string) {
	info, err := file.Info()
	if err != nil {
		log.Println(err)
	}
	output += "("
	output += fmt.Sprintf("%v", info.Size())
	output += " Bytes"
	output += ")"
	return
}
