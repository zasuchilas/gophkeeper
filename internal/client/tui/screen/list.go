package screen

import (
	"fmt"
	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/zasuchilas/gophkeeper/internal/client/grpcclient"
	"github.com/zasuchilas/gophkeeper/internal/client/model"
	"github.com/zasuchilas/gophkeeper/internal/client/tui/component"
	"github.com/zasuchilas/gophkeeper/internal/client/tui/style"
	"strconv"
	"strings"
	"time"
)

var _ State = (*ListScreen)(nil)

const (
	KeyListUpdate = KeyCtrl + "u"
)

type listMsg struct {
	items []model.ListSecretItem
}

type ListScreen struct {
	loading bool
	spinner spinner.Model
	err     error

	data  []string
	count int
	table table.Model
}

func NewListScreen() ListScreen {
	spin := component.Spinner()

	return ListScreen{
		loading: true,
		spinner: spin,

		//table:   makeTable(),
	}
}

func (s ListScreen) Update(msg tea.Msg) (State, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {

	case listMsg:
		s.loading = false
		s.err = nil
		s.count = len(msg.items)
		s.table = makeTable(msg.items)
		return s, nil
	case errMsg:
		s.err = msg
		s.loading = false
		return s, nil

	case tea.KeyMsg:
		switch msg.String() {
		case KeyCreateScreen:
			return NewCreateLogoPassScreen(), nil
		case KeyListUpdate:
			next := NewListScreen()
			return next, tea.Batch(
				next.spinner.Tick,
				next.list,
			)

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

	scr.SetScreenHeader("YOUR SECRETS",
		fmt.Sprintf("You have %d of them.", s.count))

	var b strings.Builder
	if s.count != 0 {
		//	b.WriteString("there are no secrets\n")
		//} else {
		b.WriteString(style.TableBase.Render(s.table.View()))
	}
	scr.SetBody(b.String())

	cmd := []component.PressItem{
		{KeyListUpdate, "update list"},
		{KeyCreateScreen, "create"},
		{"enter", "secret item"},
	}
	scr.SetFooter(cmd)
	return scr.String()
}

func (s *ListScreen) list() tea.Msg {

	items, err := grpcclient.ApiService.GetSecretList()
	if err != nil {
		return errMsg{
			err: err,
		}
	}

	return listMsg{items: items}
}

func makeTable(items []model.ListSecretItem) table.Model {
	columns := []table.Column{
		{Title: "ID", Width: 4},
		{Title: "Name", Width: 25},
		{Title: "Size", Width: 4},
		{Title: "Type", Width: 10},
		{Title: "Updated", Width: 15},
	}

	rows := make([]table.Row, len(items))
	for i := range items {
		id := strconv.FormatInt(items[i].ID, 10)
		name := items[i].Name
		size := strconv.FormatInt(items[i].Size, 10)
		secretType := items[i].SecretType
		updatedAt := items[i].UpdatedAt.Format(time.DateOnly)

		rows[i] = table.Row{
			id,
			name,
			size,
			secretType,
			updatedAt,
		}
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
