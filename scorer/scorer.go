// Package scorer calculate score for each file
package scorer

import (
	"io/ioutil"
	"log"
	"regexp"
)

// ScoringResult contains scoring results for one file
type ScoringResult struct {
	FilePath    string
	NumOfErrors int
}

var (
	sbacliError = regexp.MustCompile("^\\d\\d\\d\\d-\\d\\d-\\d\\d \\d\\d:\\d\\d:\\d\\d,\\d\\d\\d\\tsbacli\\tERROR\\(.*)")
)

// ScoreFile is score process implementation for one file
func ScoreFile(file string) ScoringResult {
	bytes, err := ioutil.ReadFile(file)
	if err != nil {
		log.Panic(err)
	}
	match := sbacliError.FindAllSubmatch(bytes, -1)
	if match == nil {
		return ScoringResult{file, 0}
	}
	return ScoringResult{file, len(match)}
}

// Scorer is a worker thread #n for scoring process.
func Scorer(n int, files <-chan string, processedCnt chan<- int, dmakerInChan chan<- ScoringResult) {
	log.Printf("Scorer #%d started", n)
	cnt := 0
	for file := range files {
		log.Printf("Scorer #%d: processing '%s'", n, file)
		dmakerInChan <- ScoreFile(file)
		cnt++
	}
	log.Printf("Scorer #%d ended, %d files processed", n, cnt)
	processedCnt <- cnt
}
