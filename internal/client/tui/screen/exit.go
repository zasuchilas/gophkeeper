package screen

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/zasuchilas/gophkeeper/internal/client/tui/component"
)

var _ State = (*LogoPassItemScreen)(nil)

type ExitScreen struct{}

func NewExitScreen() ExitScreen {
	return ExitScreen{}
}

func (s ExitScreen) Update(msg tea.Msg) (State, tea.Cmd) {
	return s, nil
}

func (s ExitScreen) View() string {
	scr := component.NewScreenView()
	scr.SetHeader()

	body := "Good luck!"
	scr.SetBody(body)

	return scr.String()
}
