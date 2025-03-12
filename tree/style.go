package tree

import "github.com/charmbracelet/lipgloss"

func GetStyles() (headerStyle, fileStyle, dirStyle, branchStyle lipgloss.Style) {

    headerStyle = lipgloss.NewStyle().
        Foreground(lipgloss.Color(AppConfig.ColorDir)).
        Bold(true).
        Padding(0, 1).
        Margin(0, 1)

    fileStyle = lipgloss.NewStyle().
        Foreground(lipgloss.Color(AppConfig.ColorFile))

    dirStyle = lipgloss.NewStyle().
        Foreground(lipgloss.Color(AppConfig.ColorDir)).
        Bold(true)

    branchStyle = lipgloss.NewStyle().
        Foreground(lipgloss.Color(AppConfig.ColorBranch))

    return
}