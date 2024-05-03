package main

import (
	"fmt"
	"os"
	"sync"

	tea "github.com/charmbracelet/bubbletea"
)

var waitGroup sync.WaitGroup



func main() {
	options, err := ParseOptions()
	errHandler(err, "Error parsing options")

	fmt.Println("hello world", options)

// 	// start workers in parallel
// 	for i := 0; i < options.Threads; i++ {
// 		waitGroup.Add(1)
// 		go func() {
// 			fmt.Println("do parallel stuff")
// 
// 			defer waitGroup.Done()
// 		}()
// 	}
// 	waitGroup.Wait()
p := tea.NewProgram(initialModel())
    if _, err := p.Run(); err != nil {
        fmt.Printf("Alas, there's been an error: %v", err)
        os.Exit(1)
    }



}

func errHandler(err error, message string) {
	if err != nil {
		fmt.Printf("%s: %s\n", message, err)
		os.Exit(1)
	}
}
