/*
Copyright © 2024 Edgar Avila Agramonte
*/
package main

import (
	"crucigrama/tui"
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	var m tea.Model = tui.TopicScreen()
	model := tui.RootScreenWithModel(m)
	p := tea.NewProgram(model, tea.WithAltScreen())
	_, err := p.Run()
	if err != nil {
		fmt.Println("Error al empezar el programa:", err)
		os.Exit(1)
	}
}
