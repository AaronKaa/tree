package main

import (
    "fmt"
    "os"

    tea "github.com/charmbracelet/bubbletea"
    "github.com/AaronKaa/tree/tree"
)

func main() {
    var path string
    tree.LoadConfig()

    if len(os.Args) > 1 {
        for _, arg := range os.Args[1:] {
            switch arg {
            case "--summarise", "--s":
                tree.AppConfig.Summarize = true
            case "--no-summarise", "--ns":
                tree.AppConfig.Summarize = false
            case "--dot-files", "--df":
                tree.AppConfig.HideDotFiles = false
            case "--no-dot-files", "--ndf":
                tree.AppConfig.HideDotFiles = true
            case "--dirs-only", "--do":
                tree.AppConfig.DirsOnly = true
            default:
                if path == "" {
                    path = arg
                }
            }
        }
    }

    if path == "" {
        path = "."
    }

    m := tree.NewModel(path)
    p := tea.NewProgram(m, tea.WithOutput(os.Stdout))

    if _, err := p.Run(); err != nil {
        fmt.Printf("Error: %v\n", err)
        os.Exit(1)
    }
}