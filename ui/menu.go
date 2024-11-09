package ui

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/earlysvahn/ha-tui/pkg/api"
)

var (
	headerStyle   = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("5"))
	selectedStyle = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("2"))
	normalStyle   = lipgloss.NewStyle().Foreground(lipgloss.Color("7"))
)

type MenuModel struct {
	statusMessage string
	menuOptions   []string
	selectedIndex int
	client        *api.Client
}

func NewMenuModel(statusMessage string, client *api.Client) MenuModel {
	return MenuModel{
		statusMessage: statusMessage,
		menuOptions:   []string{"Lights", "Switches", "Scenes"},
		client:        client,
	}
}

func (m MenuModel) Init() tea.Cmd {
	return nil
}

func (m MenuModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
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
			if m.selectedIndex < len(m.menuOptions)-1 {
				m.selectedIndex++
			}
		case "enter":
			switch m.menuOptions[m.selectedIndex] {
			case "Lights":
				return NewLightsModel(m.client), nil
			case "Switches":
				return NewSwitchesModel(m.client), nil
			case "Scenes":
				return NewScenesModel(m.client), nil
			}
		}
	}
	return m, nil
}

func (m MenuModel) View() string {
	s := headerStyle.Render("Welcome to HA-TUI!") + "\n"
	s += normalStyle.Render(m.statusMessage) + "\n\n"
	s += normalStyle.Render("Use the up/down arrow keys to navigate and enter to select an option.") + "\n\n"

	icons := []string{"💡", "🔌", "🎭"}
	for i, option := range m.menuOptions {
		cursor := " "
		style := normalStyle
		if i == m.selectedIndex {
			cursor = ">"
			style = selectedStyle
		}
		s += fmt.Sprintf("%s %s %s\n", cursor, icons[i], style.Render(option))
	}

	s += normalStyle.Render("\nPress q to quit.\n")
	return s
}
