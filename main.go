package main

import (
	"fmt"
	"os"
	"sync"

	tea "github.com/charmbracelet/bubbletea"
)

var waitGroup sync.WaitGroup
var TUI *tea.Program

func main() {
	_, err := ParseOptions()
	// options, err := ParseOptions()
	errHandler(err, "Error parsing options")

	model := initialModel()
	TUI = tea.NewProgram(
		&model,
		tea.WithAltScreen(), // use the full size of the terminal in its "alternate screen buffer"
		// tea.WithMouseCellMotion(), // turn on mouse support so we can track the mouse wheel
	)

	go watchFilesystem()

	if _, err := TUI.Run(); err != nil {
		errHandler(err, "Tui error:")
	}
}

func errHandler(err error, message string) {
	if err != nil {
		fmt.Printf("%s: %s\n", message, err)
		os.Exit(1)
	}
}
