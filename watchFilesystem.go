package main

// It's not an actual reflex (https://github.com/cespare/reflex). I just wrote some code as for reflex and couldn't manage to adapt it fast before moving to creating own solution

import (
	// "errors"
	// // "flag"
	"io/fs"
		"path/filepath"
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
func watchFilesystem() { //TODO watch subfolders
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

	//get subdirectories
	var subdirs []string 
	filepath.WalkDir(".", func(path string, file fs.DirEntry,  err error) error {
			if err != nil {
				return err
			}
			if file.IsDir() {
				// fmt.Println(path);
				subdirs = append(subdirs,path)	
			}
	
			return nil;
		});
	for _,dir := range subdirs {
		err = watcher.Add(dir)
			if err != nil {
				errHandler(err, "Filesystem Watcher - Can't add current dir:")
			}
	}
	// Add a current path.
	// err = watcher.Add("...") //TODO waiting for official recursive add https://github.com/fsnotify/fsnotify/issues/18
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
