package etlutils

import (
	"fmt"
)

// Print prints firstN elements from channel ch (use firstN=0 to fetch everything)
func Print(ch <-chan interface{}, firstN int) {
	n := 0
	for i := range ch {
		if n > 0 && n > firstN {
			break
		}
		n++
		switch i := i.(type) {
		case string:
			fmt.Printf("%q", i)
		default:
			fmt.Printf("%T\t%v", i, i)
		}
	}
}

// PrintS prints firstN elements from string channel (use firstN=0 to fetch everything)
func PrintS(ch <-chan string, firstN int) {
	chi := make(chan interface{}, firstN)
	n := 0
	for i := range ch {
		if n > 0 && n > firstN {
			break
		}
		n++
		chi <- i
	}
	close(chi)
	Print(chi, 0)
}
