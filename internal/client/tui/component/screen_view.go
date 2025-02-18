package component

import (
	"fmt"
	"github.com/zasuchilas/gophkeeper/internal/client/tui/style"
	"strings"
)

type ScreenView struct {
	header string
	body   string
	footer string
}

func NewScreenView() *ScreenView {
	return &ScreenView{}
}

func (scr *ScreenView) String() string {
	var b strings.Builder
	b.WriteString(scr.header)
	b.WriteString(scr.body)
	b.WriteString(scr.footer)
	return b.String()
}

func (scr *ScreenView) SetHeader() {
	var b strings.Builder
	b.WriteString(style.Header.Render("gophkeeper cli v1.0.0"))
	b.WriteRune('\n')
	scr.header = b.String()
}

func (scr *ScreenView) SetBody(data string) {
	var b strings.Builder
	b.WriteString(data)
	b.WriteRune('\n')
	scr.body = b.String()
}

type PressItem struct {
	CmdKey   string // +, -, enter, esc, q
	ToResult string // increment, decrement, quit
}

func (scr *ScreenView) SetFooter(commands []PressItem) {
	var b strings.Builder

	commands = append(commands, PressItem{
		CmdKey:   "ctrl+c or ctrl+q",
		ToResult: "quit",
	})

	cmd := make([]string, len(commands))
	for i := range commands {
		cmd[i] = fmt.Sprintf("%s to %s", strings.ToUpper(commands[i].CmdKey), commands[i].ToResult)
	}

	content := fmt.Sprintf(
		"Press %s.",
		strings.Join(cmd, ", "),
	)

	b.WriteRune('\n')
	b.WriteString(style.Help.Render(content))
	b.WriteRune('\n')
	scr.footer = b.String()
}
