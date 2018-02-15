// Package dmaker is a decision maker components, which works after scoring
package dmaker

import (
	"github.com/beloblotskiy/data-signal-detector/scorer"
)

// DecisionMaker is a decision maker #n worker process
func DecisionMaker(n int, dmakerInChan <-chan scorer.ScoringResult, processedCnt chan<- int) {

}
