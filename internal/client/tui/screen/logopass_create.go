package screen

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/zasuchilas/gophkeeper/internal/client/secret"
	"github.com/zasuchilas/gophkeeper/internal/client/tui/component"
)

var _ State = (*CreateLogoPassScreen)(nil)

type CreateLogoPassScreen struct {
	formData secret.Secret // secret.LogoPass
	loading  bool
	data     []byte
	err      error
}

func NewCreateLogoPassScreen() CreateLogoPassScreen {
	return CreateLogoPassScreen{
		loading: false,
	}
}

func (s CreateLogoPassScreen) Update(msg tea.Msg) (State, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case KeyListScreen:
			return NewListScreen(), nil
		}
	}
	return s, nil
}

func (s CreateLogoPassScreen) View() string {
	scr := component.NewScreenView()
	scr.SetHeader()

	body := fmt.Sprintf("CREATE SECRET SCREEN %s", "")
	scr.SetBody(body)

	cmd := []component.PressItem{
		{KeyListScreen, "secrets"},
	}
	scr.SetFooter(cmd)
	return scr.String()
}
