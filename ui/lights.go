package ui

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/earlysvahn/ha-tui/pkg/api"
)

type LightsModel struct {
	client         *api.Client
	lights         []string // Placeholder; replace with appropriate struct
	filteredLights []string
	selectedIndex  int
	input          string
}

func NewLightsModel(client *api.Client) LightsModel {
	lights, _ := client.FetchLights() // Fetch lights from API
	return LightsModel{
		client:         client,
		lights:         lights,
		filteredLights: lights,
	}
}

func (m LightsModel) Init() tea.Cmd {
	return nil
}

func (m LightsModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit
		case "up", "k":
			if m.selectedIndex > 0 {
				m.selectedIndex--
			}
		case "down", "j":
			if m.selectedIndex < len(m.filteredLights)-1 {
				m.selectedIndex++
			}
		case "enter":
			return NewBrightnessInputModel(m.client, m.filteredLights[m.selectedIndex]), nil
		}
	}
	return m, nil
}

func (m LightsModel) View() string {
	s := "Lights\n"
	for i, light := range m.filteredLights {
		cursor := " "
		if i == m.selectedIndex {
			cursor = ">"
		}
		s += fmt.Sprintf("%s %s\n", cursor, light)
	}
	return s
}
