package screen

import tea "github.com/charmbracelet/bubbletea"

type State interface {
	Update(msg tea.Msg) (State, tea.Cmd)
	View() string
}

const (
	KeyCtrl         = "ctrl+"
	KeyAuthScreen   = KeyCtrl + "a"
	KeyListScreen   = KeyCtrl + "s"
	KeyItemScreen   = KeyCtrl + "i"
	KeyCreateScreen = KeyCtrl + "c"
)

type statusMsg int

type errMsg struct {
	err error
}

func (e errMsg) Error() string {
	return e.err.Error()
}
