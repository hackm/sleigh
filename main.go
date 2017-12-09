package main

import (
	"fmt"
	"os"
	"sync"
	"time"

	flags "github.com/jessevdk/go-flags"
)

type Options struct {
	Room     string `short:"r" long:"room" description:"Room name" required:"true"`
	Password string `short:"p" long:"pass" description:"Password"`
	Listen   uint   `short:"l" long:"listen" description:"Using port number"`
}

var wg sync.WaitGroup
var options Options

func main() {
	var parser = flags.NewParser(&options, flags.Default)
	if _, err := parser.Parse(); err != nil {
		if flagsErr, ok := err.(*flags.Error); ok && flagsErr.Type == flags.ErrHelp {
			os.Exit(0)
		} else {
			os.Exit(1)
		}
	}
	fmt.Printf("Room: %s\n", options.Room)
	fmt.Printf("Password: %s\n", options.Password)
	fmt.Printf("Listen: %d\n", options.Listen)
	sleigh()
}

func sleigh() {
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
