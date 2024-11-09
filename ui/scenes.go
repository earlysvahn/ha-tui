package ui

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/earlysvahn/ha-tui/pkg/api"
)

type ScenesModel struct {
	client        *api.Client
	scenes        []string
	selectedIndex int
}

func NewScenesModel(client *api.Client) ScenesModel {
	scenes, _ := client.FetchScenes() // Fetch scenes from API
	return ScenesModel{
		client: client,
		scenes: scenes,
	}
}

func (m ScenesModel) Init() tea.Cmd {
	return nil
}

func (m ScenesModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
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
			if m.selectedIndex < len(m.scenes)-1 {
				m.selectedIndex++
			}
		case "enter":
			// Activate the selected scene
			selectedScene := m.scenes[m.selectedIndex]
			m.client.ActivateScene(selectedScene)
		}
	}
	return m, nil
}

func (m ScenesModel) View() string {
	s := "Scenes\n"
	for i, scene := range m.scenes {
		cursor := " "
		if i == m.selectedIndex {
			cursor = ">"
		}
		s += fmt.Sprintf("%s %s\n", cursor, scene)
	}
	s += "\nPress enter to activate scene, q to quit."
	return s
}
