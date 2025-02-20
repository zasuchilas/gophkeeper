package tui

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/zasuchilas/gophkeeper/internal/client/tui/screen"
	"os"
)

type model struct {
	state screen.State
}

func initialModel() model {
	return model{
		state: screen.NewAuthScreen(),
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "ctrl+q":
			m.state = screen.NewExitScreen()
			return m, tea.Quit
		}
	}

	switch modelState := m.state.(type) {
	default:
		nextState, nextCmd := modelState.Update(msg)
		m.state = nextState
		return m, nextCmd
	}
}

func (m model) View() string {
	switch modelState := m.state.(type) {
	default:
		return modelState.View()
	}
}

func Run() {
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("an error: %v", err)
		os.Exit(1)
	}
}
