package tui

import (
	//"fmt"	
	"strings"

	"github.com/charmbracelet/lipgloss"

	"github.com/maarqui/portfolio-tui/internal/content"
)

// ──────────────────────────────────────────────────────────────────
//  ASCII ART
// ──────────────────────────────────────────────────────────────────

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

// ──────────────────────────────────────────────────────────────────
//  TOP-LEVEL VIEW
// ──────────────────────────────────────────────────────────────────

// View is Bubble Tea's entry point to the render layer.
// It dispatches to the right screen based on m.view.
func (m Model) View() string {
	switch m.view {
	case viewDetail:
		return m.detailView()
	default:
		return m.menuView()
	}
}

// ──────────────────────────────────────────────────────────────────
//  MENU VIEW
// ──────────────────────────────────────────────────────────────────

func (m Model) menuView() string {
	// Left column: portrait art.
	leftCol := artStyle.Render(portraitArt)

	// Right column: logo + bio + horizontal menu.
	logo := logoStyle.Render(nameLogo)
	bioText := bioStyle.Render(content.Bio)

	// Build menu items.
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
			bioText,
			"",
			"",
			menu,
		),
	)

	mainContent := lipgloss.JoinHorizontal(lipgloss.Top, leftCol, rightCol)
	footer := footerStyle.Render("[← → to select · enter to open · q to quit]")

	return lipgloss.JoinVertical(lipgloss.Left, mainContent, footer)
}

// ──────────────────────────────────────────────────────────────────
//  DETAIL VIEW
// ──────────────────────────────────────────────────────────────────

func (m Model) detailView() string {
	var body string
	switch m.currentalias() {
	case "about":
		body = sectionBodyStyle.Render(content.AboutText)
	case "projects":
		body = projectsBlockStyle.Render(m.renderProjects())
	case "skills":
		body = m.renderSkills()
	case "contact":
		body = m.renderContact()
	case "cv":
		body = m.renderCV()
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

	return lipgloss.JoinVertical(lipgloss.Left, title, body, "", footer)
}

// renderProjects builds the projects list with the active one expanded.
func (m Model) renderProjects() string {
	var b strings.Builder

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

	return b.String()
}

// ──────────────────────────────────────────────────────────────────
//  SKILLS VIEW
// ──────────────────────────────────────────────────────────────────

// renderSkills builds the full skills block grouped by category.
func (m Model) renderSkills() string {
	var b strings.Builder

	for _, cat := range content.Skills {
		b.WriteString(skillCategoryStyle.Render(cat.Name) + "\n")
		// Join items with a middle dot — a classic chip-list separator.
		items := strings.Join(cat.Items, " · ")
		b.WriteString(skillItemsStyle.Render(items) + "\n")
	}

	return skillsBlockStyle.Render(b.String())
}

// ──────────────────────────────────────────────────────────────────
//  CONTACT VIEW
// ──────────────────────────────────────────────────────────────────

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

// ──────────────────────────────────────────────────────────────────
//  CV VIEW
// ──────────────────────────────────────────────────────────────────

// renderCV builds a compact CV view: intro + structured blocks + hint.
func (m Model) renderCV() string {
	var b strings.Builder

	// Intro paragraph.
	b.WriteString(sectionBodyStyle.Render(content.CVIntro) + "\n")

	// Structured blocks.
	for _, block := range content.CVBlocks {
		b.WriteString(cvHeadingStyle.Render(block.Heading) + "\n")
		for _, line := range block.Lines {
			b.WriteString(cvLineStyle.Render(line) + "\n")
		}
	}

	// Download hint at the bottom.
	b.WriteString(cvHintStyle.Render(content.CVDownloadHint))

	return cvBlockStyle.Render(b.String())
}