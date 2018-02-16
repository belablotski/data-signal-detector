// Package dmaker is a decision maker components, which works after scoring
package dmaker

import (
	"github.com/beloblotskiy/data-signal-detector/scorer"
)

// Decide makes a boolean decision about scored object, based on its score
func Decide(score scorer.ScoringResult) bool {
	return score.NumOfErrors > 0
}

// DecisionMaker is a decision maker #n worker process
func DecisionMaker(n int, dmakerInChan <-chan scorer.ScoringResult, dmakerOutChan chan<- scorer.ScoringResult) {

}
