package screen

import (
	"fmt"
	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/zasuchilas/gophkeeper/internal/client/tui/component"
	"github.com/zasuchilas/gophkeeper/internal/client/tui/style"
	"strings"
)

var _ State = (*ListScreen)(nil)

const (
	KeyListUpdate = KeyCtrl + "u"
)

type ListScreen struct {
	loading bool
	data    []string

	table table.Model
}

func NewListScreen() ListScreen {
	return ListScreen{
		loading: false,
		table:   makeTable(),
	}
}

func (s ListScreen) Update(msg tea.Msg) (State, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case KeyItemScreen:
			return NewLogoPassItemScreen(), nil
		case KeyCreateScreen:
			return NewCreateLogoPassScreen(), nil
		case KeyListUpdate:
			return NewListScreen(), nil

		case "enter":
			return NewLogoPassItemScreen(), nil
			//return s, tea.Batch(
			//	tea.Printf("Let's go to %s!", s.table.SelectedRow()[1]),
			//)
		}
	}

	s.table, cmd = s.table.Update(msg)
	return s, cmd
}

func (s ListScreen) View() string {
	scr := component.NewScreenView()
	scr.SetAppHeader()
	scr.SetScreenHeader("YOUR SECRETS", fmt.Sprintf("You have %d of them.", 15))

	var b strings.Builder
	b.WriteString(style.TableBase.Render(s.table.View()))
	scr.SetBody(b.String())

	cmd := []component.PressItem{
		{KeyListUpdate, "update list"},
		{KeyCreateScreen, "create"},
		{KeyItemScreen, "secret item"},
	}
	scr.SetFooter(cmd)
	return scr.String()
}

func makeTable() table.Model {
	columns := []table.Column{
		{Title: "ID", Width: 4},
		{Title: "Name", Width: 25},
		{Title: "Size", Width: 4},
		{Title: "Type", Width: 10},
		{Title: "Updated", Width: 15},
	}

	rows := []table.Row{
		{"1", "Tokyo", "10", "LOGO_PASS", "2025-02-02"},
		{"2", "Tokyo", "10", "LOGO_PASS", "2025-02-02"},
		{"3", "Tokyo", "10", "LOGO_PASS", "2025-02-02"},
		{"4", "Tokyo", "10", "LOGO_PASS", "2025-02-02"},
		{"5", "Tokyo", "10", "LOGO_PASS", "2025-02-02"},
		{"6", "Tokyo", "10", "LOGO_PASS", "2025-02-02"},
		{"7", "Tokyo", "10", "LOGO_PASS", "2025-02-02"},
		{"8", "Tokyo", "10", "LOGO_PASS", "2025-02-02"},
		{"9", "Tokyo", "10", "LOGO_PASS", "2025-02-02"},
	}

	t := table.New(
		table.WithColumns(columns),
		table.WithRows(rows),
		table.WithFocused(true),
		table.WithHeight(7),
	)

	s := table.DefaultStyles()
	s.Header = s.Header.
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("240")).
		BorderBottom(true).
		Bold(false)
	s.Selected = s.Selected.
		Foreground(lipgloss.Color("229")).
		Background(lipgloss.Color("57")).
		Bold(false)
	t.SetStyles(s)

	return t
}
