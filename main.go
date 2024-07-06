package main

import (
	"flag"
	"fmt"
)

func main() {
	var (
		dFlag = flag.Bool("d", false, "Print the day overview")
		wFlag = flag.Bool("w", false, "Print the week overview")
		mFlag = flag.Bool("m", false, "Print the month overview")
	)
	flag.Parse()

	switch {
	case *dFlag:
	case *wFlag:
	case *mFlag:
	default:
		startTimer()
	}
}

func startTimer() {
	createInitialFile(true)
	fmt.Println("deepwork: timer starting, good luck!")
}
