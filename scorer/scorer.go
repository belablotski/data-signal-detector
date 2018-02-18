// Package scorer calculate score for each file
package scorer

import (
	"io/ioutil"
	"log"
	"regexp"
	"sync"
)

// ScoringResult contains scoring results for one file
type ScoringResult struct {
	FilePath    string
	NumOfErrors int
}

var (
	sbacliError = regexp.MustCompile("^\\d\\d\\d\\d-\\d\\d-\\d\\d \\d\\d:\\d\\d:\\d\\d,\\d\\d\\d\\tsbacli\\tERROR\\t(.*)")
)

func scoreFile(file string) ScoringResult {
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

// Score does parallel scoring calculations for files from input channel
func Score(nWorkers int, files <-chan string) <-chan ScoringResult {
	scores := make(chan ScoringResult, 100)

	scorer := func(n int, wg *sync.WaitGroup) {
		defer wg.Done()
		log.Printf("Scorer #%d starts", n)
		for file := range files {
			scores <- scoreFile(file)
		}
		log.Printf("Scorer #%d ends", n)
	}

	go func() {
		var wg sync.WaitGroup
		wg.Add(nWorkers)
		for i := 1; i <= nWorkers; i++ {
			go scorer(i, &wg)
		}
		wg.Wait()
		close(scores)
	}()

	return scores
}
