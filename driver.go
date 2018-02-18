package main

import (
	"github.com/beloblotskiy/data-signal-detector/dmaker"
	"github.com/beloblotskiy/data-signal-detector/etlutils"
	"github.com/beloblotskiy/data-signal-detector/scanner"
	"github.com/beloblotskiy/data-signal-detector/scorer"
)

func main() {
	etlutils.PrintSR(dmaker.Decide(1, scorer.Score(2, scanner.Scan(".\\test_data"))))
}
