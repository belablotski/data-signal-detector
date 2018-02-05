package main

import (
	"log"
	"github.com/beloblotskiy/data-signal-detector/scanner"
)

func main() {
	files, size := scanner.ListFiles("c:\\")
	for _, file := range files {
		log.Println(file)
	}
	log.Printf("Total number of files: %d, total size: %d", len(files), size)
}
