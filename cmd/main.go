package main

import (
	"fmt"
	"log"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/earlysvahn/ha-tui/pkg/api"
)

// Define the TUI model
type model struct {
	statusMessage string
}

// Initialize the TUI model
func initialModel() model {
	return model{statusMessage: "Connecting to Home Assistant..."}
}

// Define the Init function for Bubble Tea, which starts any commands
func (m model) Init() tea.Cmd {
	return nil
}

// Update function to handle events and interactions
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		// Exit the app on "q" or "ctrl+c"
		if msg.String() == "q" || msg.String() == "ctrl+c" {
			return m, tea.Quit
		}
	}
	return m, nil
}

// View function renders the TUI
func (m model) View() string {
	return fmt.Sprintf(
		"%s\nPress q to quit.\n",
		m.statusMessage,
	)
}

func main() {
	baseURL := "http://localhost:8123" // Home Assistant URL
	token := os.Getenv("HA_TOKEN")     // Get token from environment variable

	client := api.NewClient(baseURL, token)
	status, err := client.GetStatus()
	if err != nil {
		log.Fatalf("Error connecting to Home Assistant: %v", err)
	}

	// Display connection status in the TUI model
	initialModel := model{statusMessage: status}

	// Start the TUI
	p := tea.NewProgram(initialModel)
	if err := p.Start(); err != nil {
		log.Fatalf("Error running TUI: %v", err)
	}
}
