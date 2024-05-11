package main

// It's not an actual reflex (https://github.com/cespare/reflex). I just wrote some code as for reflex and couldn't manage to adapt it fast before moving to creating own solution

import (
	// "errors"
	// // "flag"
	// "fmt"
	// "io/ioutil"
	// "log"
	// "os"
	// "os/exec"
	// "path"
	// "reflect"
	// "regexp"
	// "strconv"
	// // "strings"
	// "syscall"
	// "time"
	// "bytes"
	"log"
	"os/exec"

	"github.com/fsnotify/fsnotify"
	// tea "github.com/charmbracelet/bubbletea"
)

type fileUpdated bool

// func watchFilesystem(tui *tea.Program) {
func watchFilesystem() {
	// Create new watcher.
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		errHandler(err, "Can't setup filesystem watcher")
	}
	defer watcher.Close()

	// Start listening for events.
	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				// log.Println("event:", event)
				if event.Has(fsnotify.Write) {
					// log.Println("modified file:", event.Name)
					var fileUpdated1 fileUpdated
					TUI.Send(fileUpdated1)
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				log.Println("error:", err)
			}
		}
	}()

	// Add a current path.
	err = watcher.Add(".")
	if err != nil {
		errHandler(err, "Filesystem Watcher - Can't add current dir:")
	}

	// Block main goroutine forever.
	<-make(chan struct{})
}

func watchRun() string {
	//TODO if edited .go file
	return GoVet()
}

func GoVet() string {
	out1, _ := exec.Command("go", "mod", "tidy").CombinedOutput()

	out2, _ := exec.Command("go", "vet").CombinedOutput()

	output := string(out1) + string(out2)
	if output == "" {
		output = "[BUILD OK]"
	}
	return output
}
