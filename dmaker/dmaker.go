// Package dmaker is a decision maker components, which works after scoring
package dmaker

import (
	"log"
	"sync"

	"github.com/beloblotskiy/data-signal-detector/scorer"
)

// Decide makes a boolean decision about scored object, based on its score
func decide(score scorer.ScoringResult) bool {
	return score.NumOfErrors > 0
}

// Decide is a multi-thread decision maker
func Decide(nWorkers int, scores <-chan scorer.ScoringResult) <-chan scorer.ScoringResult {
	trueScores := make(chan scorer.ScoringResult, 100)

	decisionMaker := func(n int, wg *sync.WaitGroup) {
		defer wg.Done()
		log.Printf("Decision maker #%d starts", n)
		cnt := 0
		for score := range scores {
			if decide(score) {
				trueScores <- score
			}
			cnt++
		}
		log.Printf("Decision maker #%d ends, processed %d results", n, cnt)
	}

	go func() {
		var wg sync.WaitGroup
		wg.Add(nWorkers)
		for i := 1; i <= nWorkers; i++ {
			go decisionMaker(i, &wg)
		}
		wg.Wait()
		close(trueScores)
	}()

	return trueScores
}
