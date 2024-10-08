/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"crucigrama/tui"
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "crucigrama",
	Short: "Un crucigrama en la terminal",
	Long:  `Selecciona un tema y resuelve un crucigrama en la terminal.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Parse flags
		topicVal, _ := cmd.Flags().GetString("tema")
		// wordCount, _ := cmd.Flags().GetInt("palabras")
		// size, _ := cmd.Flags().GetInt("lado")

		model := tui.RootScreen()
		if topicVal != "" {
			model = tui.RootScreenWithModel(tui.OptionsScreen(topicVal))
		}
		p := tea.NewProgram(model, tea.WithAltScreen())
		_, err := p.Run()
		if err != nil {
			fmt.Println("Error al empezar el programa:", err)
			os.Exit(1)
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().StringP("tema", "t", "", "El tema del crucigrama")
	rootCmd.Flags().IntP("palabras", "p", 10, "Cantidad de palabras a poner en el crucigrama")
	rootCmd.Flags().IntP("lado", "l", 20, "Tamaño del lado del crucigrama")
}
