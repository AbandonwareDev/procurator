package main

import (
	// "fmt"

	"os"
	// "path/filepath"
	// "github.com/kardianos/osext"
	// "unicode"

	"strings"
	// "time"

	// tea "github.com/charmbracelet/bubbletea"
	// "github.com/charmbracelet/bubbles/viewport"
	"github.com/charmbracelet/lipgloss"
	// "github.com/fsnotify/fsnotify"
)

const useHighPerformanceRenderer = false

var (
	titleStyle = func() lipgloss.Style {
		b := lipgloss.RoundedBorder()
		b.Right = "├"
		return lipgloss.NewStyle().BorderStyle(b).Padding(0, 1)
	}()

	infoStyle = func() lipgloss.Style {
		b := lipgloss.RoundedBorder()
		b.Left = "┤"
		return titleStyle.Copy().BorderStyle(b)
	}()
)

func (m *model) headerView() string {
	title := titleStyle.Render("Procurator - " + getCurrentFolder())

	line := strings.Repeat("─", max(0, m.viewport.Width-lipgloss.Width(title)))
	return lipgloss.JoinHorizontal(lipgloss.Center, title, line)
}

func (m *model) footerView() string {
	// info := infoStyle.Render(fmt.Sprintf("%3.f%%", m.viewport.ScrollPercent()*100))
	line := strings.Repeat("─", max(0, m.viewport.Width))
	return lipgloss.JoinHorizontal(lipgloss.Center, line)
}

func getCurrentFolder() string { //TODO bug, shows not current folder but folder of execuutable
	// folderPath, err := osext.ExecutableFolder()
	folderPath, err := os.Getwd()
	// folderPath, err := filepath.Abs(filepath.Dir(os.Args[0]))
	folders := strings.Split(folderPath, "/")
	if err != nil {
		errHandler(err, "Can't get current folder")
	}
	capitalizedCurFolder := folders[len(folders)-1]
	capitalizedCurFolder = strings.ToUpper(capitalizedCurFolder[:1]) + capitalizedCurFolder[1:]
	return capitalizedCurFolder
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

///////////
