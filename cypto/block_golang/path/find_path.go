package path

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func PrintFilePath(path string) {
	fileInfo, err := os.Stat(path)
	if err != nil {
		log.Printf("Failed to read file info for %s : %v", path, err)
	}
	if fileInfo.IsDir() {
		dirEntries, err := os.ReadDir(path)
		if err != nil {
			log.Printf("Failed to read dir entries for %s : %v", dirEntries, err)
		}
		for _, entry := range dirEntries {
			PrintFilePath(filepath.Join(path, entry.Name()))
		}
	} else {
		// fmt.Printf("File %s has %d size \n", path, fileInfo.Size())
		_, state, ok := strings.Cut(path, ".")
		if ok {
			// fmt.Println(state)
			if state == "txt" {
				fmt.Println(state)
			}
		}
	}
}
