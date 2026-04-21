package main

import (
	"fmt"
	"os"

	tea 	 "github.com/charmbracelet/bubbletea"
	lipgloss "github.com/charmbracelet/lipgloss"
)

// ──────────────────────────────────────────────────────────────────
//  COLOR PALETTE
// ──────────────────────────────────────────────────────────────────

var (
	colorPrimary   = lipgloss.Color("#A7C7E7") // pastel blue
	colorHighlight = lipgloss.Color("#F5E6A8") // pastel yellow — active item
	colorText      = lipgloss.Color("#D0D0D0") // body text
	colorDim       = lipgloss.Color("#707070") // inactive items, footer
)

// ──────────────────────────────────────────────────────────────────
//  STYLES
// ──────────────────────────────────────────────────────────────────

var (
	logoStyle = lipgloss.NewStyle().
			Foreground(colorPrimary).
			Bold(true)

	bioStyle = lipgloss.NewStyle().
			Foreground(colorText).
			Width(58)

	menuItemStyle = lipgloss.NewStyle().
			Foreground(colorDim).
			Padding(0, 2)

	menuItemActiveStyle = lipgloss.NewStyle().
				Foreground(colorHighlight).
				Bold(true).
				Padding(0, 2)

	artStyle = lipgloss.NewStyle().
			Foreground(colorPrimary).
			Padding(1, 4, 0, 4)

	rightColumnStyle = lipgloss.NewStyle().
				PaddingLeft(2).
				PaddingTop(1)

	footerStyle = lipgloss.NewStyle().
			Foreground(colorDim).
			Padding(1, 2)
)

// ──────────────────────────────────────────────────────────────────
//  ASCII ART
// ──────────────────────────────────────────────────────────────────

// logo. 
const nameLogo = `
 _ __ ___   __ _  __ _ _ __ __ _ _   _(_)
| '_ ' _ \ / _' |/ _' | '__/ _' | | | | |
| | | | | | (_| | (_| | | | (_| | |_| | |
|_| |_| |_|\__,_|\__,_|_|  \__, |\__,_|_|
                              |_|        
`

// ascii art (placeholder)
// TODO: replace ascii art
const asciiArtLogo = `
░░░░▒▒▒▒▓▓▓▓████████▓▓▓▓▒▒▒▒░░░░
░░▒▒▒▓▓▓▓██████████████▓▓▓▒▒▒░░
░▒▒▓▓▓████████████████████▓▓▓▒▒░
▒▓▓████████████████████████████▓▒
▓██████████████████████████████▓
████████████████████████████████
████████████████████████████████
████████████████████████████████
████████████████████████████████
████████████████████████████████
▓██████████████████████████████▓
▒▓▓████████████████████████████▓▒
░▒▒▓▓▓████████████████████▓▓▓▒▒░
░░▒▒▒▓▓▓▓██████████████▓▓▓▒▒▒░░
░░░░▒▒▒▒▓▓▓▓████████▓▓▓▓▒▒▒▒░░░░
`

// ──────────────────────────────────────────────────────────────────
//  CONTENT
// ──────────────────────────────────────────────────────────────────

const bio = `is a Computer Engineering student at
Universidad San Jorge (Spain), currently
on exchange at Hochschule Reutlingen (Germany).

Focused on software development, data science,
and machine learning. Hands-on experience with
Python, SQL, Java, and C.

Built projects exploring personality data,
substance consumption patterns, and action
recognition for live events.

Explore the sections below →`

// ──────────────────────────────────────────────────────────────────
//  MODEL
// ──────────────────────────────────────────────────────────────────

// section rerpresents a single option in the menu
type section struct {
	title 		string
}

// model holds state of the TUI
type model struct{
	sections  []section		// array of items that compose the screen
	cursor 	  int			// index of the currently highlighted item
}

//initialModel returns the starting state of the app
func initialModel() model {
	// fill the model with the initial elements
	return model{
		sections: []section{
			{title: "About"}, 
			{title: "Projects"},
			{title: "Skills"},
			{title: "Contact"},
			{title: "CV"},
		}, 
		cursor: 0,
	}
}

// model.Init function is called when the program starts, 
// it can return a command to run 
func (m model) Init() tea.Cmd{
	return nil
}

// model.Update function is called every time something happens, 
// it receives a message and returns the new state of the model (and a command to run)
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd){
	switch msg := msg.(type){
	case tea.KeyMsg: 
		switch msg.String(){
		// quit condition
		case "ctrl+c", "q": 
			return m, tea.Quit
		// move cursor left (keys) - without exceding the avaliable elements
		case "left", "h": 
			if m.cursor > 0 {
				m.cursor--
			}
		// move cursor right (keys) - without exceding the avaliable elements
		case "right", "l": 
			if m.cursor < len(m.sections)-1 {
				m.cursor++
			}
		// enter into section
		case "enter": 
			// TODO: declare enter action
		}  
	}
	return m, nil
}

// model.View function is called every time the state changes, 
// it returns the new state of the model
func (m model) View() string{
	// LEFT COLUMN contains the ascii logo
	leftCol := artStyle.Render(asciiArtLogo)

	// RIGHT COLUMN contains name + bio + horizontal menu
	logo := logoStyle.Render(nameLogo)
	bioText := bioStyle.Render(bio)

	// build the horizontal menu
	var menuItems []string
	for i, s := range m.sections{
		if i == m.cursor{
			menuItems = append(menuItems, 
			menuItemActiveStyle.Render("✦ "+s.title))
		}else {
			menuItems = append(menuItems, 
			menuItemStyle.Render("  "+s.title))
		}
	}
	// align the menu items horizontally
	menu := lipgloss.JoinHorizontal(lipgloss.Top, menuItems...)

	rightCol := rightColumnStyle.Render(
		lipgloss.JoinVertical(
			lipgloss.Left,
			logo,
			"",
			bioText, 
			"", 
			"", 
			menu,
		),
	)

	// align and combine the columns
	mainContent := lipgloss.JoinHorizontal(lipgloss.Top, leftCol, rightCol)
	// footer with keybinds
	footer := footerStyle.Render(
		"[← → to select · enter to open · q to quit]",
	)

	// return the complete layout
	return lipgloss.JoinVertical(lipgloss.Left, mainContent, footer)
}

func main(){
	// bubbletea program with the model as the initial state
	p := tea.NewProgram(initialModel(), tea.WithAltScreen())

	// run, if there is an error, print it and exit with err status 1
	if _, err := p.Run(); err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
}
