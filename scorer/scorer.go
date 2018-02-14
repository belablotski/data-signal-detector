package scorer

import (
	"log"
)

// Score is score process implementation for one file
func Score(file string) {
	log.Printf("Scoring file %s", file)
}

// Scorer is a worker thread for scoring process.
func Scorer(n int, files <-chan string, processedCnt chan<- int) {
	log.Printf("Scorer #%d started", n)
	cnt := 0
	for file := range files {
		log.Printf("Scorer #%d: processing '%s'", n, file)
		cnt++
	}
	log.Printf("Scorer #%d ended, %d files processed", n, cnt)
	processedCnt <- cnt
}
