package style

import "github.com/charmbracelet/lipgloss"

var (
	Header = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("#FAFAFA")).
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("#FAFAFA")).
		BorderTop(true).
		BorderBottom(true).
		Bold(false)
	Title = lipgloss.NewStyle().Bold(true)
	Help  = lipgloss.NewStyle().Foreground(lipgloss.Color("240"))

	NoStyle = lipgloss.NewStyle()
	Focused = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))
	Blurred = lipgloss.NewStyle().Foreground(lipgloss.Color("240"))

	TableBase = lipgloss.NewStyle().
			BorderStyle(lipgloss.NormalBorder()).
			BorderForeground(lipgloss.Color("240"))

	Error = lipgloss.NewStyle().Foreground(lipgloss.Color("9"))
)
