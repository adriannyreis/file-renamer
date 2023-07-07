package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/dlclark/regexp2"
)

var rgx *regexp2.Regexp

func main() {
	rgx = regexp2.MustCompile(`(?i)(?!(\.))\s(\w+)(?=(\.))|(?!(\.))\s(\w+)$`, 0)

	// Folder/File name to be renamed
	originalPath := "./doc/"

	RenameFilesAndFolders(originalPath)
}

func RenameFilesAndFolders(data string) {
	files, _ := ioutil.ReadDir(data)

	for _, e := range files {
		temporaryPath := data
		currentName := temporaryPath + e.Name()

		name, err := rgx.Replace(e.Name(), "", -1, -1)
		if err != nil {
			fmt.Println("Error trying to replace file's or folder's name", err)
		}

		temporaryPath = temporaryPath + name
		err = os.Rename(currentName, temporaryPath)
		if err != nil {
			fmt.Println("Error renaming file", err)
		}

		if e.IsDir() {
			RenameFilesAndFolders(data + name + "/")
		}
	}
}

// \s(\w+)(\.)
// (?i)(?!(\.))\s+\S*.+?(?=(\.))|(?!(\.))\s+\S*
