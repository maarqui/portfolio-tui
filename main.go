package main 

import( 
	"fmt"
	"os"
	"strings"
	// imprort bubble tea modules from repository
	tea "github.com/charmbracelet/bubbletea"
)

// menuItem rerpesents a single option in the menu
type menuItem struct {
	title 		string
	description string
}

// model holds state of the TUI
type model struct{
	menuItems []menuItem	// array of items that compose the screen
	cursor 	  int			// index of the currently highlighted item
}

//initialModel returns the starting state of the app
func initialModel() model {
	// fill the model with the initial elements
	return model{
		menuItems: []menuItem{
			{title: "About me", description: "Who I am and what I do"}, 
			{title: "Projects", description: "Things I've built"},
			{title: "Skills", description: "Tech I work with"},
			{title: "Contact", description: "How to reach me"},
			{title: "CV", description: "Get my resume as a PDF"},
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
		// move cursor up (keys) - without exceding the avaliable elements
		case "up", "k": 
			if m.cursor > 0 {
				m.cursor--
			}
		// move cursor down (keys) - without exceding the avaliable elements
		case "down", "j": 
			if m.cursor < len(m.menuItems)-1 {
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
	var b strings.Builder

	//header of the TUI screen
	b.WriteString("\n")
	b.WriteString("  maarqui-tui — personal portfolio\n")
	b.WriteString("  ─────────────────────────────────\n\n")

	// menu items
	for i, item := range m.menuItems{
		cursor := "  "
		if m.cursor == i{
			cursor = "▸ "
		}
		b.WriteString(fmt.Sprintf("  %s%s\n", cursor, item.title))
	}

	// description of the highlighted item
	b.WriteString("\n  ─────────────────────────────────\n")
	b.WriteString(fmt.Sprintf("  %s\n", m.menuItems[m.cursor].description))

	// footer of the TUI screen
	b.WriteString("\n  ↑/↓ navigate • enter select • q quit\n")

	return b.String()
}

func main(){
	// bubbletea program with the model as the initial state
	p := tea.NewProgram(initialModel())

	// run, if there is an error, print it and exit with err status 1
	if _, err := p.Run(); err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
}
