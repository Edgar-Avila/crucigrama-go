package tui

import (
	tea "github.com/charmbracelet/bubbletea"
)

type RootScreenModel struct {
	model tea.Model
	Err   error
}

func RootScreen() RootScreenModel {
	var rootModel tea.Model

	rootModel = TopicScreen()

	return RootScreenModel{
		model: rootModel,
	}
}

func RootScreenWithModel(model tea.Model) RootScreenModel {
	return RootScreenModel{
		model: model,
	}
}


func (m RootScreenModel) Init() tea.Cmd {
	return m.model.Init()
}

func (m RootScreenModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		containerStyle = containerStyle.Width(msg.Width - 2).Height(msg.Height - 2)
	case error:
		m.Err = msg
		return m.model, tea.Quit
	}
	return m.model.Update(msg)
}

func (m RootScreenModel) View() string {
	return m.model.View()
}

func (m RootScreenModel) SwitchScreen(model tea.Model) (tea.Model, tea.Cmd) {
	m.model = model
	return m.model, m.model.Init()
}
