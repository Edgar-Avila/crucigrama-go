package tui

import (
	"strconv"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type sizeScreenModel struct {
	sizeInput      textinput.Model
	wordCountInput textinput.Model
	size           int
	wordCount      int
	title          string
	validationMsg string
}

func sizeScreen(title string) sizeScreenModel {
	sizeInput := textinput.New()
	sizeInput.Placeholder = "20"
	sizeInput.Focus()
	wordCountInput := textinput.New()
	wordCountInput.Placeholder = "10"
	wordCountInput.Blur()

	return sizeScreenModel{
		sizeInput:      sizeInput,
		wordCountInput: wordCountInput,
		size:           0,
		wordCount:      0,
		title:          title,
		validationMsg:  "",
	}
}

func (m sizeScreenModel) Init() tea.Cmd {
	return nil
}

func (m sizeScreenModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		containerStyle = containerStyle.Width(msg.Width - 2).Height(msg.Height - 2)
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "esc":
			return m, tea.Quit
		case "enter":
			sizeStr := m.sizeInput.Value()
			wordCountStr := m.wordCountInput.Value()
			size, err := strconv.Atoi(sizeStr)
			m.validationMsg = ""
			if err != nil {
				m.sizeInput.SetValue("")
				m.validationMsg += "El tamaño debe ser un número.\n"
			}
			if size < 10 {
				m.sizeInput.SetValue("")
				m.validationMsg += "El tamaño debe ser al menos 10.\n"
			}
			wordCount, err := strconv.Atoi(wordCountStr)
			if err != nil {
				m.wordCountInput.SetValue("")
				m.validationMsg += "El número de palabras debe ser un número.\n"
			}
			if wordCount < 2 {
				m.wordCountInput.SetValue("")
				m.validationMsg += "El número de palabras debe ser al menos 2.\n"
			}
			if m.validationMsg != "" {
				return m, nil
			}
			return RootScreen().SwitchScreen(CrosswordScreen(m.title, size, wordCount))
		case "up", "down", "tab":
			if m.sizeInput.Focused() {
				m.sizeInput.Blur()
				m.wordCountInput.Focus()
			} else {
				m.wordCountInput.Blur()
				m.sizeInput.Focus()
			}
		}
	}

	m.sizeInput, cmd = m.sizeInput.Update(msg)
	m.wordCountInput, cmd = m.wordCountInput.Update(msg)
	return m, cmd
}

func (m sizeScreenModel) View() string {
	return containerStyle.Render(
		lipgloss.JoinVertical(
			lipgloss.Top,
			titleStyle.Render("Ingresa el tamaño del crucigrama"),
			"",
			m.sizeInput.View(),
			"",
			titleStyle.Render("Ingresa el número de palabras"),
			"",
			m.wordCountInput.View(),
			"",
			helpStyle.Render("(Presiona Esc para salir)"),
			helpStyle.Render("(Presiona Arriba/Abajo/Tab para cambiar de campo)"),
			helpStyle.Render("(Presiona Enter para continuar)"),
			errorStyle.Render(m.validationMsg),
		),
	)
}
