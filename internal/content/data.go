// package content holds all the static portfolio data.
package content

// main menu bio.
const Bio = `is a Computer Engineering student at
Universidad San Jorge (Spain), currently
on exchange at Hochschule Reutlingen (Germany).

Focused on data science, software development,
and machine learning. Hands-on experience with
Python, SQL, Java, and C.

Built projects exploring personality data,
substance consumption patterns, and action
recognition for live events.

Explore the sections below →`

// about section.
const AboutText = `I'm Daniel Marquino Pérez, a third-year Computer Engineering student at Universidad San Jorge in Zaragoza, Spain.

This year (2025–2026) I'm on an Erasmus exchange at Hochschule Reutlingen, Germany, where I'm studying at the Fakultät Informatik.

My main interests are software development and data science. I enjoy the full pipeline of a data project: from exploring a raw dataset to deploying a trained model. But I also love backend work, distributed systems, and anything that involves making software talk to other software.

Outside of code, I speak Spanish (native), English (C1), German (A1, learning), and French (A2). I previously worked as a network engineer deploying fibre-optic infrastructure, which gave me a solid understanding of how the internet actually works at the cable level.

This portfolio itself is a Go project using Bubble Tea and Wish — feel free to check the source on GitHub.`

// Project represents a single portfolio project entry based on the following structure:
type Project struct {
	Title       string
	Stack       string
	Description string
	Link        string
}

// Projects is the list of highlighted portfolio projects.
var Projects = []Project{
	{
		Title: "Drug Consumption Analysis",
		Stack: "Python · pandas · scikit-learn · Jupyter",
		Description: "End-to-end analysis of a 1,885-record dataset linking " +
			"personality traits and demographics to substance use. Three " +
			"phases: exploratory data analysis with demographic and " +
			"personality correlations; binary classification benchmarking " +
			"Logistic Regression, Random Forest, and SVM; and ordinal " +
			"regression to predict consumption levels.",
		Link: "github.com/maarqui/drug-consumption-analysis",
	},
	{
		Title: "Intelligent Streaming System with Action Recognition",
		Stack: "Python · Computer Vision · Machine Learning",
		Description: "Contributing to an action recognition streaming " +
			"system powered by computer vision models, designed for " +
			"intelligent monitoring of live music events with direct " +
			"commercial application.",
		Link: "-- not publicly available --",
	},
	{
		Title: "portfolio-tui",
		Stack: "Go · Bubble Tea · Lip Gloss · Wish",
		Description: "This very portfolio you are reading, exposed as a " +
			"public SSH service. Built as a learning project to pick up " +
			"Go, TUI design, and real-world deployment.",
		Link: "github.com/maarqui/portfolio-tui",
	},
}