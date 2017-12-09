package main

import (
	"encoding/json"
	"fmt"
	"os"
	"os/signal"
	"path/filepath"
	"strings"
	"sync"
	"syscall"

	"github.com/fsnotify/fsnotify"
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

func ignore(evt Event) bool {
	return false
}

func sleigh() {
	showTextLogo()
	done := make(chan os.Signal, 1)
	signal.Notify(done, syscall.SIGINT)

	hostname, err := os.Hostname()
	if err != nil {
		fmt.Println("Cannot get hostname.")
		os.Exit(1)
	}
	wd, err := os.Getwd()
	if err != nil {
		fmt.Println("Cannot get current working directory.")
		os.Exit(1)
	}
	items, err := GetItems(wd, wd, ignore)
	if err != nil {
		fmt.Println("Cannot get items in current working directory.")
		os.Exit(1)
	}
	hey := Hey{
		Hostname: hostname,
		Items:    items,
	}

	conn := NewConn("8986")

	tracker := NewTracker(wd, ignore)
	defer tracker.Close()

	differ := NewDiffer(hostname, int(options.Listen), options.Room)
	defer differ.Close()

	go func() {
		for {
			select {
			case d := <-conn.Datagram:
				var h Hey
				var n Notification
				str := string(d.Payload)

				if strings.Contains(str, `"type"`) {
					_ = json.Unmarshal(d.Payload, &n)
					fmt.Println("Notification")
					differ.Notifications <- n
				} else if err := json.Unmarshal(d.Payload, &h); err == nil {
					fmt.Println("Hey")
					for _, item := range h.Items {
						path := filepath.Join(wd, item.RelPath)
						fmt.Println(path)

						info, err := os.Stat(path)
						if err != nil {
							differ.Notifications <- Notification{
								Hostname: h.Hostname,
								Event:    fsnotify.Write,
								Type:     File,
								Path:     item.RelPath,
								ModTime:  item.ModTime,
							}
							continue
						}
						checksum, err := GetChecksum(path)
						if err != nil {
							fmt.Printf("ERROR: %v\n", err)
							continue
						}
						modtime := info.ModTime().UnixNano()
						if checksum != item.Checksum {
							fmt.Printf("%v != %v\n", checksum, item.Checksum)
							fmt.Printf("%v , %v\n", item.ModTime, modtime)
							if item.ModTime > modtime {
								differ.Notifications <- Notification{
									Hostname: h.Hostname,
									Event:    fsnotify.Write,
									Type:     File,
									Path:     item.RelPath,
									ModTime:  item.ModTime,
								}
							} else {
								conn.Notify(Notification{
									Hostname: hostname,
									Event:    fsnotify.Write,
									Type:     File,
									Path:     item.RelPath,
									ModTime:  modtime,
								})
							}
						}
					}
					for _, local := range items {
						hit := false
						for _, remote := range h.Items {
							if local.RelPath == remote.RelPath {
								hit = true
							}
						}
						if hit == false {
							path := filepath.Join(wd, local.RelPath)
							info, _ := os.Stat(path)

							conn.Notify(Notification{
								Hostname: hostname,
								Event:    fsnotify.Write,
								Type:     File,
								Path:     local.RelPath,
								ModTime:  info.ModTime().UnixNano(),
							})
						}
					}
				}
			case evt := <-tracker.Events:
				fmt.Println("Tracker")
				info, err := os.Stat(evt.FullPath)
				if err != nil {
					fmt.Printf("ERROR: %v\n", err)
					continue
				}
				var itemType = File
				if evt.Dir {
					itemType = Dir
				}
				conn.Notify(Notification{
					Hostname: hostname,
					Event:    evt.Op,
					Type:     itemType,
					Path:     evt.RelPath,
					ModTime:  info.ModTime().UnixNano(),
				})
			case err := <-tracker.Errors:
				fmt.Printf("ERROR: %v\n", err)
			case err := <-differ.Errors:
				fmt.Printf("ERROR: %v\n", err)
			case err := <-conn.Errors:
				fmt.Printf("ERROR: %v\n", err)
			}
		}
	}()

	conn.Listen()
	differ.Start()
	tracker.Start()
	conn.Hey(hey)

	// /*
	// 	*** mock work ***
	// 	spawn some goroutines to do arbitrary work, updating their
	// 	respective progress bars as they see fit
	// */
	// progressChannel := make(chan int)

	// go showProgress("ProgressBar", progressChannel, 100)

	// wg.Add(1)
	// // do something asyn that we can get updates upon
	// // every time an update comes in, tell the bar to re-draw
	// // this could be based on transferred bytes or similar
	// for i := 0; i <= 100; i++ {
	// 	progressChannel <- i
	// 	time.Sleep(time.Millisecond * 10)
	// }
	// close(progressChannel)
	// wg.Done()
	<-done
	fmt.Println("Bye!")
}
