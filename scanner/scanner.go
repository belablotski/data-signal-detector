// Package scanner scans directory tree
package scanner

import (
	"io/ioutil"
	"log"
	"path"
	"strings"
	"sync"
	"time"
)

// ListFiles returns list of files and their total size in startDir and its sub-directiries
func ListFiles(startDir string) (files []string, size int64) {
	filesAndDirs, err := ioutil.ReadDir(startDir)

	if err != nil {
		if strings.Contains(err.Error(), "Access is denied") {
			log.Println(err)
		} else {
			log.Panicln(err)
		}
	}

	files = make([]string, 0, 100)

	for _, f := range filesAndDirs {
		if f.IsDir() {
			fl, sz := ListFiles(path.Join(startDir, f.Name()))
			files = append(files, fl...)
			size += sz
		} else {
			files = append(files, path.Join(startDir, f.Name()))
			size += f.Size()
		}
	}

	return files, size
}

// Scanner does file system scan, starting from specified folder and submits found files into output channel
func Scanner(startDir string) chan<- string {
	dirs := make(chan string, 200)
	files := make(chan<- string, 1000)

	dirs <- startDir

	lister := func(n int, dirs chan string) {
		for timeout := 0; timeout < 3; {
			select {
			case dir := <-dirs:
				filesAndDirs, err := ioutil.ReadDir(dir)
				if err != nil {
					if strings.Contains(err.Error(), "Access is denied") {
						log.Println(err)
					} else {
						log.Panicln(err)
					}
				}

				for _, f := range filesAndDirs {
					fpath := path.Join(startDir, f.Name())
					if f.IsDir() {
						dirs <- fpath
					} else {
						files <- fpath
					}
				}
			case <-time.After(1 * time.Second):
				timeout++
			}
		}
	}

	go func() {
		var wg sync.WaitGroup
		wg.Add(1)
		for i := 1; i <= 2; i++ {
			go lister(i, dirs)
		}
		wg.Wait()
		close(files)
	}()

	return files
}
