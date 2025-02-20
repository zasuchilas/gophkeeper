package screen

import (
	"fmt"
	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/zasuchilas/gophkeeper/internal/client/grpcclient"
	"github.com/zasuchilas/gophkeeper/internal/client/tui/component"
	"github.com/zasuchilas/gophkeeper/internal/client/tui/style"
	"google.golang.org/grpc/codes"
	"strings"
)

var _ State = (*AuthScreen)(nil)

const (
	ActionAuthLogin    = "login"
	ActionAuthRegister = "register"

	KeyAuthSwitch = KeyCtrl + "r"
)

type AuthScreen struct {
	// login or register
	action string

	// form
	focusIndex int
	inputs     []textinput.Model

	// request
	loading bool
	spinner spinner.Model
	err     error
}

func NewAuthScreen() AuthScreen {
	return AuthScreen{
		action:  ActionAuthLogin,
		spinner: component.Spinner(),
		inputs:  makeTextInputs(),
	}
}

func (s AuthScreen) Update(msg tea.Msg) (State, tea.Cmd) {
	switch msg := msg.(type) {

	case statusMsg:
		next := NewListScreen()
		return next, tea.Batch(
			next.spinner.Tick,
			next.list,
		)
	case errMsg:
		s.err = msg
		s.loading = false
		return s, nil

	case tea.KeyMsg:
		k := msg.String()
		switch k {
		case KeyAuthSwitch:
			s.action = s.getNextAction()
			return s, nil

		// Set focus to next input
		case "tab", "shift+tab", "enter", "up", "down":

			// Did the user press enter while the submit button was focused?
			// If so, exit.
			if k == "enter" && s.focusIndex == len(s.inputs)-1 {
				s.loading = true

				action := s.login
				if s.action == ActionAuthRegister {
					action = s.register
				}

				return s, tea.Batch(
					s.spinner.Tick,
					action,
				)
			}

			// Cycle indexes
			if k == "up" || k == "shift+tab" {
				s.focusIndex--
			} else {
				s.focusIndex++
			}

			if s.focusIndex > len(s.inputs)-1 {
				s.focusIndex = 0
			} else if s.focusIndex < 0 {
				s.focusIndex = len(s.inputs) - 1
			}

			cmds := make([]tea.Cmd, len(s.inputs))
			for i := 0; i <= len(s.inputs)-1; i++ {
				if i == s.focusIndex {
					// Set focused state
					cmds[i] = s.inputs[i].Focus()
					s.inputs[i].PromptStyle = style.Focused
					s.inputs[i].TextStyle = style.Focused
					continue
				}
				// Remove focused state
				s.inputs[i].Blur()
				s.inputs[i].PromptStyle = style.NoStyle
				s.inputs[i].TextStyle = style.NoStyle
			}

			return s, tea.Batch(cmds...)
		}

	case spinner.TickMsg:
		var cmd tea.Cmd
		s.spinner, cmd = s.spinner.Update(msg)
		return s, cmd
	}

	// Handle character input and blinking
	cmd := s.updateInputs(msg)

	return s, cmd // textinput.Blink
}

func (s AuthScreen) View() string {
	scr := component.NewScreenView()
	scr.SetAppHeader()

	switch s.action {
	case ActionAuthLogin:
		scr.SetScreenHeader("LOGIN", "Log in to the app.")
	case ActionAuthRegister:
		scr.SetScreenHeader("REGISTER", "Register in the app.")
	}

	var b strings.Builder

	if s.loading {
		b.WriteString(fmt.Sprintf("loading ... %s\n", s.spinner.View()))
		scr.SetBody(b.String())
		scr.SetFooter(nil)
		return scr.String()
	}

	for i := range s.inputs {
		b.WriteString(s.inputs[i].View())
		if i < len(s.inputs)-1 {
			b.WriteRune('\n')
		}
	}

	// error message
	if s.err != nil {
		b.WriteRune('\n')
		b.WriteString(style.Error.Render(fmt.Sprintf("Something went wrong: %s", s.err.Error())))
	}

	scr.SetBody(b.String())

	nextAction := s.getNextAction()
	cmd := []component.PressItem{
		{KeyAuthSwitch, nextAction},
		//{KeyCtrl + nextAction[:1], nextAction},
	}
	scr.SetFooter(cmd)
	return scr.String()
}

func (s *AuthScreen) updateInputs(msg tea.Msg) tea.Cmd {
	cmds := make([]tea.Cmd, len(s.inputs))

	// Only text inputs with Focus() set will respond, so it's safe to simply
	// update all of them here without any further logic.
	for i := range s.inputs {
		s.inputs[i], cmds[i] = s.inputs[i].Update(msg)
	}

	return tea.Batch(cmds...)
}

func (s *AuthScreen) login() tea.Msg {

	login := s.inputs[0].Value()
	password := s.inputs[1].Value()

	// validate inputs
	if len(login) == 0 || len(password) == 0 {
		return errMsg{
			err: fmt.Errorf("login and password are required"),
		}
	}

	err := grpcclient.ApiService.Login(login, password)
	if err != nil {
		return errMsg{
			err: err,
		}
	}

	return statusMsg(codes.OK)
}

func (s *AuthScreen) register() tea.Msg {

	login := s.inputs[0].Value()
	password := s.inputs[1].Value()

	// validate inputs
	if len(login) == 0 || len(password) == 0 {
		return errMsg{
			err: fmt.Errorf("login and password are required"),
		}
	}

	err := grpcclient.ApiService.Register(login, password)
	if err != nil {
		return errMsg{
			err: err,
		}
	}

	return statusMsg(codes.OK)
}

func (s AuthScreen) getNextAction() string {
	if s.action == ActionAuthLogin {
		return ActionAuthRegister
	}
	return ActionAuthLogin
}

func makeTextInputs() []textinput.Model {
	inputs := make([]textinput.Model, 2)

	var t textinput.Model
	for i := range inputs {
		t = textinput.New()
		t.Cursor.Style = style.Focused
		t.CharLimit = 32

		switch i {
		case 0:
			t.Placeholder = "Login"
			t.Focus()
			t.PromptStyle = style.Focused
			t.TextStyle = style.Focused
			t.CharLimit = 64
		case 1:
			t.Placeholder = "Password"
			t.EchoMode = textinput.EchoPassword
			t.EchoCharacter = '•'
		}

		inputs[i] = t
	}

	return inputs
}
