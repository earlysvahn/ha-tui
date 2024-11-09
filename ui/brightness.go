package ui

import (
	"fmt"
	"strconv"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/earlysvahn/ha-tui/pkg/api"
)

type BrightnessModel struct {
	client   *api.Client
	lightID  string
	input    string
	finished bool
}

func NewBrightnessInputModel(client *api.Client, lightID string) BrightnessModel {
	return BrightnessModel{
		client:  client,
		lightID: lightID,
	}
}

func (m BrightnessModel) Init() tea.Cmd {
	return nil
}

func (m BrightnessModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "enter":
			brightness, _ := strconv.Atoi(m.input)
			if brightness == 0 {
				m.client.TurnOffLight(m.lightID)
			} else {
				m.client.SetLightBrightness(m.lightID, brightness)
			}
			m.finished = true
			return NewMenuModel("Action complete!", m.client), nil
		case "q", "ctrl+c":
			return m, tea.Quit
		default:
			m.input += msg.String()
		}
	}
	return m, nil
}

func (m BrightnessModel) View() string {
	if m.finished {
		return "Brightness set. Returning to menu..."
	}
	return fmt.Sprintf("Set brightness (0-100) for light %s: %s", m.lightID, m.input)
}
