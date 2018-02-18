// Package scanner scans directory tree
package scanner

import (
	"io/ioutil"
	"log"
	"path"
	"strings"
)

func listFiles(startDir string, files chan<- string) {
	filesAndDirs, err := ioutil.ReadDir(startDir)
	if err != nil {
		if strings.Contains(err.Error(), "Access is denied") {
			log.Println(err)
		} else {
			log.Panicln(err)
		}
	}

	for _, f := range filesAndDirs {
		p := path.Join(startDir, f.Name())
		if f.IsDir() {
			listFiles(p, files)
		} else {
			files <- p
		}
	}
}

// Scan does file system scan, starting from specified folder and submits found files into output channel
func Scan(startDir string) <-chan string {
	files := make(chan string, 100)

	go func() {
		log.Println("File system scanner starts")
		listFiles(startDir, files)
		close(files)
		log.Println("File system scanner ends")
	}()

	return files
}
