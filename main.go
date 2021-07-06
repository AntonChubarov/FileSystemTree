package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	folderPath, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	fmt.Println(folderPath)

	files, err := ioutil.ReadDir(folderPath)
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		fmt.Println(f.Name())
	}

	file := files[4]

	fmt.Println(file)
}

//type FolderInfo struct {
//	name string
//	subfoldersCount int
//	filesCount int
//}
//
//type FileInfo struct {
//	name string
//	size int
//}
