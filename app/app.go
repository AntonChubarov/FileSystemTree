package app

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

func PrintFolder(path string, isInitial bool, isLast bool, prefix []string, fileTypes map[string]int) {
	folderName := GetFolderName(path)

	fmt.Println(GetTreePrefix(isInitial, isLast, prefix) + folderName + " " + DirInfo(path, fileTypes))

	// modify folder tree-prefix
	if !isLast {
		prefix = append(prefix, linePrefix)
	} else if !isInitial {
		prefix = append(prefix, emptyPrefix)
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
			PrintFolder(filepath.Join(path, fsoName), false, i == len(fso)-1, prefix, fileTypes)
		} else {
			PrintFile(f, i == len(fso)-1, prefix, fileTypes)
		}
	}
}

func PrintFile(file os.DirEntry, isLast bool, prefix []string, fileTypes map[string]int) {
	name := file.Name()
	extension, _ := GetFileExtension(name)
	if _, ok := fileTypes[extension]; ok {
		if isLast {
			fmt.Println(PrefixToString(prefix) + lastPrefix + name + " " + FileSizeInfo(file))
			prefix = append(prefix, emptyPrefix)
		} else {
			fmt.Println(PrefixToString(prefix) + nonLastPrefix + name + " " + FileSizeInfo(file))
			prefix = append(prefix, linePrefix)
		}
	}
}

func PrefixToString(prefix []string) string {
	var s string
	for _, p := range prefix {
		s += p
	}
	return s
}

func DirInfo(path string, fileTypes map[string]int) (output string) {
	folderCount, fileCount := DirCount(path, fileTypes)
	output += "("

	if folderCount == 0 && fileCount == 0 {
		output += "empty)"
		return
	}

	if folderCount == 1 {
		output += "1 folder"
	} else if folderCount > 1 {
		output += fmt.Sprintf("%v", folderCount) + " folders"
	}

	if folderCount != 0 && fileCount != 0 {
		output += ", "
	}

	if fileCount == 1 {
		output += "1 file"
	} else if fileCount > 1 {
		output += fmt.Sprintf("%v", fileCount) + " files"
	}

	output += ")"
	return output
}

func DirCount(path string, fileTypes map[string]int) (folderCount int, fileCount int) {
	fso, err := os.ReadDir(path)
	if err != nil {
		log.Println(err)
	}
	for _, f := range fso {
		if f.IsDir() {
			folderCount++
		} else if extension, ok := GetFileExtension(f.Name()); ok {
			if _, ok := fileTypes[extension]; ok {
				fileCount++
			}
		}
	}
	return
}

func FileSizeInfo(file os.DirEntry) (output string) {
	info, err := file.Info()
	if err != nil {
		log.Println(err)
	}
	size := info.Size()
	if size >= 1024 && size < 1024*1024 {
		output = "(" + fmt.Sprintf("%.2f", float32(size)/1024) + " KB)"
	} else if size >= 1024*1024 {
		output = "(" + fmt.Sprintf("%.2f", float32(size)/1024/1024) + " MB)"
	} else {
		output = "(" + fmt.Sprintf("%v", size) + " Bytes)"
	}
	return
}

func GetFileExtension(fileName string) (string, bool) {
	if strings.Contains(fileName, ".") &&
		strings.Count(fileName, ".") == 1 &&
		strings.Index(fileName, ".") != 0 {
		slice := strings.Split(fileName, ".")
		return "." + slice[len(slice)-1], true
	}
	return "", false
}

func GetFolderName(path string) (folderName string) {
	var ss []string
	if runtime.GOOS == "windows" {
		ss = strings.Split(path, "\\")
	} else {
		ss = strings.Split(path, "/")
	}
	folderName = ss[len(ss)-1]
	return
}

func GetTreePrefix(isInitial bool, isLast bool, prefix []string) (treePrefix string) {
	if isLast && isInitial {
		treePrefix = PrefixToString(prefix)
	} else if isLast && !isInitial {
		treePrefix = PrefixToString(prefix) + lastPrefix
	} else {
		treePrefix = PrefixToString(prefix) + nonLastPrefix
	}
	return
}
