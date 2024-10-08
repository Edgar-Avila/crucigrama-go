package tui

import (
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type topicScreenModel struct {
	topicInput textinput.Model
}

func TopicScreen() topicScreenModel {
	ti := textinput.New()
	ti.Placeholder = "Golang"
	ti.Focus()
	return topicScreenModel{
		topicInput: ti,
	}
}

func (m topicScreenModel) Init() tea.Cmd {
	return nil
}

func (m topicScreenModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "esc":
			return m, tea.Quit
		case "enter":
			topic := m.topicInput.Value()
			return RootScreen().SwitchScreen(OptionsScreen(topic))
		}
	}

	m.topicInput, cmd = m.topicInput.Update(msg)
	return m, cmd
}

func (m topicScreenModel) View() string {
	return containerStyle.Render(
		lipgloss.JoinVertical(
			lipgloss.Top,
			titleStyle.Render("Ingresa un tema"),
			"",
			m.topicInput.View(),
			"",
			helpStyle.Render("(Presiona Esc para salir)"),
		),
	)
}
