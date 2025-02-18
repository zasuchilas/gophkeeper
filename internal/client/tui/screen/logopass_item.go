package screen

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/zasuchilas/gophkeeper/internal/client/secret"
	"github.com/zasuchilas/gophkeeper/internal/client/tui/component"
)

var _ State = (*LogoPassItemScreen)(nil)

type LogoPassItemScreen struct {
	loading    bool
	data       []byte
	err        error
	secretData *secret.LogoPass // secret.Secret // secret.LogoPass
}

func NewLogoPassItemScreen() LogoPassItemScreen {

	data := []byte("eyJsb2dpbiI6ItCQ0L3QvdCwIiwicGFzc3dvcmQiOiIxMjMiLCJpbmZvIjoicXdlIiwibWV0YSI6IjAwNyJ9")
	secretData := secret.NewEmptyLogoPass()
	err := secretData.DecryptFromBase64(data)

	return LogoPassItemScreen{
		loading:    false,
		data:       data,
		err:        err,
		secretData: secretData,
	}
}

func (s LogoPassItemScreen) Update(msg tea.Msg) (State, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case KeyListScreen:
			return NewListScreen(), nil
		case "enter":
			//s.err = s.secretData.DecryptFromBase64(s.data)
			return s, nil
		}
	}
	return s, nil
}

func (s LogoPassItemScreen) View() string {
	scr := component.NewScreenView()
	scr.SetHeader()

	body := fmt.Sprintf("SECRET ITEM SCREEN\nLOGIN: %s PASSWORD: %s",
		s.secretData.Login, s.secretData.Password)
	scr.SetBody(body)

	cmd := []component.PressItem{
		{KeyListScreen, "secrets"},
	}
	scr.SetFooter(cmd)
	return scr.String()
}
