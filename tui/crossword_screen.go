package tui

import (
	"crucigrama/core"
	"crucigrama/wikipedia"
	"strings"

	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// *****************************************************************************
// Styles
// *****************************************************************************
var crosswordStyle = lipgloss.
	NewStyle().
	Foreground(colorWhite).
	Border(lipgloss.RoundedBorder()).
	BorderForeground(colorPrimary).
	PaddingLeft(1).
	PaddingRight(1).
	MarginTop(1).
	MarginBottom(1).
	MarginLeft(2)

var crosswordWordsStyle = lipgloss.NewStyle().
	Foreground(colorGray).
	Border(lipgloss.RoundedBorder()).
	BorderForeground(colorWhite).
	MarginTop(2).
	PaddingLeft(1).
	PaddingRight(1)

// *****************************************************************************
// Model
// *****************************************************************************
type crosswordScreenModel struct {
	loading   bool
	spinner   spinner.Model
	crossword [][]string
	words     []string
	size      int
	wordCount int
	title     string
	error     error
}

func CrosswordScreen(title string, size int, wordCount int) crosswordScreenModel {
	s := spinner.New()
	s.Spinner = spinner.Dot
	s.Style = spinnerStyle

	return crosswordScreenModel{
		spinner:   s,
		loading:   true,
		title:     title,
		size:      size,
		wordCount: wordCount,
	}
}

// *****************************************************************************
// Event handlers
// *****************************************************************************
func (m crosswordScreenModel) onExtractFetched(extract string) (tea.Model, tea.Cmd) {
	return m, getWords(extract, m.size, m.wordCount)
}

func (m crosswordScreenModel) onWordsFetched(words []string) (tea.Model, tea.Cmd) {
	m.words = words
	return m, getCrossword(m.words, m.size)
}

func (m crosswordScreenModel) onCrosswordFetched(crossword [][]string) (tea.Model, tea.Cmd) {
	m.loading = false
	m.crossword = crossword
	return m, nil
}

// *****************************************************************************
// Lifecycle methods
// *****************************************************************************
func (m crosswordScreenModel) Init() tea.Cmd {
	return tea.Batch(m.spinner.Tick, getExtract(m.title))
}

func (m crosswordScreenModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "esc", "q":
			return m, tea.Quit
		case "enter":
			return m, nil
		}
	case spinner.TickMsg:
		if m.loading {
			m.spinner, cmd = m.spinner.Update(msg)
		}
		return m, cmd
	case extractMsg:
		return m.onExtractFetched(string(msg))
	case wordsMsg:
		return m.onWordsFetched(msg)
	case crosswordMsg:
		return m.onCrosswordFetched(msg)
	case error:
		m.loading = false
		m.error = msg
		return m, nil
	}

	return m, cmd
}

func (m crosswordScreenModel) View() string {
	view := ""
	if m.loading {
		view = spinnerStyle.Render(m.spinner.View() + " Generando crucigrama...")
	} else if m.error != nil {
		view = lipgloss.JoinVertical(
			lipgloss.Top,
			errorStyle.Render(m.error.Error()),
			helpStyle.Render("(Presiona Esc para salir)"),
		)
	} else {
		lines := make([]string, len(m.crossword))
		for i, row := range m.crossword {
			lines[i] = strings.Join(row, " ")
		}
		crosswordStr := strings.Join(lines, "\n")
		wordsStr := strings.Join(m.words, "\n")
		titleRendered := titleStyle.PaddingLeft(1).PaddingRight(1).Render(m.title)
		helpRendered := helpStyle.Render("(Presiona Esc para salir)")
		crosswordRendered := crosswordStyle.Render(crosswordStr)
		wordsRendered := crosswordWordsStyle.Height(m.size).Render(wordsStr)
		// Join with crossword
		renderedCrossword := lipgloss.JoinVertical(
			lipgloss.Top,
			lipgloss.PlaceHorizontal(m.size*2+8, lipgloss.Center, titleRendered),
			crosswordRendered,
			lipgloss.PlaceHorizontal(m.size*2+8, lipgloss.Center, helpRendered),
		)

		// Join with
		view = lipgloss.JoinHorizontal(
			lipgloss.Top,
			wordsRendered,
			renderedCrossword,
		)
	}

	return containerStyle.Render(view)
}

// *****************************************************************************
// Messages
// *****************************************************************************
type extractMsg string

func getExtract(title string) tea.Cmd {
	return func() tea.Msg {
		extract, err := wikipedia.QueryExtract(title)
		if err != nil {
			return err
		}
		return extractMsg(extract)
	}
}

type wordsMsg []string

func getWords(extract string, size int, wordCount int) tea.Cmd {
	return func() tea.Msg {
		words, err := core.MostImportantWords(extract, wordCount, size)
		if err != nil {
			return err
		}
		return wordsMsg(words)
	}
}

type crosswordMsg [][]string

func getCrossword(words []string, size int) tea.Cmd {
	return func() tea.Msg {
		crossword, err := core.Crossword(words, size)
		if err != nil {
			return err
		}
		return crosswordMsg(crossword)
	}
}
