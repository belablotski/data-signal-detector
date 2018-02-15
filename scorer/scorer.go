// Package scorer calculate score for each file
package scorer

import (
	"io/ioutil"
	"log"
	"regexp"
)

// ScoringResult contains scoring results for one file
type ScoringResult struct {
	file        string
	numOfErrors int
}

var (
	sbacliError = regexp.MustCompile("^\\d\\d\\d\\d-\\d\\d-\\d\\d \\d\\d:\\d\\d:\\d\\d,\\d\\d\\d\\tsbacli\\tERROR\\(.*)")
)

// ScoreFile is score process implementation for one file
func ScoreFile(file string) {
	bytes, err := ioutil.ReadFile(file)
	if err != nil {
		log.Panic(err)
	}
	match := sbacliError.FindAllSubmatch(bytes, -1)
	if match == nil {

	}
}

// Scorer is a worker thread for scoring process.
func Scorer(n int, files <-chan string, processedCnt chan<- int) {
	log.Printf("Scorer #%d started", n)
	cnt := 0
	for file := range files {
		log.Printf("Scorer #%d: processing '%s'", n, file)
		ScoreFile(file)
		cnt++
	}
	log.Printf("Scorer #%d ended, %d files processed", n, cnt)
	processedCnt <- cnt
}
