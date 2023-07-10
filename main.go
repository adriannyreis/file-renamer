package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/dlclark/regexp2"
)

var rgx *regexp2.Regexp
var rgxSplit *regexp2.Regexp

func main() {
	rgx = regexp2.MustCompile(`(?i)(?!(\.))\s(\w+)(?=(\.))|(?!(\.))\s(\w+)$`, 0)
	rgxSplit = regexp2.MustCompile(`(?i)\s`, 0)

	// Folder/File name to be renamed
	originalPath := "./doc/"

	RenameFilesAndFolders(originalPath)
}

func RenameFilesAndFolders(data string) {
	files, _ := ioutil.ReadDir(data)

	for _, e := range files {
		temporaryPath := data
		currentName := temporaryPath + e.Name()

		if strings.Contains(e.Name(), "png") {
			continue
		}

		namePath, err := rgx.Replace(e.Name(), "", -1, -1)
		if err != nil {
			fmt.Println("Error trying to replace file's or folder's name", err)
		}

		nameWithoutSpaces, err := rgxSplit.Replace(namePath, "-", -1, -1)
		if err != nil {
			fmt.Println("Error trying to replace file's or folder's name", err)
		}

		nameWithoutAccents := removeAcentos(nameWithoutSpaces)

		temporaryPath = temporaryPath + strings.ToLower(nameWithoutAccents)
		err = os.Rename(currentName, temporaryPath)
		if err != nil {
			fmt.Println("Error renaming file", err)
		}

		if e.IsDir() {
			RenameFilesAndFolders(data + strings.ToLower(nameWithoutAccents) + "/")
		}
	}
}

func removeAcentos(name string) string {
	mapper := map[int]rune{226: 97, 194: 97, 224: 97, 192: 97, 225: 97, 193: 97, 227: 97, 195: 97, 234: 101, 202: 101, 232: 101, 200: 101, 233: 101, 201: 101, 238: 105, 206: 105, 236: 105, 204: 105, 237: 105, 205: 105, 245: 111, 213: 111, 244: 111, 212: 111, 242: 111, 210: 111, 243: 111, 211: 111, 252: 117, 220: 117, 251: 117, 219: 117, 250: 117, 218: 117, 249: 117, 217: 117, 231: 99, 199: 99}
	result := []rune{}

	runes := []rune(name)

	for _, word := range runes {
		if word == 771 || word == 770 || word == 769 || word == 807 {
			continue
		}

		hasStringWithAccent := mapper[int(word)]
		if hasStringWithAccent != 0 {
			result = append(result, mapper[int(word)])
		} else {
			result = append(result, word)
		}
	}
	return string(result)
}
