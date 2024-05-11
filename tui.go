package main

import (
	"fmt"

	// "os"
	// "path/filepath"
	// "github.com/kardianos/osext"
	// "unicode"

	// "strings"
	// "time"
	// "log"

	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	// "github.com/fsnotify/fsnotify"
)

var skipNextViewportUpdate bool

type model struct {
	viewport viewport.Model
	ready    bool

	// choices  []string           // items on the to-do list
	list list.Model
	// choice   string
	quitting bool
	// cursor   int                // which to-do list item our cursor is pointing at
	// selected map[int]struct{}   // which to-do items are selected

}

func initialModel() model {
	items := []list.Item{
		item("fmt code"),
		item("git add ."),
		// item("git commit"), TODO
		item("git push"),
		// item("add to gitignore"), TODO
		// item("go mod init github_repo"), TODO
		// item("remove file from git history"), TODO
	}
	l := list.New(items, itemDelegate{}, defaultWidth, listHeight)
	// l.Title = "What do you want for dinner?"
	l.SetShowStatusBar(false)
	l.SetShowTitle(false)
	// l.SetShowStatusBar(true)
	l.SetFilteringEnabled(false)
	l.Styles.Title = titleStyle
	l.Styles.PaginationStyle = paginationStyle
	l.Styles.HelpStyle = helpStyle
	return model{
		// Our to-do list is a grocery list
		// choices:  []string{"Buy carrots", "Buy celery", "Buy kohlrabi"},

		// A map which indicates which choices are selected. We're using
		// the  map like a mathematical set. The keys refer to the indexes
		// of the `choices` slice, above.
		// selected: make(map[int]struct{}),
		list: l,
	}
}

type dataMsg struct {
	// reflex string
}

func (m *model) Init() tea.Cmd {
	// Just return `nil`, which means "no I/O right now, please."
	// return m.reflex("")
	return nil
}

type continueExec string

var fileUpdatedBool bool
var choiceActionBool bool

func (m *model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var (
		cmd tea.Cmd
		// cmds []tea.Cmd
		// continueExec string

	)

	if fileUpdatedBool {
		m.viewport.SetContent(watchRun())
		m.viewport.GotoBottom()
		fileUpdatedBool = false
	}
	if choiceActionBool {
		i, ok := m.list.SelectedItem().(item)
		if ok {
			m.viewport.SetContent(choiceAction(string(i)))
			m.viewport.GotoBottom()
		}
		fileUpdatedBool = false
		choiceActionBool = false
	}

	switch msg := msg.(type) {

	// Is it a key press?
	case tea.KeyMsg:

		// Cool, what was the actual key pressed?
		switch msg.String() {

		case "ctrl+c", "q":
			return m, tea.Quit

			// 	        // The "up" and "k" keys move the cursor up
			// 	        case "up", "k":
			// 	            if m.cursor > 0 {
			// 	                m.cursor--
			// 	            }
			//
			// 	        // The "down" and "j" keys move the cursor down
			// 	        case "down", "j":
			// 	            if m.cursor < len(m.choices)-1 {
			// 	                m.cursor++
			// 	            }

		// The "enter" key and the spacebar (a literal space) toggle
		// the selected state for the item that the cursor is pointing at.
		case "enter", " ":
			skipNextViewportUpdate = true
			m.viewport.SetContent("[Running...]") //TODO not working, probably need do in another update, so need to mkae it set here and return msg
			// continueExecVar = "choiceAction"
			// TUI.Send(func() tea.Msg{return continueExecVar})
			// continueExecVar = ""
			go TUI.Send(func() tea.Msg { // foroce rerender probably there are better solutions
				// time.Sleep(time.Second / 10)
				return ""
			})
			choiceActionBool = true
			return m, nil

		}
	case tea.WindowSizeMsg:
		headerHeight := lipgloss.Height(m.headerView())
		footerHeight := lipgloss.Height(m.footerView())
		verticalMarginHeight := headerHeight + footerHeight
		m.list.SetWidth(msg.Width)
		m.list.SetHeight(msg.Height / 2)

		if !m.ready {
			// Since this program is using the full size of the viewport we
			// need to wait until we've received the window dimensions before
			// we can initialize the viewport. The initial dimensions come in
			// quickly, though asynchronously, which is why we wait for them
			// here.
			m.viewport = viewport.New(msg.Width, msg.Height/2-verticalMarginHeight)
			m.viewport.YPosition = headerHeight
			m.viewport.HighPerformanceRendering = useHighPerformanceRenderer
			m.viewport.SetContent("[READY]")
			m.ready = true

			// This is only necessary for high performance rendering, which in
			// most cases you won't need.
			//
			// Render the viewport one line below the header.
			m.viewport.YPosition = headerHeight + 1
		} else {
			m.viewport.Width = msg.Width
			m.viewport.Height = msg.Height/2 - verticalMarginHeight
		}

		// if useHighPerformanceRenderer {
		// 	// Render (or re-render) the whole viewport. Necessary both to
		// 	// initialize the viewport and when the window is resized.
		// 	//
		// 	// This is needed for high-performance rendering only.
		// 	cmds = append(cmds, viewport.Sync(m.viewport))
		// }

	case fileUpdated:
		// if !skipNextViewportUpdate { //TODO useless
		m.viewport.SetContent("[Running...]")

		go TUI.Send(func() tea.Msg { //probably there are better solutions
			// time.Sleep(time.Second / 10)
			return ""
		})
		fileUpdatedBool = true
		// time.Sleep(time.Second*2)

		// continueExecVar = ""
		// return m, nil
		// }
		// skipNextViewportUpdate = false //TODO useless

		// case continueExec:
		// 	switch msg{
		// 		case "choiceAction":
		// 			i, ok := m.list.SelectedItem().(item)
		// 			if ok {
		// 				m.viewport.SetContent(choiceAction(string(i)))
		// 				m.viewport.GotoBottom()
		// 			}
		// 			// return m, nil
		// 		case "watchRun":
		// 			// log.Println("tets")
		// 			m.viewport.SetContent(watchRun())
		// 			m.viewport.GotoBottom()
		// 			// return m, nil
		//
		// 	}

	}

	// Return the updated model to the Bubble Tea runtime for processing.
	// Note that we're not returning a command.
	// var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
	// return m, nil
}

func (m *model) View() string {

	if !m.ready {
		return "\n  Initializing..."
	}

	// return fmt.Sprintf("%s\n%s\n%s\n%s", m.headerView(), m.viewport.View(), m.footerView(), m.choicesView())
	// TODO if window too small, hide list.View
	return fmt.Sprintf("%s\n%s\n%s\n%s", m.headerView(), m.viewport.View(), m.footerView(), m.list.View())

}
