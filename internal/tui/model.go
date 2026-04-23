package tui

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/maarqui/portfolio-tui/internal/content"
)

// view identifies which screen the user is currently seeing.
type view int

const (
	viewMenu view = iota
	viewDetail
)

// section is one entry in the main menu.
type section struct {
	title string
	alias string // stable identifier used to pick what to render
}

// Model holds the entire TUI state.
// NOTE: exported (capital M) because main.go creates it and hands it to tea.NewProgram.
type Model struct {
	sections []section
	cursor   int
	view     view

	// width and height track the current terminal size. Updated by tea.WindowSizeMsg
	width  int
	height int

	// projectCursor tracks which project is highlighted inside the Projects detail view.
	projectCursor int
}

// InitialModel builds the starting state.
func InitialModel() Model {
	return Model{
		sections: []section{
			{title: "About", alias: "about"},
			{title: "Projects", alias: "projects"},
			{title: "Skills", alias: "skills"},
			{title: "Contact", alias: "contact"},
			{title: "CV", alias: "cv"},
		},
		cursor: 0,
		view:   viewMenu,
	}
}

// Init runs once on startup.
func (m Model) Init() tea.Cmd {
	return nil
}

// Update receives messages (keypresses, window resizes, ...) and returns
// the new state, it dispatches based on the current view.
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.WindowSizeMsg:
		// store the new terminal size for later use (responsive layout).
		m.width = msg.Width
		m.height = msg.Height
		return m, nil

	case tea.KeyMsg:
		// Global keys: work from any view.
		switch msg.String() {
		case keyQuit, keyCtrlC:
			return m, tea.Quit
		}

		// View-specific keys.
		switch m.view {
		case viewMenu:
			return m.updateMenu(msg)
		case viewDetail:
			return m.updateDetail(msg)
		}
	}

	return m, nil
}

// updateMenu handles keys while the main menu is visible.
func (m Model) updateMenu(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch msg.String() {
	case keyLeft, keyLeftV:
		if m.cursor > 0 {
			m.cursor--
		}
	case keyRight, keyRightV:
		if m.cursor < len(m.sections)-1 {
			m.cursor++
		}
	case keyEnter:
		m.view = viewDetail
		m.projectCursor = 0 // reset when entering the projects view
	}
	return m, nil
}

// updateDetail handles keys while a section detail is visible.
func (m Model) updateDetail(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch msg.String() {
	case keyEsc, keyBack:
		m.view = viewMenu
		return m, nil
	}

	// within the Projects view, up/down navigate between projects.
	if m.currentalias() == "projects" {
		switch msg.String() {
		case keyUp, keyUpV:
			if m.projectCursor > 0 {
				m.projectCursor--
			}
		case keyDown, keyDownV:
			if m.projectCursor < len(content.Projects)-1 {
				m.projectCursor++
			}
		}
	}

	return m, nil
}

// currentalias returns the alias of the section currently highlighted in
// the menu. convenient for switches in the view layer.
func (m Model) currentalias() string {
	return m.sections[m.cursor].alias
}
