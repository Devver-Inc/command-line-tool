package app

import (
	"flag"
	"fmt"
	"os"

	"github.com/Devver-Inc/cli/internal/commands"
	"github.com/Devver-Inc/cli/internal/types"
	tea "github.com/charmbracelet/bubbletea"
)

var mockedNodes = [5]types.Node{
	{ID: "1", Name: "Hello", URL: "https://devver.fr/node1"},
	{ID: "2", Name: "World", URL: "https://devver.fr/node2"},
	{ID: "3", Name: "Test", URL: "https://devver.fr/node3"},
	{ID: "4", Name: "Demo", URL: "https://devver.fr/node4"},
	{ID: "5", Name: "Example", URL: "https://devver.fr/node5"},
}

type Model struct {
	types.Model
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "up", "k":
			if m.Cursor > 0 {
				m.Cursor--
			}
		case "down", "j":
			if m.Cursor < len(m.Nodes)-1 {
				m.Cursor++
			}
		case "enter", " ":
			_, ok := m.Selected[m.Cursor]
			if ok {
				delete(m.Selected, m.Cursor)
			} else {
				m.Selected[m.Cursor] = struct{}{}
			}
		}
	}
	return m, nil
}

func (m Model) View() string {
	s := "What should we buy at the market?\n\n"
	for i, choice := range m.Nodes {
		cursor := " "
		if m.Cursor == i {
			cursor = ">"
		}

		checked := " "
		if _, ok := m.Selected[i]; ok {
			checked = "x"
		}

		s += fmt.Sprintf("%s [%s] %s\n", cursor, checked, choice.Name)
	}
	s += "\nPress q to quit.\n"
	return s
}

func InitialModel() Model {
	return Model{
		types.Model{
			Nodes:    mockedNodes[:],
			Selected: make(map[int]struct{}),
		},
	}
}

func ParseArgs() types.Config {
	interactive := flag.Bool("i", false, "Run in interactive mode")
	flag.Parse()

	args := flag.Args()
	config := types.Config{
		Interactive: *interactive,
	}

	if len(args) > 0 {
		config.Command = args[0]
		if len(args) > 1 {
			config.Args = args[1:]
		}
	}

	return config
}

func RunCommand(m Model, cmd string, args []string) {
	switch cmd {
	case "list":
		items := commands.List(m.Model)
		for _, item := range items {
			fmt.Printf("%s: %s (%s)\n", item.ID, item.Name, item.URL)
		}
	case "interactive":
		RunInteractive(m)
	default:
		fmt.Printf("Unknown command: %s\n", cmd)
		os.Exit(1)
	}
}

func RunInteractive(model Model) {
	p := tea.NewProgram(model)
	if _, err := p.Run(); err != nil {
		fmt.Printf("Error: %v", err)
		os.Exit(1)
	}
}
