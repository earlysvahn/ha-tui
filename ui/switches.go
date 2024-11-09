package ui

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/earlysvahn/ha-tui/pkg/api"
)

type SwitchesModel struct {
	client        *api.Client
	switches      []string
	selectedIndex int
}

func NewSwitchesModel(client *api.Client) SwitchesModel {
	switches, _ := client.FetchSwitches() // Fetch switches from API
	return SwitchesModel{
		client:   client,
		switches: switches,
	}
}

func (m SwitchesModel) Init() tea.Cmd {
	return nil
}

func (m SwitchesModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
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
			if m.selectedIndex < len(m.switches)-1 {
				m.selectedIndex++
			}
		case "enter":
			// Toggle the selected switch
			selectedSwitch := m.switches[m.selectedIndex]
			// Assuming the switch can be either "on" or "off"
			if err := m.client.TurnOnSwitch(selectedSwitch); err == nil {
				// Turn on successful, now toggle to off for next selection
				m.client.TurnOffSwitch(selectedSwitch)
			}
		}
	}
	return m, nil
}

func (m SwitchesModel) View() string {
	s := "Switches\n"
	for i, sw := range m.switches {
		cursor := " "
		if i == m.selectedIndex {
			cursor = ">"
		}
		s += fmt.Sprintf("%s %s\n", cursor, sw)
	}
	s += "\nPress enter to toggle switch, q to quit."
	return s
}
