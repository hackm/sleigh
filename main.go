package main

import (
	"fmt"
	"log"
	"net/url"
	"os"
	"os/signal"
	"path/filepath"
	"strings"
	"syscall"

	"github.com/fsnotify/fsnotify"

	"github.com/fatih/color"
)

func main() {
	port := 8986
	done := make(chan os.Signal, 1)
	signal.Notify(done, syscall.SIGINT)

	hostname, err := os.Hostname()
	if err != nil {
		color.Yellow("Cannot get hostname.")
		os.Exit(1)
	}

	wd, err := os.Getwd()
	if err != nil {
		color.Yellow("Cannot get current working directory.")
		os.Exit(1)
	}

	localPathResolver := func(relpath string) string {
		return filepath.Join(wd, relpath)
	}

	remoteURLResolver := func(n Notification) string {
		v := url.Values{}
		v.Set("path", n.RelPath)
		return fmt.Sprintf("http://%s:%d/contents?%s", n.IP, port, v.Encode())
	}

	target := func(evt Event) bool {
		if evt.Op == fsnotify.Rename || evt.Op == fsnotify.Remove {
			return true
		}
		if (evt.Op == fsnotify.Create || evt.Op == fsnotify.Write) && evt.Dir == false {
			stat, err := os.Stat(evt.FullPath)
			if err == nil && stat.Size() > 0 {
				return true
			}
		}
		return false
	}

	server := NewServer(hostname, port, localPathResolver)
	patcher := NewPatcher(server.Notifications, remoteURLResolver, localPathResolver)
	tracker := NewTracker(wd, ignore)

	go func() {
		for {
			select {
			case evt := <-tracker.Events:
				if target(evt) == false {
					continue
				}
				log.Printf("%s %s\n", evt.Op.String(), evt.RelPath)
				for _, ip := range server.Others() {
					n, err := createNotification(evt, hostname)
					if err != nil {
						color.Yellow("Cannot create Notification: %s", err)
						continue
					}
					err = Notify(ip, port, n)
					if err != nil {
						color.Yellow("Cannot notify to %s:%d : %s", ip, port, err)
						continue
					}
				}
			case hey := <-server.Heys:
				log.Printf("%s(%s): Hey!\n", hey.Hostname, hey.IP)
			case err := <-server.Errors:
				log.Printf("%+v\n", err)
			case err := <-patcher.Errors:
				log.Printf("%+v\n", err)
			case err := <-tracker.Errors:
				log.Printf("%+v\n", err)
			}
		}
	}()

	patcher.Start()
	tracker.Start()
	server.Start()
	defer server.Close()

	<-done
	color.Red("Bye!")
	defer color.Unset()
}

func ignore(evt Event) bool {
	return strings.Contains(evt.RelPath, ".sleigh")
}
