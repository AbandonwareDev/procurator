package main

// https://github.com/charmbracelet/bubbletea/blob/master/examples/list-simple/main.go
import (
	"fmt"
	"io"
	"os/exec"
	"strings"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type item string

func (i item) FilterValue() string { return "" }

type itemDelegate struct{}

const listHeight = 14
const defaultWidth = 20

func (d itemDelegate) Height() int                             { return 1 }
func (d itemDelegate) Spacing() int                            { return 0 }
func (d itemDelegate) Update(_ tea.Msg, _ *list.Model) tea.Cmd { return nil }
func (d itemDelegate) Render(w io.Writer, m list.Model, index int, listItem list.Item) {
	i, ok := listItem.(item)
	if !ok {
		return
	}

	str := fmt.Sprintf("%d. %s", index+1, i)

	fn := itemStyle.Render
	if index == m.Index() {
		fn = func(s ...string) string {
			return selectedItemStyle.Render("> " + strings.Join(s, " "))
		}
	}

	fmt.Fprint(w, fn(str))
}

var (
	// titleStyle        = lipgloss.NewStyle().MarginLeft(2)
	itemStyle         = lipgloss.NewStyle().PaddingLeft(4)
	selectedItemStyle = lipgloss.NewStyle().PaddingLeft(2).Foreground(lipgloss.Color("170"))
	paginationStyle   = list.DefaultStyles().PaginationStyle.PaddingLeft(4)
	helpStyle         = list.DefaultStyles().HelpStyle.PaddingLeft(4).PaddingBottom(1)
	quitTextStyle     = lipgloss.NewStyle().Margin(1, 0, 2, 4)
)

// func (m *model) choicesView() string {
//     // The header
//     s := "What should we buy at the market?\n\n"
//
//     // Iterate over our choices
//     for i, choice := range m.choices {
//
//         // Is the cursor pointing at this choice?
//         cursor := " " // no cursor
//         if m.cursor == i {
//             cursor = ">" // cursor!
//         }
//
//         // Is this choice selected?
//         checked := " " // not selected
//         if _, ok := m.selected[i]; ok {
//             checked = "x" // selected!
//         }
//
//         // Render the row
//         s += fmt.Sprintf("%s [%s] %s\n", cursor, checked, choice)
//     }
//
//     // The footer
//     s += "\nPress q to quit.\n"
//
//     // Send the UI for rendering
//     // return s
//     return s
// }
//

// type errorRunningCommand string
// const errorMsg=""

func choiceAction(input string) string {
	switch input {
	case "fmt code":
		// TODO if go
		output, err := exec.Command("go", "fmt").CombinedOutput()
		if err != nil {
			// var errorRnCmd errorRunningCommand
			// errorRnCmd = errorMsg+string(output)
			return "[ERROR] - executing command failed\n" + string(output)
			// tui.Send(errorRnCmd)
		}
		return "[OK]\n" + string(output)
	case "git add .":
		output, err := exec.Command("git", "add", ".").CombinedOutput()
		if err != nil {
			return "[ERROR] - executing command failed\n" + string(output)

		}
		return "[OK]\n" + string(output)
	case "git commit":
		// TODO add input message box
	case "git push":
		output, err := exec.Command("git", "push").CombinedOutput()
		if err != nil {
			return "[ERROR] - executing command failed\n" + string(output)
		}
		return "[OK]\n" + string(output)
	case "add file to gitignore":
		// TODO add input message box with option to choose file or import filter
	case "go mod init github_repo":
		// TODO add input message box OR add fetching link via .git
	case "remove file from git history":
		// TODO add file chooser
	default:
		return "[ERROR] - unexisting option"
	}
	return "[ERROR] - unexisting option"
}
