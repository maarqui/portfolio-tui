package tui

import (
	//"fmt"	
	"strings"

	"github.com/charmbracelet/lipgloss"

	"github.com/maarqui/portfolio-tui/internal/content"
)

//  ASCII ART

const nameLogo = `
 _ __ ___   __ _  __ _ _ __ __ _ _   _(_)
| '_ ' _ \ / _' |/ _' | '__/ _' | | | | |
| | | | | | (_| | (_| | | | (_| | |_| | |
|_| |_| |_|\__,_|\__,_|_|  \__, |\__,_|_|
                              |_|        
`

const portraitArt = `
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

//  TOP-LEVEL VIEW

// View is Bubble Tea's entry point to the render layer.
// it checks if the terminal is wide enough to render, if not shows message.
func (m Model) View() string {
	// if the terminal dimensions are not known (width == 0), return an empty string
	// Bubble Tea will call View again after receiving the first WindowSizeMsg 
	if m.width == 0 {
		return ""
	}

	// if the terminal is too narrow, show a message instead of trying to render
	if m.width < minWidth {
		msg := tooNarrowStyle.Render(
			"Please resize your terminal\n" +
				"to at least 80 columns.",
		)
		return placeFullScreen(msg, m.width, m.height)
	}

	// normal dispatch by view.
	switch m.view {
	case viewDetail:
		return m.detailView()
	default:
		if m.isNarrow(){
			return m.menuViewNarrow()
		}
		return m.menuView()
	}
}

// isNarrow returns true when the terminal is narrow enough that the
// two-column layout should collapse to a vertical stack.
func (m Model) isNarrow() bool {
	return m.width < wideWidth
}

//  MENU VIEW

func (m Model) menuView() string {
	// effective content width: capped at contentWidth, but never wider than the terminal
	w := fitContentWidth(m.width)
	// left column: portrait art.
	leftCol := artStyle.Render(portraitArt)

	// right column width: what remains after the left column.
	leftColW := lipgloss.Width(leftCol)
	rightColW := w - leftColW
	if rightColW < 40 {
		rightColW = 40 // safety floor
	}

	// right column: logo + bio + horizontal menu.
	logo := logoStyle.Render(nameLogo)
	// bio width is adaptable to variable space.
	bio := bioStyle.Width(rightColW - 4).Render(content.Bio)

	// build menu items.
	items := make([]string, 0, len(m.sections))
	for i, s := range m.sections {
		if i == m.cursor {
			items = append(items, menuItemActiveStyle.Render("✦ "+s.title))
		} else {
			items = append(items, menuItemStyle.Render("  "+s.title))
		}
	}
	menu := lipgloss.JoinHorizontal(lipgloss.Top, items...)

	rightCol := rightColumnStyle.Render(
		lipgloss.JoinVertical(
			lipgloss.Left,
			logo,
			"",
			bio,
			"",
			"",
			menu,
		),
	)

	mainContent := lipgloss.JoinHorizontal(lipgloss.Top, leftCol, rightCol)
	footer := footerStyle.Render("[← → ↑ ↓ to select · enter to open · q to quit]")

	// assemble the view, then center horizontally
	full := lipgloss.JoinVertical(lipgloss.Left, mainContent, footer)
	return centerHorizontally(full, m.width)
}

// menuViewNarrow renders a stacked, single-column version of the main
// menu for terminals between minWidth and wideWidth.
func (m Model) menuViewNarrow() string {
	w := m.width

	// logo first.
	logo := logoStyle.Render(nameLogo)

	// bio at full available width.
	bio := bioStyle.Width(w - 4).Render(content.Bio)

	// menu items stacked vertically.
	items := make([]string, 0, len(m.sections))
	for i, s := range m.sections {
		if i == m.cursor {
			items = append(items, menuItemActiveStyle.Render("✦ "+s.title))
		} else {
			items = append(items, menuItemStyle.Render("  "+s.title))
		}
	}
	menu := lipgloss.JoinVertical(lipgloss.Left, items...)

	content := lipgloss.JoinVertical(
		lipgloss.Left,
		logo,
		"",
		bio,
		"",
		menu,
	)

	footer := footerStyle.Render("[↑ ↓ to select · enter · q quit]")
	full := lipgloss.JoinVertical(lipgloss.Left, content, footer)

	return centerHorizontally(full, m.width)
}

//  DETAIL VIEW

func (m Model) detailView() string {
	w := fitContentWidth(m.width)

	var body string
	switch m.currentalias() {
	case "about":
		// adapts its width dinamically 
		body = sectionBodyStyle.Width(w - 4).Render(content.AboutText)
	case "projects":
		body = m.renderProjects(w)
	case "skills":
		body = m.renderSkills(w)
	case "contact":
		body = m.renderContact()
	case "cv":
		body = m.renderCV(w)
	default:
		body = sectionBodyStyle.Render("Unknown section.")
	}

	title := sectionTitleStyle.Render("▸ " + m.sections[m.cursor].title)

	var footerText string
	if m.currentalias() == "projects" {
		footerText = "[↑ ↓ browse projects · esc back · q quit]"
	} else {
		footerText = "[esc back · q quit]"
	}
	footer := footerStyle.Render(footerText)

	// assamble & allign view
	full := lipgloss.JoinVertical(lipgloss.Left, title, body, "", footer)
	return centerHorizontally(full, m.width)
}



// PROJECTS VIEW

// renderProjects builds the projects list with the active one expanded.
func (m Model) renderProjects(contentW int) string {
	var b strings.Builder

	// description width adapts to available content width.
	descW := contentW - 10
	if descW < 40 {
		descW = 40
	}

	for i, p := range content.Projects {
		marker := "  "
		if i == m.projectCursor {
			marker = "▸ "
		}

		titleLine := projectTitleStyle.Render(marker + p.Title)
		b.WriteString(titleLine + "\n")

		// only show stack + description for the currently selected project.
		if i == m.projectCursor {
			b.WriteString("    " + projectStackStyle.Render(p.Stack) + "\n")
			b.WriteString(projectDescStyle.Render(p.Description) + "\n")
			if p.Link != "" {
				b.WriteString("    " + projectLinkStyle.Render(p.Link) + "\n")
			}
		}
		b.WriteString("\n")
	}

	return projectsBlockStyle.Render(b.String())
}

//  SKILLS VIEW

// renderSkills builds the full skills block grouped by category.
func (m Model) renderSkills(contentW int) string {
	var b strings.Builder

	itemsW := contentW - 6
	if itemsW < 40 {
		itemsW = 40
	}

	for _, cat := range content.Skills {
		b.WriteString(skillCategoryStyle.Render(cat.Name) + "\n")
		// join items with a middle dot
		items := strings.Join(cat.Items, " · ")
		b.WriteString(skillItemsStyle.Width(itemsW).Render(items) + "\n")
	}

	return skillsBlockStyle.Render(b.String())
}

//  CONTACT VIEW

// renderContact builds the contact list with icons, labels and values.
func (m Model) renderContact() string {
	var b strings.Builder

	for _, c := range content.Contacts {
		icon := contactIconStyle.Render(c.Icon)
		label := contactLabelStyle.Render(c.Label)
		value := contactValueStyle.Render(c.Value)

		// JoinHorizontal keeps icon / label / value on the same line
		// with consistent widths (the label has a fixed Width(10)).
		line := lipgloss.JoinHorizontal(lipgloss.Top,
			icon, " ", label, value,
		)
		b.WriteString(line + "\n")
	}

	return contactBlockStyle.Render(b.String())
}

//  CV VIEW

// renderCV builds a compact CV view: intro + structured blocks + hint.
func (m Model) renderCV(contentW int) string {
	var b strings.Builder

	introW := contentW - 4
	if introW < 40 {
		introW = 40
	}

	// intro paragraph.
	b.WriteString(sectionBodyStyle.Width(introW).Render(content.CVIntro) + "\n")

	// structured blocks.
	for _, block := range content.CVBlocks {
		b.WriteString(cvHeadingStyle.Render(block.Heading) + "\n")
		for _, line := range block.Lines {
			b.WriteString(cvLineStyle.Render(line) + "\n")
		}
	}

	// download hint at the bottom.
	b.WriteString(cvHintStyle.Render(content.CVDownloadHint))

	return cvBlockStyle.Render(b.String())
}