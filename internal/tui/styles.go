package tui

import "github.com/charmbracelet/lipgloss"

// LAYOUT CONSTANTS
// breakpoints in terminal columns.
//
//   width < minWidth          → "terminal too narrow" message
//   minWidth <= w < wideWidth → collapsed vertical layout
//   width >= wideWidth        → full two-column layout

const (
	minWidth     = 80
	wideWidth    = 100
	contentWidth = 130 // caps how wide the content grows
)

// COLOR PALLETE
// pastel blue for highlights/art, 
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
	// -- main view styles --
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

	// -- about view styles --

	sectionTitleStyle = lipgloss.NewStyle().
				Foreground(colorHighlight).
				Bold(true).
				Padding(1, 0, 1, 2)

	sectionBodyStyle = lipgloss.NewStyle().
				Foreground(colorText).
				Padding(0, 2).
				Width(80)

	// -- projects view styles --

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

	// -- Skills view styles --

	skillsBlockStyle = lipgloss.NewStyle().
				MarginLeft(2)

	skillCategoryStyle = lipgloss.NewStyle().
				Foreground(colorHighlight).
				Bold(true).
				MarginTop(1)

	skillItemsStyle = lipgloss.NewStyle().
				Foreground(colorText).
				MarginLeft(2).
				Width(80)

	// -- Contact view styles --

	contactBlockStyle = lipgloss.NewStyle().
				MarginLeft(2).
				MarginTop(1)

	contactIconStyle = lipgloss.NewStyle().
				Foreground(colorPrimary).
				Bold(true)

	contactLabelStyle = lipgloss.NewStyle().
				Foreground(colorDim).
				Width(10)

	contactValueStyle = lipgloss.NewStyle().
				Foreground(colorText)

	// -- CV view styles --

	cvBlockStyle = lipgloss.NewStyle().
				MarginLeft(2)

	cvHeadingStyle = lipgloss.NewStyle().
				Foreground(colorHighlight).
				Bold(true).
				MarginTop(1)

	cvLineStyle = lipgloss.NewStyle().
				Foreground(colorText).
				MarginLeft(2)

	cvHintStyle = lipgloss.NewStyle().
				Foreground(colorDim).
				Italic(true).
				MarginTop(2).
				MarginLeft(2)
	
	// -- responsiveness styles --

	// case: terminal too narrow
	tooNarrowStyle = lipgloss.NewStyle().
				Foreground(colorHighlight).
				Bold(true).
				Align(lipgloss.Center)
)

//  LAYOUT HELPERS

// fitContentWidth returns the effective content width for the current terminal.
func fitContentWidth(termWidth int) int {
	if termWidth > contentWidth {
		return contentWidth
	}
	return termWidth
}

// centerHorizontally centers a block of content horizontally within
// the given terminal width, padding both sides with blank space.
func centerHorizontally(content string, termWidth int) string {
	contentW := lipgloss.Width(content)
	if contentW >= termWidth {
		return content
	}
	leftPad := (termWidth - contentW) / 2
	return lipgloss.NewStyle().PaddingLeft(leftPad).Render(content)
}

// placeFullScreen places content centered both horizontally and
// vertically within the given terminal dimensions.
func placeFullScreen(content string, termWidth, termHeight int) string {
	return lipgloss.Place(
		termWidth, termHeight,
		lipgloss.Center, lipgloss.Center,
		content,
	)
}