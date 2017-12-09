package main

import (
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	showTextLogo()
	showLogo()
	/*
		*** mock work ***
		spawn some goroutines to do arbitrary work, updating their
		respective progress bars as they see fit
	*/
	progressChannel := make(chan int)

	go showProgress("ProgressBar", progressChannel, 100)

	wg.Add(1)
	// do something asyn that we can get updates upon
	// every time an update comes in, tell the bar to re-draw
	// this could be based on transferred bytes or similar
	for i := 0; i <= 100; i++ {
		progressChannel <- i
		time.Sleep(time.Millisecond * 10)
	}
	close(progressChannel)
	wg.Done()
}
