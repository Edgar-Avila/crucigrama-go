package tui

import "github.com/charmbracelet/lipgloss"

var colorPrimary = lipgloss.Color("#5f5fd7")
var colorBlack = lipgloss.Color("#000000")
var colorWhite = lipgloss.Color("#ffffff")
var colorGray = lipgloss.Color("#666666")

var containerStyle = lipgloss.NewStyle().
	Padding(2).
	Border(lipgloss.DoubleBorder()).
	BorderForeground(colorPrimary)

var titleStyle = lipgloss.NewStyle().
	Background(colorPrimary).
	Foreground(colorWhite)

var helpStyle = lipgloss.NewStyle().
	Foreground(colorGray)

var spinnerStyle = lipgloss.NewStyle().
	Foreground(colorPrimary)

