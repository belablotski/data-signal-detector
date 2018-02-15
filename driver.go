package main

import (
	"log"
	"time"

	"github.com/beloblotskiy/data-signal-detector/scanner"
	"github.com/beloblotskiy/data-signal-detector/scorer"
)

func doScoring(files []string) {
	scorerPoolSize := 10
	if len(files) < scorerPoolSize {
		scorerPoolSize = 1
	}
	chanBufferSize := (len(files) / scorerPoolSize) + 1
	log.Printf("Scoring of %d files will be done in %d buffered threads with buffer size %d", len(files), scorerPoolSize, chanBufferSize)

	t1 := time.Now()
	scorerInChans := make([]chan string, scorerPoolSize)
	scorerEndChans := make([]chan int, scorerPoolSize)
	dmakerInChan := make(chan scorer.ScoringResult, len(files)/10)
	for i := 0; i < scorerPoolSize; i++ {
		scorerInChans[i] = make(chan string, chanBufferSize)
		scorerEndChans[i] = make(chan int)
		go scorer.Scorer(i, scorerInChans[i], scorerEndChans[i], dmakerInChan)
	}

	for i, file := range files {
		n := i % scorerPoolSize
		scorerInChans[n] <- file
	}

	for i := 0; i < scorerPoolSize; i++ {
		close(scorerInChans[i])
	}

	processedCnt := 0
	for i := 0; i < scorerPoolSize; i++ {
		processedCnt += <-scorerEndChans[i]
	}

	if processedCnt != len(files) {
		log.Panicf("Files count %d doesn't match number of processed files %d", len(files), processedCnt)
	}

	log.Printf("Scoring of %d files has beeen done in %s", len(files), time.Since(t1).String())
}

func main() {
	files, size := scanner.ListFiles("c:\\")
	log.Printf("Total number of files: %d, total size: %d", len(files), size)

	doScoring(files)
}
