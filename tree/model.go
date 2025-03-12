package tree

import (
    "strings"

    tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
    tree []string
    path string
}

func NewModel(path string) Model {
    LoadConfig()
    tree := BuildTree(path)
    return Model{tree: tree, path: path}
}

func (m Model) Init() tea.Cmd {
    return tea.Quit
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    return m, tea.Quit
}

func (m Model) View() string {
    output := strings.Join(m.tree, "\n")
    return output
}