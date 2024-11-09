package main

import (
	"log"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/earlysvahn/ha-tui/pkg/api"
	"github.com/earlysvahn/ha-tui/ui"
)

func main() {
	baseURL := os.Getenv("HA_LOCAL_URL")
	token := os.Getenv("HA_CLITOKEN")

	client := api.NewClient(baseURL, token)
	status, err := client.GetStatus()
	if err != nil {
		log.Fatalf("Error connecting to Home Assistant: %v", err)
	}

	// Initialize main menu model with status
	mainMenu := ui.NewMenuModel(status, client)

	p := tea.NewProgram(mainMenu)
	if _, err := p.Run(); err != nil {
		log.Fatalf("Error running TUI: %v", err)
	}
}
