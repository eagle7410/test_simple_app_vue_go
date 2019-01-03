package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
	"regexp"
	"time"

	"github.com/radovskyb/watcher"
)

func main() {
	watchDir := "./back"
	runWatch(&watchDir)
}

func runWatch(dir *string) {
	w := watcher.New()

	// SetMaxEvents to 1 to allow at most 1 event's to be received
	// on the Event channel per watching cycle.
	//
	// If SetMaxEvents is not set, the default is to send all events.
	w.SetMaxEvents(1)

	// Only files that match the regular expression during file listings
	// will be watched.
	// r := regexp.MustCompile(`\.go$`)
	r := regexp.MustCompile(`\.go$`)
	w.AddFilterHook(watcher.RegexFilterHook(r, false))
	pwd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	go func() {
		for {
			select {
			case event := <-w.Event:
				fmt.Println(event) // Print the event's info.
				fmt.Println("go run " + pwd + "/back/main.go")
				cmd := exec.Command("sh", "go run "+pwd+"/back/main.go")
				var out bytes.Buffer
				cmd.Stdout = &out
				err := cmd.Run()
				if err != nil {
					log.Fatal(err)
				}

				fmt.Printf("in all caps: %q\n", out.String())

			case err := <-w.Error:
				log.Fatalln(err)
			case <-w.Closed:
				return
			}
		}
	}()

	// Watch recursively for changes.
	if err := w.AddRecursive(*dir); err != nil {
		log.Fatalln(err)
	}

	// Print a list of all of the files and folders currently
	// being watched and their paths.
	for path, f := range w.WatchedFiles() {
		fmt.Printf("%s: %s\n", path, f.Name())
	}

	fmt.Println()

	// Start the watching process - it'll check for changes every 100ms.
	if err := w.Start(time.Millisecond * 100); err != nil {
		log.Fatalln(err)
	}
}
