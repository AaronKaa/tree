package main

import (
    "os"

    tea "github.com/charmbracelet/bubbletea"
    "github.com/AaronKaa/tree/tree"
)

func main() {
    tree.LoadConfig()

    path := "."
    if len(os.Args) > 1 {
        for _, arg := range os.Args[1:] {
            switch arg {
            case "--help", "--h":
                tree.PrintHelp()
                return
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
                if path == "." {
                    path = arg
                }
            }
        }
    }

    m := tree.NewModel(path)
    p := tea.NewProgram(m)
    p.Run()
}