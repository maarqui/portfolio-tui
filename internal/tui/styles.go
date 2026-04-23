package tui

import "github.com/charmbracelet/lipgloss"

// color palette — pastel blue for highlights/art, 
// pastel yellow for the active cursor,
// grey for secondary info.
var (
	colorPrimary   = lipgloss.Color("#A7C7E7") // pastel blue
	colorHighlight = lipgloss.Color("#F5E6A8") // pastel yellow
	colorText      = lipgloss.Color("#D0D0D0") // body text
	colorDim       = lipgloss.Color("#707070") // inactive, footer
)

// styles used across all views.
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

	// ── styles for the detail view ──

	sectionTitleStyle = lipgloss.NewStyle().
				Foreground(colorHighlight).
				Bold(true).
				Padding(1, 0, 1, 2)

	sectionBodyStyle = lipgloss.NewStyle().
				Foreground(colorText).
				Padding(0, 2).
				Width(80)

	projectTitleStyle = lipgloss.NewStyle().
				Foreground(colorPrimary).
				Bold(true)

	projectStackStyle = lipgloss.NewStyle().
				Foreground(colorDim).
				Italic(true)

	projectLinkStyle = lipgloss.NewStyle().
				Foreground(colorHighlight).
				Underline(true)

	projectDescStyle = lipgloss.NewStyle().
				Foreground(colorText).
				MarginLeft(4).
				Width(75)

	projectsBlockStyle = lipgloss.NewStyle().
    			MarginLeft(2)
)
