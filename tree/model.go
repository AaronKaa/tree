package tree

import (
    liptree "github.com/charmbracelet/lipgloss/tree"
    tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
    Path string
    Tree *liptree.Tree
}

func NewModel(path string) Model {
    m := Model{Path: path}
    m.Tree = BuildTree(path)
    return m
}

func (m Model) View() string {
    return m.Tree.String() + "\n"
}

func (m Model) Init() tea.Cmd {
    return tea.Quit
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    return m, tea.Quit
}
