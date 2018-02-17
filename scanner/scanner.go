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

// Scan does file system scan, starting from specified folder and submits found files into output channel
func Scan(startDir string) <-chan string {
	dirs := make(chan string, 100000)
	files := make(chan string, 100000)

	fslister := func(n int, wg *sync.WaitGroup) {
		log.Printf("FSLister #%d started", n)
		for timeout := 0; timeout < 3; {
			select {
			case dir := <-dirs:
				log.Printf("FSLister #%d procesing directory '%s'", n, dir)
				filesAndDirs, err := ioutil.ReadDir(dir)
				if err != nil {
					if strings.Contains(err.Error(), "Access is denied") {
						log.Println(err)
					} else {
						log.Panicln(err)
					}
				}

				for _, f := range filesAndDirs {
					fpath := path.Join(dir, f.Name())
					if f.IsDir() {
						dirs <- fpath
					} else {
						files <- fpath
					}
				}
				log.Printf("FSLister #%d xx", n)
			case <-time.After(1 * time.Second):
				timeout++
				log.Printf("FSLister #%d timeout #%d", n, timeout)
			}
		}
		wg.Done()
	}

	dirs <- startDir

	go func() {
		var wg sync.WaitGroup
		n := 1
		wg.Add(n)
		for i := 1; i <= n; i++ {
			go fslister(i, &wg)
		}
		wg.Wait()
		close(dirs)
		close(files)
	}()

	return files
}
