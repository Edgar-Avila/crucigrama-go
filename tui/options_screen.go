package tui

import (
	"crucigrama/wikipedia"
	"errors"

	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
)

// *****************************************************************************
// Model
// *****************************************************************************
type optionsScreenModel struct {
	loading   bool
	spinner   spinner.Model
	list      list.Model
	topic     string
}

func OptionsScreen(topic string) optionsScreenModel {
	s := spinner.New()
	s.Spinner = spinner.Dot
	s.Style = spinnerStyle

	return optionsScreenModel{
		spinner:   s,
		loading:   true,
		list:      list.New(nil, list.NewDefaultDelegate(), 0, 0),
		topic:     topic,
	}
}

// *****************************************************************************
// Event handlers
// *****************************************************************************
func (m optionsScreenModel) onOptionsFetched(items []wikipedia.OpenSearchOption) (tea.Model, tea.Cmd) {
	m.loading = false
	listItems := make([]list.Item, 0, len(items))
	for _, i := range items {
		listItems = append(listItems, listItem(i))
	}
	m.list = list.New(listItems, list.NewDefaultDelegate(), 0, 0)
	m.list.Title = "Artículos relacionados"
	h, v := containerStyle.GetFrameSize()
	m.list.SetSize(containerStyle.GetWidth()-h, containerStyle.GetHeight()-v)
	return m, nil
}

// *****************************************************************************
// Lifecycle methods
// *****************************************************************************
func (m optionsScreenModel) Init() tea.Cmd {
	if m.topic == "" {
		return func() tea.Msg {
			return errors.New("no se ha especificado un tema")
		}
	}
	return tea.Batch(m.spinner.Tick, fetchOptions(m.topic))
}

func (m optionsScreenModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit
		case "esc":
			if !m.list.SettingFilter() && !m.list.IsFiltered() {
				return m, tea.Quit
			}
		case "enter":
			item, ok := m.list.SelectedItem().(listItem)
			if !ok {
				return m, tea.Quit
			}
			if !m.list.SettingFilter() {
				title := item.Title()
				return RootScreen().SwitchScreen(sizeScreen(title))
			}
		}
	case spinner.TickMsg:
		if m.loading {
			m.spinner, cmd = m.spinner.Update(msg)
		}
		return m, cmd
	case fetchOptionsMsg:
		return m.onOptionsFetched(msg)
	}

	if !m.loading {
		m.list, cmd = m.list.Update(msg)
	}
	return m, cmd
}

func (m optionsScreenModel) View() string {
	view := ""
	if m.loading {
		view = spinnerStyle.Render(m.spinner.View() + " Obteniendo artículos relacionados...")
	} else {
		view = m.list.View()
	}

	return containerStyle.Render(view)
}

// *****************************************************************************
// Messages
// *****************************************************************************
type listItem wikipedia.OpenSearchOption

func (i listItem) Title() string       { return i.Text }
func (i listItem) Description() string { return i.Link }
func (i listItem) FilterValue() string { return i.Text }

type fetchOptionsMsg []wikipedia.OpenSearchOption

func fetchOptions(topic string) tea.Cmd {
	return func() tea.Msg {
		options, err := wikipedia.OpenSearch(topic)
		if err != nil {
			return err
		}
		return fetchOptionsMsg(options)
	}
}
