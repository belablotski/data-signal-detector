// Scan directory tree
package main

import (
	"html"
	"io/ioutil"
	"log"
	"path"
	"strings"
)

// Returns list of sub-directories in dataDir
func listDirs(dataDir string) (result []string) {
	files, err := ioutil.ReadDir(dataDir)

	if err != nil {
		log.Panic(err)
	}

	result = make([]string, 0, len(files))

	for _, f := range files {
		if f.IsDir() {
			result = append(result, f.Name())
		} else {
			log.Panicf("There is file '%s' in '%s' dir, expected sub-directories only", f.Name(), dataDir)
		}
	}

	return result
}

// Returns list of files (i.e. list of files in dataDir/subdir dir)
func listFiles(dataDir, subdir string, skipOfflineValidationRequests bool) (result []string) {
	files, err := ioutil.ReadDir(path.Join(dataDir, subdir))

	if err != nil {
		log.Panic(err)
	}

	result = make([]string, 0, len(files))

	for _, f := range files {
		if !f.IsDir() {
			theRequestId := html.UnescapeString(f.Name())
			if skipOfflineValidationRequests && strings.HasPrefix(theRequestId, "temp-") {
				log.Printf("Temp '%s'", theRequestId)
			} else {
				result = append(result, theRequestId)
			}
		} else {
			log.Panicf("There is dir '%s' in '%s' dir, expected files only", f.Name(), dataDir)
		}
	}

	return result
}
