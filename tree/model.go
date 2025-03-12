package tree

import (
    "strings"

    tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
    Path string
    Tree []string
}

func NewModel(path string) Model {
    m := Model{Path: path}
    m.Tree = BuildTree(path)
    m.Tree = append(m.Tree, "")

    return m
}

func (m Model) Init() tea.Cmd {
    return tea.Quit
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    return m, tea.Quit
}

func (m Model) View() string {
    output := strings.Join(m.Tree, "\n")
    return output
}