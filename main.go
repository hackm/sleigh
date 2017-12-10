package main

import (
	"encoding/json"
	"os"
	"os/signal"
	"path/filepath"
	"strings"
	"sync"
	"syscall"
	"time"

	"github.com/fatih/color"
	"github.com/fsnotify/fsnotify"
	flags "github.com/jessevdk/go-flags"
)

type Options struct {
	// Room     string `short:"r" long:"room" description:"Room name" required:"true"`
	// Password string `short:"p" long:"pass" description:"Password"`
	Listen uint `short:"l" long:"listen" description:"Using port number" default:"8986"`
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
	sleigh()
}

func ignore(evt Event) bool {
	return strings.Contains(evt.RelPath, ".sleigh")
}

func sleigh() {
	showTextLogo()
	defer color.Unset()

	done := make(chan os.Signal, 1)
	signal.Notify(done, syscall.SIGINT)
	hostname, err := os.Hostname()
	if err != nil {
		color.Yellow("Cannot get hostname.")
		os.Exit(1)
	}
	color.Green("Version\t\t\t%s\n", "alpha")
	color.Green("Hostname\t\t%s\n", hostname)

	wd, err := os.Getwd()
	if err != nil {
		color.Yellow("Cannot get current working directory.")
		os.Exit(1)
	}
	items, err := GetItems(wd, wd, ignore)
	if err != nil {
		color.Yellow("Cannot get items in current working directory.")
		os.Exit(1)
	}

	hey := Hey{
		Hostname: hostname,
		Items:    items,
	}

	conn := NewConn(options.Listen)

	tracker := NewTracker(wd, ignore)
	defer tracker.Close()

	differ := NewDiffer(hostname, int(options.Listen), wd)
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
					differ.Notifications <- n
				} else if err := json.Unmarshal(d.Payload, &h); err == nil {
					color.Magenta("Hey! I'm %s", h.Hostname)
					for _, item := range h.Items {
						path := filepath.Join(wd, item.RelPath)

						info, err := os.Stat(path)
						if err != nil {
							differ.Notifications <- Notification{
								Hostname: h.Hostname,
								Event:    fsnotify.Create,
								Type:     File,
								Path:     item.RelPath,
								ModTime:  item.ModTime,
							}
							continue
						}
						checksum, err := GetChecksum(path)
						if err != nil {
							color.Yellow("ERROR in GetChecksum: %v\n", err)
							continue
						}
						modtime := info.ModTime().UnixNano()
						if checksum != item.Checksum {
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
								Event:    fsnotify.Create,
								Type:     File,
								Path:     local.RelPath,
								ModTime:  info.ModTime().UnixNano(),
							})
						}
					}
				}
			case evt := <-tracker.Events:
				n := Notification{
					Hostname: hostname,
					Event:    evt.Op,
					Type:     File,
					Path:     evt.RelPath,
					ModTime:  time.Now().UnixNano(),
				}
				if evt.Dir {
					n.Type = Dir
				}
				if evt.Op != fsnotify.Rename && evt.Op != fsnotify.Remove {
					info, err := os.Stat(evt.FullPath)
					if err != nil {
						color.Yellow("ERROR in tracker: %v\n", err)
						continue
					}
					n.ModTime = info.ModTime().UnixNano()
				}

				conn.Notify(n)
			case err := <-tracker.Errors:
				color.Yellow("ERROR in tracker: %v\n", err)
			case err := <-differ.Errors:
				color.Yellow("ERROR in patcher: %v\n", err)
			case err := <-conn.Errors:
				color.Yellow("ERROR in connector: %v\n", err)
			}
		}
	}()

	conn.Listen()
	differ.Start()
	tracker.Start()
	conn.Hey(hey)

	<-done
	color.Red("Bye!")
	defer color.Unset()
}
